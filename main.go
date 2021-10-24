package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Greeting{
			Message: "Merhaba Onur Hocam ve Değerli Öğrencileri",
			Status:  "OK",
		})
	})

	e.GET("/health", getHealth)
	e.GET("/ping", ping)
	e.GET("/data", data)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
