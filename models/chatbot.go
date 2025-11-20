package models

// ChatRequest representa la solicitud del usuario al chatbot
type ChatRequest struct {
	Message string `json:"message" binding:"required" example:"Hola, ¿cómo estás?"`
}

// ChatResponse representa la respuesta del chatbot
type ChatResponse struct {
	Response string `json:"response" example:"¡Hola! Estoy bien, gracias por preguntar. ¿En qué puedo ayudarte hoy?"`
	Model    string `json:"model" example:"deepseek-chat"`
}

// ErrorResponse representa un error
type ErrorResponse struct {
	Error string `json:"error" example:"mensaje de error"`
}

// DeepSeekRequest estructura para la API de DeepSeek
type DeepSeekRequest struct {
	Model    string          `json:"model"`
	Messages []DeepSeekMessage `json:"messages"`
}

// DeepSeekMessage representa un mensaje en el formato de DeepSeek
type DeepSeekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// DeepSeekResponse estructura de respuesta de la API de DeepSeek
type DeepSeekResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message DeepSeekMessage `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
