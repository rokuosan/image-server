package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rokuosan/image-server/internal/router"
)

func main() {
	e := echo.New()
	r := router.NewRouter()
	r.Configure(e)

	e.Logger.Fatal(e.Start(":8080"))
}
