package main

import (
	"PULZO-TEST/controllers"
	"PULZO-TEST/middleware"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar Variables de Entorno (.env)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}

	router := gin.Default()
	router.Use(setupCORS())

	// Rutas Publicas
	router.POST("/login", controllers.LoginHandler)
	router.GET("/admin", controllers.ListActiveTokens)

	// Rutas Protegidas
	protected := router.Group("/")
	protected.Use(middleware.TokenMiddleware())
	protected.GET("/characters", controllers.GetCharacters)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}

// CORS (Evita Conflicto en Puertos al hacer pruebas Locales)
func setupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
