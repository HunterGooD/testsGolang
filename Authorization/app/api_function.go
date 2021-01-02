package app

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) signUp(c *gin.Context) {
	c.Header("Content-type", "application/json")
	dataUser := Authorization{}
	if err := c.BindJSON(&dataUser); err != nil {
		erre := ErrorResponse{
			Error: map[string]interface{}{
				"data_valid": "Not valid",
				"is_valid":   false,
				"example":    `{"login":"ExampleLogin","password":"ExamplePassword"}`,
			},
		}
		c.JSON(http.StatusBadRequest, erre)
		return
	}

	token, err := GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Ошибка генерации токена",
			"code":    500,
			"message": "попробуйте обновить страницу или обратитесь за помощью к администратору",
		})
		return
	}
	response := ResponseToken{
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}

func (a *App) signIn(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"action": "вход",
	})
}

// HashPassword хэш паролей bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash хэш паролей bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken Генерирует токен для JWT
func GenerateToken() (string, error) {
	randomData := make([]byte, 10)
	_, err := rand.Read(randomData)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write(randomData)
	return hex.EncodeToString(h.Sum(nil)), nil
}
