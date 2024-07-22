package main

import (
	"encoding/json"
	"encoding/xml"
	"log"

	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type LogEntry struct {
	Message string `json:"message" xml:"message"`
	Level   string `json:"level" xml:"level"`
}

func main() {
	http.HandleFunc("/logs", logsHandler)
	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	rotateHook, err := logrotate.NewRotateHook(logrotate.Options{
		Filename:   "logs/app.log",
		MaxSize:    10 * 1024 * 1024, // 10MB
		MaxBackups: 3,
		MaxAge:     28 * 24 * time.Hour, // 28 days
	})
	if err != nil {
		log.Fatal(err)
	}
	logger.Hooks.Add(rotateHook)
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
	logger.WithFields(logrus.Fields{
		"level":   logEntry.Level,
		"message": logEntry.Message,
	}).Info()
	w.WriteHeader(http.StatusAccepted)
}
