package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LogEntry struct {
	Message string `json:"message"`
	Level   string `json:"level"`
}

func main() {
	http.HandleFunc("/logs", logsHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	var logEntry LogEntry
	err := json.NewDecoder(r.Body).Decode(&logEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%s: %s", logEntry.Level, logEntry.Message)
	w.WriteHeader(http.StatusAccepted)
}
