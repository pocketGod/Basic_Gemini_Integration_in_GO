package main

import (
	"net/http"

	"github.com/pocketGod/Basic_Gemini_Integration_in_GO/controllers"

	"github.com/joho/godotenv"
)

func main() {

	// get .env file
	err := godotenv.Load(".env")
	if err != nil {
		//  log.Fatalf("Error loading .env file: %s", err)
	}

	http.HandleFunc("/chatbot", controllers.ChatbotHandler)
	http.HandleFunc("/tokencount", controllers.TokenCountHandler)

	http.ListenAndServe(":8080", nil)
}
