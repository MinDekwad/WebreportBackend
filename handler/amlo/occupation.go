package amlo

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreateConfigOccupation(c echo.Context) error {

	var result interface{}
	var req model.CreateConfigOccupation
	today := c.QueryParam("date")
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}

	data, err := h.amloService.ServiceCreateConfigOccupation(c.Request().Context(), req, today)
	if err != nil {
		h.logger.Errorf("[Handler] func create config occupation has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500131,
			Message: "Create config occupation has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetOccupationList(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllOccupationList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all occupation list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500132,
			Message: "Get all occupation list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) DelConfigOccupation(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveDelConfigOccupation(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Delete config occupation has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500133,
			Message: "Delete config occupation has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) GetOccupation(c echo.Context) error {
	var result interface{}
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	data, err := h.amloService.GetDataOccupation(ID)
	if err != nil {
		h.logger.Errorf("[Handler] Get data occupation has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500134,
			Message: "Get data occupation has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateConfigOccupation(c echo.Context) error {
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	updateDate := c.QueryParam("date")

	var req model.CreateConfigOccupation
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}

	_, err = h.amloService.SaveEditConfigOccupation(c.Request().Context(), ID, req, updateDate)
	if err != nil {
		h.logger.Errorf("[Handler] Edit config occupation has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500135,
			Message: "Edit config occupation has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) ApproveConfigOccupation(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)
	userData := c.QueryParam("userData")

	_, err = h.amloService.SaveApproveConfigOccupation(c.Request().Context(), ID, userData)
	if err != nil {
		h.logger.Errorf("[Handler] Approve config occupation has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500136,
			Message: "Approve config occupation has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) GetOccupationListSearch(c echo.Context) error {
	var result interface{}

	occupationName := c.QueryParam("occupationname")
	rank := c.QueryParam("rank")

	data, err := h.amloService.GetAllOccupationListSearch(occupationName, rank)
	if err != nil {
		h.logger.Errorf("[Handler] Get search all occupation list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500137,
			Message: "Get search all occupation list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetOccuWaitingApprove(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetOccuListWaitingApprove()
	if err != nil {
		h.logger.Errorf("[Handler] Get occupation list waiting approve has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500203,
			Message: "Get occupation list waiting approve has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
