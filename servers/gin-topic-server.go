package servers

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func StartGinServer() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	router.Run(":" + port)
}
