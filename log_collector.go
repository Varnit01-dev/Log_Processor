package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type LogEntry struct {
	Message string `json:"message" xml:"message"`
	Level   string `json:"level" xml:"level"`
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	logs := collectLogs()
	for _, log := range logs {
		err = insertLog(db, log)
		if err != nil {
			logger.Fatal(err)
		}
	}

	setupHTTP()
}

func setupHTTP() {
	http.HandleFunc("/logs", logsHandler)
	logger := logrus.New()
	logger.Level = logrus.InfoLevel

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

func connectDB() (*sql.DB, error) {
	connStr := "user=myuser dbname=mydb sslmode=disable"
	return sql.Open("postgres", connStr)
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id SERIAL PRIMARY KEY,
			level TEXT NOT NULL,
			message TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}

func collectLogs() []LogEntry {
	return []LogEntry{
		{Level: "info", Message: "Log message 1"},
		{Level: "error", Message: "Log message 2"},
	}
}

func insertLog(db *sql.DB, log LogEntry) error {
	_, err := db.Exec(`
		INSERT INTO logs (level, message)
		VALUES ($1, $2);
	`, log.Level, log.Message)
	return err
}
