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

func (sc *SertServiceController) CreateCalSertController(c echo.Context) error {
	sert := model.CalonSertificate{}
	c.Bind(&sert)

	_, err := sc.SertServ.CreateCalSertService(sert)
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

func (sc *SertServiceController) GetCalSertsController(c echo.Context) error {
	serts, err := sc.SertServ.GetCalSertsService()

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

func (sc *SertServiceController) GetCalSertByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	sert, err := sc.SertServ.GetCalSertByIDService(int(intID))

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

func (sc *SertServiceController) UpdateCalSertController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)
	sert := model.CalonSertificate{}
	sert.ID = int(intID)
	c.Bind(&sert)

	err := sc.SertServ.UpdateCalSertService(sert, int(intID))
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

func (sc *SertServiceController) DeleteCalSertController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.ParseInt(id, 10, 64)

	err := sc.SertServ.DeleteCalSertService(int(intID))
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
