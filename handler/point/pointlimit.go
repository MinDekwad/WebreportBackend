package point

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetLimitPoint(c echo.Context) error {
	var result interface{}

	data, err := h.pointService.GetAllLimitPoint()
	if err != nil {
		h.logger.Errorf("[Handler] Get limit point has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500120,
			Message: "Get limit point has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateLimitPoint(c echo.Context) error {
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	var req model.LimitPoint
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}
	_, err = h.pointService.SaveEditLimitPoint(c.Request().Context(), Uid, req)
	if err != nil {
		h.logger.Errorf("[Handler] Edit config point has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500119,
			Message: "Edit limit point has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}
