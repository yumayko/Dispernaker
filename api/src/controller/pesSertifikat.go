package controller

import (
	"disperinaker-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (sc *SertServiceController) CreatePesSertController(c echo.Context) error {
	sert := model.PesertaSertificate{}
	c.Bind(&sert)

	_, err := sc.SertServ.CreatePesSertService(sert)
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

func (sc *SertServiceController) GetPesSertsController(c echo.Context) error {
	serts, err := sc.SertServ.GetPesSertsService()

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

func (sc *SertServiceController) GetPesSertByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	sert, err := sc.SertServ.GetPesSertByIDService(int(intID))

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

func (sc *SertServiceController) UpdatePesSertController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	sert := model.PesertaSertificate{}
	sert.ID = int(intID)
	c.Bind(&sert)

	err := sc.SertServ.UpdatePesSertService(sert, int(intID))
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

func (sc *SertServiceController) DeletePesSertController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.SertServ.DeletePesSertService(int(intID))
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
