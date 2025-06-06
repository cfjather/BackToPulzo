// Servicio de Consulta de Datos, para acceder a información proveniente de una API pública (como Rick and Morty API).

package services

import (
	"io"
	"net/http"
)

// Solicitud HTTP GET a la API de Rick and Morty
func GetCharacters() ([]byte, error) {
	resp, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Devuelve la respuesta con los personajes
	return io.ReadAll(resp.Body)
}
