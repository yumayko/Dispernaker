package main

import (
	"disperinaker-api/config"
	"disperinaker-api/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.InitConfiguration()

	e := echo.New()
	e.Use(middleware.CORS())

	route.HealthAPI(e, conf)
	route.RegisterSertGroupAPI(e, conf)
	route.RegisterTrainGroupAPI(e, conf)

	e.Logger.Fatal(e.Start(config.InitConfiguration().SERVER_ADDRESS))
}
