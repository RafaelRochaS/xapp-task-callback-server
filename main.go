package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	TaskID        string  `json:"taskId"`
	DeviceID      int     `json:"deviceId"`
	WorkloadSize  int     `json:"workloadSize"`
	ExecutionSite string  `json:"executionSite"`
	CreatedAt     int64   `json:"createdAt"`
	Duration      float64 `json:"duration"`
	Timestamp     int64   `json:"timestamp"`
}

var logChan = make(chan Data, 10000)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received from: ", r.RemoteAddr)
	defer r.Body.Close()

	var d Data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logChan <- d

	log.Println("Request finished processing")
	w.WriteHeader(http.StatusOK)
}

func main() {
	go fileWriter("results.jsonl")

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
}
