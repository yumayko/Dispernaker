package route

import (
	"disperinaker-api/config"
	"disperinaker-api/controller"
	"disperinaker-api/database"
	"disperinaker-api/repository"
	"disperinaker-api/service"

	"github.com/labstack/echo/v4"
)

func RegisterSertGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewSert(db)

	svc := service.NewSertService(repo, conf)

	cont := controller.SertServiceController{
		SertServ: svc,
	}

	apiSert := e.Group("/sertificate/calon")

	apiSert.POST("", cont.CreateSertController)
	apiSert.GET("/all", cont.GetSertsController)
	apiSert.GET("/:id", cont.GetSertByIDController)
	apiSert.PUT("/:id", cont.UpdateSertController)
	apiSert.DELETE("/:id", cont.DeleteSertController)
}
