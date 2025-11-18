package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	TaskID        string  `json:"task_id"`
	ExecutionSite string  `json:"execution_site"`
	Timestamp     int64   `json:"timestamp"`
	Duration      float64 `json:"duration"`
}

var logChan = make(chan Data, 10000)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var d Data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	select {
	case logChan <- d:
	default:
		logChan <- d
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	done := make(chan struct{})

	go fileWriter("results.jsonl", done)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Listening on port 8080...")

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	close(logChan)
	<-done
}
