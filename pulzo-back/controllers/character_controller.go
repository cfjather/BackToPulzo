package controllers

import (
	"encoding/json"
	"net/http"

	"PULZO-TEST/services"

	"github.com/gin-gonic/gin"
)

func GetCharacters(c *gin.Context) {
	// Obtiene los usos restantes del Token
	usesLeftRaw, exists := c.Get("usesLeft")
	usesLeft := 0
	if exists {
		usesLeft = usesLeftRaw.(int)
	}

	// Obtener los personajes desde la API
	data, err := services.GetCharacters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener los personajes"})
		return
	}

	// Decodifica el JSON en estructura genérica para enviar al cliente
	var characters interface{}
	if err := json.Unmarshal(data, &characters); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los personajes"})
		return
	}

	// Devuelve un JSON que incluye usos restantes y personajes
	c.JSON(http.StatusOK, gin.H{
		"usesLeft":   usesLeft,
		"characters": characters,
	})
}
