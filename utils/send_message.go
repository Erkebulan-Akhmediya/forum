package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

func SendMessage(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	m := message{Message: msg}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Println("Error sending message:", err)
	}
}
