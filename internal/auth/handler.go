package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler обрабатывает запросы, связанные с аутентификацией.
type AuthHandler struct {
	AuthService AuthService
}

// NewAuthHandler создает новый экземпляр обработчика аутентификации.
func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

// Login обрабатывает запрос на аутентификацию.
func (h *AuthHandler) Login(c *gin.Context) {
	// Извлечение данных аутентификации из запроса
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Вызов метода сервиса для выполнения аутентификации
	token, err := h.AuthService.Authenticate(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Отправка ответа с токеном аутентификации
	c.JSON(http.StatusOK, gin.H{"token": token})
}
