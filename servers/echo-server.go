package servers

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func StartEchoServer() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(200, "Hello, Echo!")
	})

	e.Start(":" + port)
}
