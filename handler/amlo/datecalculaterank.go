package amlo

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetDateCalculateRankList(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllDateCalculateRankList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all date calculate rank list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500143,
			Message: "Get all date calculate rank list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetDateCalculateRank(c echo.Context) error {

	var result interface{}
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	data, err := h.amloService.GetDataDateCalculateRank(ID)
	if err != nil {
		h.logger.Errorf("[Handler] Get data date calculate rank has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500144,
			Message: "Get data date calculate rank has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateConfigDateCalculateRank(c echo.Context) error {
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	updateDate := c.QueryParam("date")

	var req model.ConfigDateCalculateRank
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}

	_, err = h.amloService.SaveEditConfigDateCalculateRank(c.Request().Context(), ID, req, updateDate)
	if err != nil {
		h.logger.Errorf("[Handler] Edit config occupation has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500145,
			Message: "Edit config occupation has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) ApproveConfigDateCalculateRank(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveApproveConfigDateCalculateRank(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Approve config date calculate rank has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500146,
			Message: "Approve config date calculate rank has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}
