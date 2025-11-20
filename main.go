package main

import (
	_ "google-ai-service/docs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Hello World
// @version         1.0
// @description     API simple de ejemplo con Gin y Swagger
// @host            localhost:8080
// @BasePath        /api/v1

type HelloResponse struct {
	Message string `json:"message" example:"Hello World!"`
}

// HelloWorld godoc
// @Summary      Hello World
// @Description  Retorna un mensaje de saludo
// @Tags         hello
// @Accept       json
// @Produce      json
// @Success      200  {object}  HelloResponse
// @Router       /hello [get]
func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, HelloResponse{
		Message: "Hello World!",
	})
}

func main() {
	router := gin.Default()

	// Redirección de raíz a Swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Ruta de Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", HelloWorld)
	}

	// Imprimir URL de Swagger
	port := "8080"
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/index.html", port)
	fmt.Printf("\n🚀 Servidor iniciado en el puerto %s\n", port)
	fmt.Printf("📚 Documentación Swagger disponible en: %s\n\n", swaggerURL)

	// Iniciar servidor en puerto 8080
	router.Run(":" + port)
}
