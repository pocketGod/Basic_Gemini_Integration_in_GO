package models

type ChatbotRequest struct {
	Prompt  string `json:"prompt"`
	Context string `json:"context"`
}

type ChatbotResponse struct {
	Response string `json:"response"`
}

type TokenResponse struct {
	TokenCount int32 `json:"token_count"`
}
