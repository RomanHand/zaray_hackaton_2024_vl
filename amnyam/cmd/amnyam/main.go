package main

import (
	"amnyam/internal/engine"
	"amnyam/internal/logger"
	stdlog "log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	listenPort = "12013"
	logLevel   = "INFO"
	debugMode  = false
)

func main() {
	logLevel = getEnv("AMNYAM_LOG_LEVEL", logLevel)

	envDebug := os.Getenv("AMNYAM_LOG_DEBUG_MODE")
	if val, err := strconv.ParseBool(envDebug); err != nil {
		debugMode = val
	} else {
		stdlog.Fatalf("режим отладки журнала должен быть: false или true")
	}

	logMan, err := logger.New(logLevel, debugMode)
	if err != nil {
		stdlog.Fatalf("ошибка запуска журнала, %v", err)
	}

	listenPort = getEnv("AMNYAM_PORT", listenPort)
	engine, err := engine.New(listenPort, logMan)
	if err != nil {
		logMan.Error(err.Error())
	}

	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	engine.Run()
	logMan.Info("приложение запущено")

	<-doneChan

	engine.Shutdown()

	logMan.Info("приложение остановлено")
}

func getEnv(envValue, defaultValue string) string {
	value := os.Getenv(envValue)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
