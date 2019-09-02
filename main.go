package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app/handlers"
)

func main() {
	fmt.Println("Hello, world.")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.Hello)
	e.GET("/sign", handlers.Sign)

	e.Logger.Fatal(e.Start(":8080"))
}
