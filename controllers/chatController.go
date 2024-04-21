package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	controllerHelper "github.com/pocketGod/Basic_Gemini_Integration_in_GO/helpers"
	"github.com/pocketGod/Basic_Gemini_Integration_in_GO/models"
	"github.com/pocketGod/Basic_Gemini_Integration_in_GO/services"
)

func ChatbotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		controllerHelper.HandleError(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.ChatbotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		controllerHelper.HandleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := services.GenerateResponse(request)
	controllerHelper.EncodeResponse(w, response)
}

func TokenCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		controllerHelper.HandleError(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	tokenCount, err := services.GetTokenCount()
	if err != nil {
		log.Printf("Error getting token count: %v", err)
		controllerHelper.HandleError(w, "Error getting token count", http.StatusInternalServerError)
		return
	}

	tokenResponse := models.TokenResponse{
		TokenCount: tokenCount,
	}
	controllerHelper.EncodeResponse(w, tokenResponse)
}
