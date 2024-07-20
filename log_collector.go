package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type LogEntry struct {
	Message string `json:"message" xml:"message"`
	Level   string `json:"level" xml:"level"`
}

func main() {
	http.HandleFunc("/logs", logsHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	var logEntry LogEntry

	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&logEntry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&logEntry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	log.Printf("%s: %s", logEntry.Level, logEntry.Message)
	w.WriteHeader(http.StatusAccepted)
}


