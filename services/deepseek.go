package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google-ai-service/models"
	"io"
	"net/http"
	"os"
)

// DeepSeekService maneja las interacciones con la API de DeepSeek
type DeepSeekService struct {
	APIKey  string
	APIURL  string
	Client  *http.Client
}

// NewDeepSeekService crea una nueva instancia del servicio DeepSeek
func NewDeepSeekService() *DeepSeekService {
	return &DeepSeekService{
		APIKey: os.Getenv("DEEPSEEK_API_KEY"),
		APIURL: os.Getenv("DEEPSEEK_API_URL"),
		Client: &http.Client{},
	}
}

// Chat envía un mensaje al chatbot de DeepSeek y retorna la respuesta
func (s *DeepSeekService) Chat(message string) (*models.ChatResponse, error) {
	// Preparar la solicitud para DeepSeek
	requestBody := models.DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []models.DeepSeekMessage{
			{
				Role:    "system",
				Content: "Eres un asistente útil y amigable.",
			},
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error al serializar la solicitud: %v", err)
	}

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", s.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud: %v", err)
	}

	// Agregar headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	// Enviar la solicitud
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al enviar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	// Verificar el código de estado
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error de la API: %s - %s", resp.Status, string(body))
	}

	// Parsear la respuesta
	var deepSeekResp models.DeepSeekResponse
	if err := json.Unmarshal(body, &deepSeekResp); err != nil {
		return nil, fmt.Errorf("error al parsear la respuesta: %v", err)
	}

	// Validar que haya al menos una respuesta
	if len(deepSeekResp.Choices) == 0 {
		return nil, fmt.Errorf("no se recibió respuesta del modelo")
	}

	// Crear la respuesta
	chatResponse := &models.ChatResponse{
		Response: deepSeekResp.Choices[0].Message.Content,
		Model:    deepSeekResp.Model,
	}

	return chatResponse, nil
}
