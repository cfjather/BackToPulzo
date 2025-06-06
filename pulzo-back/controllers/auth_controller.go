package controllers

import (
	"net/http"

	"PULZO-TEST/services"

	"github.com/gin-gonic/gin"
)

// Genera Token y lo devuelve en formato JSON
func LoginHandler(c *gin.Context) {

	// Solo acepta solicitudes POST
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Genera el Token
	token, err := services.GenerateToken()
	if err != nil {
		// Si ocurre error interno al generar el Token
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// Devuelve el Token al cliente
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Verifica si el token es valido
func ValidateTokenHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token requerido en header Authorization"})
		return
	}

	_, err := services.ValidateToken(token)
	if err != nil {

		// Si es Invalido o no tiene usos disponibles
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Si el token es Valido, puede continuar
	c.JSON(http.StatusOK, gin.H{"message": "Token válido", "status": "success"})
}
