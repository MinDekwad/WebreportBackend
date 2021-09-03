package amlo

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetListProvinceNameTH(c echo.Context) error {

	var result interface{}

	provinceNameTH, err := h.amloService.GetProvinceNameTH()
	if err != nil {
		h.logger.Errorf("[Handler] func get province name has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500138,
			Message: "Get province name has error",
		})
	}

	result = provinceNameTH
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetAreaListSearch(c echo.Context) error {
	var result interface{}

	provinceNameTH := c.QueryParam("provincenameth")

	data, err := h.amloService.GetAllAreaListSearch(provinceNameTH)
	if err != nil {
		h.logger.Errorf("[Handler] Get search all area list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500139,
			Message: "Get search all area list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetArea(c echo.Context) error {
	var result interface{}
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	data, err := h.amloService.GetDataArea(ID)
	if err != nil {
		h.logger.Errorf("[Handler] Get data area has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500140,
			Message: "Get data area has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateConfigArea(c echo.Context) error {
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	updateDate := c.QueryParam("date")

	var req model.ConfigArea
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}

	_, err = h.amloService.SaveEditConfigArea(c.Request().Context(), ID, req, updateDate)
	if err != nil {
		h.logger.Errorf("[Handler] Edit config area has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500141,
			Message: "Edit config area has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) ApproveConfigArea(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)
	userData := c.QueryParam("userData")

	_, err = h.amloService.SaveApproveConfigArea(c.Request().Context(), ID, userData)
	if err != nil {
		h.logger.Errorf("[Handler] Approve config area has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500142,
			Message: "Approve config area has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) GetAreaList(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllAreaList()
	if err != nil {
		h.logger.Errorf("[Handler] Get search all area list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500139,
			Message: "Get search all area list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetAreaWaitingApprove(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAreaListWaitingApprove()
	if err != nil {
		h.logger.Errorf("[Handler] Get arae list waiting approve has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500203,
			Message: "Get arae list waiting approve has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
