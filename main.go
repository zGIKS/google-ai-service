package main

import (
	_ "google-ai-service/docs"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google-ai-service/models"
	"google-ai-service/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Chatbot con DeepSeek
// @version         1.0
// @description     API con chatbot integrado usando DeepSeek AI
// @host            localhost:8080
// @BasePath        /api/v1

// Chat godoc
// @Summary      Chat con DeepSeek
// @Description  Envía un mensaje al chatbot de DeepSeek y obtiene una respuesta
// @Tags         chatbot
// @Accept       json
// @Produce      json
// @Param        request  body      models.ChatRequest  true  "Mensaje del usuario"
// @Success      200      {object}  models.ChatResponse
// @Failure      400      {object}  models.ErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /chat [post]
func Chat(c *gin.Context) {
	var req models.ChatRequest

	// Validar el request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Mensaje requerido: " + err.Error(),
		})
		return
	}

	// Crear el servicio de DeepSeek
	deepSeekService := services.NewDeepSeekService()

	// Enviar el mensaje al chatbot
	response, err := deepSeekService.Chat(req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Error al comunicarse con DeepSeek: " + err.Error(),
		})
		return
	}

	// Retornar la respuesta
	c.JSON(http.StatusOK, response)
}

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró archivo .env, usando variables de entorno del sistema")
	}

	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todos los orígenes (puedes especificar dominios específicos)
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Redirección de raíz a Swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Ruta de Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/chat", Chat)
	}

	// Obtener puerto desde variable de entorno o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Imprimir URL de Swagger
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/index.html", port)
	fmt.Printf("\n🚀 Servidor iniciado en el puerto %s\n", port)
	fmt.Printf("📚 Documentación Swagger disponible en: %s\n\n", swaggerURL)

	// Iniciar servidor
	router.Run(":" + port)
}
