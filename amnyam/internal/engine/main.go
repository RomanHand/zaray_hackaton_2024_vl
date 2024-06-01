package engine

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"

	"encoding/json"
)

type Engine struct {
	httpserver *http.Server
	log        *zap.Logger

	workDir string
}

var LogCompiledCommand bool = false

func New(listenPort string, log *zap.Logger) (*Engine, error) {
	ex, err := os.Executable() // полный путь файла приложения
	if err != nil {
		return nil, err
	}

	return &Engine{
		httpserver: &http.Server{
			Addr: fmt.Sprintf(":%s", listenPort),
			// WriteTimeout: time.Second * 15,
			// ReadTimeout:  time.Second * 15,
			// IdleTimeout:  time.Second * 60,
		},
		log:     log,
		workDir: filepath.Dir(ex),
	}, nil
}

func (e *Engine) Run() {
	fragmentsFS := http.StripPrefix("/fragments/", http.FileServer(http.Dir("./fragments/")))
	previewFS := http.StripPrefix("/previews/", http.FileServer(http.Dir("./previews/")))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", e.mainHandler).Methods("GET") // for self debug
	router.HandleFunc("/upload", e.uploadHandler).Methods("POST")
	router.PathPrefix("/fragments/").Handler(fragmentsFS)
	router.PathPrefix("/previews/").Handler(previewFS)

	e.httpserver.Handler = router

	go func() {
		if err := e.httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.log.Fatal("engine server error", zap.Error(err))
		}
	}()

	e.log.Info("engine: started.")
}

func (e *Engine) mainHandler(w http.ResponseWriter, r *http.Request) {
	e.log.Info(fmt.Sprintf("запрос '%s'; тип запроса '%s' \n", r.URL.Path, r.Method))
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func (e *Engine) uploadHandler(w http.ResponseWriter, r *http.Request) {
	e.log.Info(fmt.Sprintf("запрос '%s'; тип запроса '%s' \n", r.URL.Path, r.Method))
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	originFileName := fileHeader.Filename

	err = os.MkdirAll("./origins", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(fmt.Sprintf("./origins/%s", originFileName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	modelPath := fmt.Sprintf("%s/best.pt", strings.Replace(e.workDir, "\\", "/", -1))
	recordPath := fmt.Sprintf("%s/origins/%s", strings.Replace(e.workDir, "\\", "/", -1), originFileName)

	//
	// Запуск скрипта обнаружения нарушений
	cmdDetector := exec.Command("./detector.exe", originFileName, recordPath, modelPath)

	outputDetector, err := cmdDetector.CombinedOutput()
	if err != nil {
		e.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e.log.Info(fmt.Sprintf("Received file: %s, frames: %v", originFileName, string(outputDetector)))

	//
	// Парсинг файла с отрезками нарушений
	fileContent, err := os.ReadFile("data.json")
	if err != nil {
		e.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(fileContent) == 0 {
		respondWithJSON(w, ResponseViolations{
			Violations: []ResponseViolation{},
		})
		return
	}

	recordFragments := &RecordFragments{}
	err = json.Unmarshal(fileContent, recordFragments)
	if err != nil {
		e.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fragmentsDir := fmt.Sprintf("%s/fragments", strings.Replace(e.workDir, "\\", "/", -1))
	fragmentsFileDir := fmt.Sprintf("%s/%s", fragmentsDir, originFileName)
	// Каталог для нарезанных отрезков с нарушениями
	err = os.MkdirAll(fragmentsFileDir, os.ModePerm)
	if err != nil {
		e.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//
	// Нарезка видео записи на фрагменты с нарушениями
	files, err := e.Cutter(recordPath, recordFragments.Violations, fragmentsFileDir)
	if err != nil {
		e.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возврат массива с нарушениями
	respViolations := &ResponseViolations{}
	for i, p := range files {
		fileName := filepath.Base(p)
		respViolation := ResponseViolation{
			Preview: fmt.Sprintf("/previews/%s_%v.png", originFileName, i),
			Clip:    fmt.Sprintf("/fragments/%s/%s", originFileName, fileName),
		}
		respViolations.Violations = append(respViolations.Violations, respViolation)
	}

	respondWithJSON(w, respViolations)
}

// Cutter - нарезка видео записи по фрагментам
func (e *Engine) Cutter(source string, fragments []Violation, outputDir string) ([]string, error) {
	files := []string{}

	for id, frag := range fragments {
		outFile := fmt.Sprintf("%s/clip_%v.mp4", outputDir, id)

		endPosition := frag.End - frag.Start
		stream := ffmpeg.Input(source, ffmpeg.KwArgs{"ss": frag.Start})
		// .Silent(true) без вывода результирующей команды запуска ffmpeg
		err := stream.Output(outFile, ffmpeg.KwArgs{"t": endPosition}).OverWriteOutput().Silent(true).Run()
		if err != nil {
			return nil, err
		}
		files = append(files, outFile)
	}

	return files, nil
}

func (e *Engine) Shutdown() {
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := e.httpserver.Shutdown(shutdownCtx); err != nil {
		e.log.Error("Engine server shutdown error: %v", zap.Error(err))
	}
	e.log.Info("Engine server shutdown complete")
}

// ответ в формате JSON
func respondWithJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Printf("Error: response writer: %v", err)
		return
	}
}
