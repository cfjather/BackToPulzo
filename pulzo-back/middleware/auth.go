package middleware

import (
	"net/http"
	"strings"

	"PULZO-TEST/services"

	"github.com/gin-gonic/gin"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Obtener Token
		authHeader := c.GetHeader("Authorization")

		// Si no hay Token, retornar Error
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta token"})
			c.Abort()
			return
		}

		token := strings.TrimSpace(authHeader)

		// Se Valida el Token
		usesLeft, err := services.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Guardar cuantos usos le quedan al Token
		c.Set("usesLeft", usesLeft)

		c.Next()
	}
}
