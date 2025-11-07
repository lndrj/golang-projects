package routes

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/:code", handlers.RedirectURL)
	r.POST("/shorten", handlers.ShortenURL)

	return r
}
