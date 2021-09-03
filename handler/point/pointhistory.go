package point

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreatePoint(c echo.Context) error {
	date := c.QueryParam("date")
	var result interface{}

	calculatePoint, err := h.pointService.CalculatePoint(date)
	if err != nil {
		h.logger.Errorf("[Handler] Calculate point has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500121,
			Message: "Calculate point has rrror",
		})
	}

	result = calculatePoint
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetListHistoryPoint(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	var result interface{}

	historyPoint, err := h.pointService.GetHistoryPoint(startDate, endDate)
	if err != nil {
		h.logger.Errorf("[Handler] Get history point has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500122,
			Message: "Get history point has rrror",
		})
	}
	result = historyPoint
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
