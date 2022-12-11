package controller

import (
	"disperinaker-api/domain"
	"disperinaker-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TrainServiceController struct {
	TrainServ domain.TrainServiceAdapter
}

func (sc *TrainServiceController) CreateCalTrainController(c echo.Context) error {
	train := model.PesertaTestMinatBakat{}
	c.Bind(&train)

	_, err := sc.TrainServ.CreateCalTrainService(train)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "Success Creating Training",
		Data:    train,
	}, "\t")
}

func (sc *TrainServiceController) GetCalTrainsController(c echo.Context) error {
	trains, err := sc.TrainServ.GetCalTrainsService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting trainings",
		Data:    trains,
	}, "\t")
}

func (sc *TrainServiceController) GetCalTrainByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	train, err := sc.TrainServ.GetCalTrainByIDService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting training",
		Data:    train,
	}, "\t")
}

func (sc *TrainServiceController) UpdateCalTrainController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	train := model.PesertaTestMinatBakat{}
	train.ID = int(intID)
	c.Bind(&train)

	err := sc.TrainServ.UpdateCalTrainService(train, int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success updating training",
		Data:    train,
	}, "\t")
}

func (sc *TrainServiceController) DeleteCalTrainController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.TrainServ.DeleteCalTrainService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting training",
		Data:    nil,
	}, "\t")
}
