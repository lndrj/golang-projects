package handlers

import (
	"math/rand"
	"net/http"
	"url-shortener/db"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	code := generateCode(6)

	_, err := db.DB.Exec(`INSERT INTO urls (code, original_url) VALUES (?, ?)`, code, req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot insert new url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080" + code})
}

func RedirectURL(c *gin.Context) {

	code := c.Param("code")
	var original string
	err := db.DB.QueryRow(`SELECT original_url FROM urls WHERE code = ?`, code).Scan(&original)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "code not found"})
		return
	}

	c.JSON(http.StatusOK, original)
}

func generateCode(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, n)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)

}
