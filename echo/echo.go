package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{
            "message": "pong",
        })
    })

    e.Start(":8080")
}
