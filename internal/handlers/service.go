package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type AllRounderStruct struct {
}

func (h *AllRounderStruct) StatusHandler(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Service string    `json:"service"`
		Level   string    `json:"level"`
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}

	var body Body

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if body.Level == "" {
		body.Level = "level : info"
	}
	if body.Service == "" {
		body.Service = "service : unknown-service"
	}
	if body.Message == "" {
		body.Message = "message : no message provided"
	}
	if body.Time.IsZero() {
		body.Time = time.Now().UTC()
	}

	log.Printf("Log received: %+v\n", body)

	WriteToLogFile(body)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "log received"})
}

func WriteToLogFile(logMessage interface{}) {
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	jsonBytes, err := json.Marshal(logMessage)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON bytes plus newline
	if _, err := f.Write(append(jsonBytes, '\n')); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
