package route

import (
	"disperinaker-api/config"
	"disperinaker-api/model"

	"github.com/labstack/echo/v4"
)

func HealthAPI(e *echo.Echo, conf config.Config) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSONPretty(200, model.Response{
			Code:    200,
			Message: "sehat gan",
		}, "\t")
	})
}
