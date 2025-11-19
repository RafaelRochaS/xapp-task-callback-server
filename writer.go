package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func fileWriter(path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer f.Close()

	enc := json.NewEncoder(f)

	for data := range logChan {
		data.Timestamp = time.Now().Unix()
		if err := enc.Encode(data); err != nil {
			log.Printf("Error writing log: %v", err)
		}
	}
}
