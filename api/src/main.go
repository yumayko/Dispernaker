package main

import (
	"disperinaker-api/config"
	"disperinaker-api/route"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfiguration()

	e := echo.New()

	route.HealthAPI(e, conf)
	route.RegisterSertGroupAPI(e, conf)

	e.Logger.Fatal(e.Start(config.InitConfiguration().SERVER_ADDRESS))
}
