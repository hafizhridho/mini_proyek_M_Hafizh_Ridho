package main

import (
	"proyek/configs"
	"proyek/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDb()
	configs.Loadenv()
	e := echo.New()
	e.GET("/user", controllers.Register)
	e.Start(":8000")

}