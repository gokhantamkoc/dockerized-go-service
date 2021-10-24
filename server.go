package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Response struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type Ping struct {
	Message string `json:"message"`
}

type Landmark struct {
	Name string	`json:"name"`
	Country string `json:"country"`
}

func getHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{
		Message: "Healthy",
		Data: nil,
	})
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &Ping{
		Message: "Pong",
	})
}

func data(c echo.Context) error {
	query := "SELECT name, country FROM landmarks"
	dbClient := NewDBClient()
	rows, err := dbClient.db.Query(query)
	if err != nil {
		return c.JSON(500, &Response{
			Message: err.Error(),
			Data: nil,
		})
	}
	defer rows.Close()

	var landmarks []Landmark

	for rows.Next() {
		var landmark Landmark
		if err = rows.Scan(&landmark.Name, &landmark.Country); err != nil {
			return c.JSON(500, &Response{
				Message: err.Error(),
				Data: nil,
			})
		}
		landmarks = append(landmarks, landmark)
	}
	return c.JSON(200, &Response{
		Message: "",
		Data: landmarks,
	})
}