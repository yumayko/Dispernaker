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

	apiCalSert := e.Group("/sertificate/calon")

	apiCalSert.POST("", cont.CreateCalSertController)
	apiCalSert.GET("/all", cont.GetCalSertsController)
	apiCalSert.GET("/:id", cont.GetCalSertByIDController)
	apiCalSert.PUT("/:id", cont.UpdateCalSertController)
	apiCalSert.DELETE("/:id", cont.DeleteCalSertController)

	apiPesSert := e.Group("/sertificate/peserta")

	apiPesSert.POST("", cont.CreatePesSertController)
	apiPesSert.GET("/all", cont.GetPesSertsController)
	apiPesSert.GET("/:id", cont.GetPesSertByIDController)
	apiPesSert.PUT("/:id", cont.UpdatePesSertController)
	apiPesSert.DELETE("/:id", cont.DeletePesSertController)

}
