package main

import (
	"log"
	"net/http"

	"hello/middlewere"
	"hello/models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello golang")
	})

	e.POST("/users", SaveUser)
	e.PUT("/users/:id", UpdateUser)
	e.GET("/users", GetUser)
	e.GET("/users/:name2/", GetUser)
	e.DELETE("/users", DeleteUser)

	e.Logger.Fatal(e.Start(":5000"))
}

func GetUser(c echo.Context) error {
	name := c.QueryParam("name1")
	name2 := c.Param("name2")

	var response = map[string]interface{}{
		"name1": name,
		"name2": name2,
	}
	return c.JSON(http.StatusOK, response)
}

func SaveUser(c echo.Context) error {
	middlewere.Form(c)
	jsonData := c.Get("params").(map[string]interface{})

	user := models.NewParamsFromData(jsonData)
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	middlewere.Form(c)
	id := c.Param("id")
	log.Print(id)

	jsonData := c.Get("params").(map[string]interface{})
	user := models.NewParamsFromData(jsonData)
	return c.JSON(http.StatusCreated, user)
}

func DeleteUser(c echo.Context) error {
	response := map[string]interface{}{
		"message": "Deleted!!",
	}

	return c.JSON(http.StatusOK, response)
}
