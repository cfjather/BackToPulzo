package controllers

import (
	"net/http"

	"PULZO-TEST/services"

	"github.com/gin-gonic/gin"
)

func ListActiveTokens(c *gin.Context) {

	// Obtiene copia del mapa de tokens activos y sus usos
	allTokens := services.GetAllTokens()

	// Retorna tokens activos como JSON
	c.JSON(http.StatusOK, allTokens)
}
