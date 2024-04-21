package services

import (
	"context"
	"log"
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"github.com/pocketGod/Basic_Gemini_Integration_in_GO/models"
	"google.golang.org/api/option"
)

func getClientAndContext() (*genai.Client, context.Context) {
	ctx := context.Background()

	var GEMINI_API_KEY string = os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func GenerateResponse(req models.ChatbotRequest) models.ChatbotResponse {
	client, ctx := getClientAndContext()
	defer client.Close()

	fullPrompt := req.Context + ". " + req.Prompt

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		if text, ok := part.(genai.Text); ok {
			return models.ChatbotResponse{Response: string(text)}
		}
	}

	return models.ChatbotResponse{Response: "No valid response generated"}
}

func GetTokenCount() (int32, error) {
	client, ctx := getClientAndContext()
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	resp, err := model.CountTokens(ctx, genai.Text(""))
	if err != nil {
		return 0, err
	}

	return resp.TotalTokens, nil
}
