package controller

import (
	"disperinaker-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (sc *TrainServiceController) CreatePesTrainController(c echo.Context) error {
	train := model.PesertaPelatihan{}
	c.Bind(&train)

	_, err := sc.TrainServ.CreatePesTrainService(train)
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

func (sc *TrainServiceController) GetPesTrainsController(c echo.Context) error {
	trains, err := sc.TrainServ.GetPesTrainsService()

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

func (sc *TrainServiceController) GetPesTrainByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	train, err := sc.TrainServ.GetPesTrainByIDService(int(intID))

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

func (sc *TrainServiceController) UpdatePesTrainController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	train := model.PesertaPelatihan{}
	train.ID = int(intID)
	c.Bind(&train)

	err := sc.TrainServ.UpdatePesTrainService(train, int(intID))
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

func (sc *TrainServiceController) DeletePesTrainController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.TrainServ.DeletePesTrainService(int(intID))
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
