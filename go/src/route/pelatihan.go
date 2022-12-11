package route

import (
	"disperinaker-api/config"
	"disperinaker-api/controller"
	"disperinaker-api/database"
	"disperinaker-api/repository"
	"disperinaker-api/service"

	"github.com/labstack/echo/v4"
)

func RegisterTrainGroupAPI(e *echo.Echo, conf config.Config) {
	db := database.InitDB(conf)

	repo := repository.NewTrain(db)

	svc := service.NewTrainService(repo, conf)

	cont := controller.TrainServiceController{
		TrainServ: svc,
	}

	apiCalTrain := e.Group("/minat")

	apiCalTrain.POST("", cont.CreateCalTrainController)
	apiCalTrain.GET("/all", cont.GetCalTrainsController)
	apiCalTrain.GET("/:id", cont.GetCalTrainByIDController)
	apiCalTrain.PUT("/:id", cont.UpdateCalTrainController)
	apiCalTrain.DELETE("/:id", cont.DeleteCalTrainController)

	apiPesTrain := e.Group("/pelatihan")

	apiPesTrain.POST("", cont.CreatePesTrainController)
	apiPesTrain.GET("/all", cont.GetPesTrainsController)
	apiPesTrain.GET("/:id", cont.GetPesTrainByIDController)
	apiPesTrain.PUT("/:id", cont.UpdatePesTrainController)
	apiPesTrain.DELETE("/:id", cont.DeletePesTrainController)

}
