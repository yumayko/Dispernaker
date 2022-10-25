package controller

import (
	"disperinaker-api/domain"
	"disperinaker-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SertServiceController struct {
	SertServ domain.SertServiceAdapter
}

func (sc *SertServiceController) CreateSertController(c echo.Context) error {
	sert := model.Sertificate{}
	c.Bind(&sert)

	_, err := sc.SertServ.CreateSertService(sert)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "Success Creating Sertificate",
		Data:    sert,
	}, "\t")
}

func (sc *SertServiceController) GetSertsController(c echo.Context) error {
	serts, err := sc.SertServ.GetSertsService()

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting sertificates",
		Data:    serts,
	}, "\t")
}

func (sc *SertServiceController) GetSertByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	sert, err := sc.SertServ.GetSertByIDService(int(intID))

	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success getting sertificate",
		Data:    sert,
	}, "\t")
}

func (sc *SertServiceController) UpdateSertController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	sert := model.Sertificate{}
	sert.ID = int(intID)
	c.Bind(&sert)

	err := sc.SertServ.UpdateSertService(sert, int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success updating sertificate",
		Data:    sert,
	}, "\t")
}

func (sc *SertServiceController) DeleteSertController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.SertServ.DeleteSertService(int(intID))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}, "\t")
	}

	return c.JSONPretty(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success deleting sertificate",
		Data:    nil,
	}, "\t")
}
