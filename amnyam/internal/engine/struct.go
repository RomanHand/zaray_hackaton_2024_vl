package engine

type RecordFragments struct {
	Name       string      `json:"name"`
	Source     string      `json:"source"`
	Violations []Violation `json:"violations"`
}

type Violation struct {
	Preview string  `json:"preview"`
	Start   float32 `json:"start"`
	End     float32 `json:"end"`
	//Timestamp string  `json:"timestamp"`
}

type ResponseViolations struct {
	Violations []ResponseViolation `json:"violations"`
}

type ResponseViolation struct {
	Preview string `json:"preview"`
	Clip    string `json:"clip"`
}

// type Record struct {
// 	Name       string      `json:"name"`
// 	Source     string      `json:"source"`
// 	Violations []Violation `json:"violations"`
// }
// Clipes  []Clip `json:"clipes"`
// type Clip struct {
// 	Name string `json:"name"`
// 	Path string `json:"path"`
// }

// type VideoRecord struct {
// 	Name string `json:"name"`
// }

// {
// 	"name": "under1.mp4",
// 	"violations": [
// 	  {
// 		"preview": "under1.mp4_0.jpeg",
// 		"start": 0.2,
// 		"end": 3.7
// 	  },
// 	  {
// 		"preview": "under1.mp4_1.jpeg",
// 		"start": 9.9,
// 		"end": 10.1
// 	  }
// 	]
// }

// {
// 	"name": "xxx.mp4",
// 	"path_source": "temp\\xxxxxx.mp4",
// 	"violations": [
// 	  {
// 		"preview": "123.jpeg",
// 		"timestamp": "00:12:36:789"
// 	  },
// 	  {
// 		"preview": "111.jpeg",
// 		"timestamp": "00:12:38:789"
// 	  },
// 	  {
// 		"preview": "211.jpeg",
// 		"timestamp": "00:12:38:789"
// 	  }
// 	]
// }
