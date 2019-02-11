package main

import (
	"github.com/goodrain/go-demo/middleware"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":5000"))
}
