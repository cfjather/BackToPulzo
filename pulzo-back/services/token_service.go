// Servicio de Autenticación que genera y gestiona tokens temporales.
package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// secretKey almacena la clave secreta para firmar los tokens JWT
var secretKey []byte

// Carga la llave secreta desde .env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando .env file en init(): %v", err)
	}

	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		panic("La variable 'JWT_SECRET_KEY' es requerida, y no ha sido definida aun")
	}
	secretKey = []byte(key)
}

// TokenInfo lleva la cuenta de los usos de cada token
type TokenInfo struct {
	Uses int
}

// tokens guarda los tokens activos e informacion de su uso
var tokens = make(map[string]*TokenInfo)

// mu asegura acceso concurrente seguro al mapa de tokens
var mu sync.Mutex

// GenerateToken crea un token JWT y lo registra con cero usos
func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"created": time.Now().Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	//Guarda el Token
	mu.Lock()
	tokens[tokenString] = &TokenInfo{Uses: 0}
	mu.Unlock()

	log.Printf("[TOKEN] Token generado: %s", tokenString)

	return tokenString, nil
}

// ValidateToken verifica si un token es válido y si tiene usos restantes
func ValidateToken(tokenStr string) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	info, exists := tokens[tokenStr]
	if !exists {
		log.Printf("[TOKEN] Token no encontrado: %s", tokenStr)
		return 0, errors.New("Token inválido o no encontrado")
	}
	if info.Uses >= 5 {
		log.Printf("[TOKEN] Token expirado: %s", tokenStr)
		return 0, errors.New("Token expirado. Solicita uno nuevo")
	}

	// Incrementa el uso del Token y lleva la cuenta de sus usos restantes
	info.Uses++
	usesLeft := 5 - info.Uses

	return usesLeft, nil
}

// CreateTokenHandler genera y retorna un token por HTTP
func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := GenerateToken()
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// GetAllTokens devuelve todos los tokens y su info (Admin View)
func GetAllTokens() map[string]*TokenInfo {
	return tokens
}
