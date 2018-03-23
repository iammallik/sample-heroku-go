package handlers

import (
	"encoding/json"
	"net/http"
)

// Default Render Handler
func Index(w http.ResponseWriter, r *http.Request) {
	object := DefaultObject{Value: "Hello!"}
	body, err := json.Marshal(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type DefaultObject struct {
	Value string `json:"value"`
}
