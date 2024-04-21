package controllerHelper

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, errMessage string, errCode int) {
	http.Error(w, errMessage, errCode)
}

func EncodeResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		HandleError(w, "Error processing response", http.StatusInternalServerError)
	}
}
