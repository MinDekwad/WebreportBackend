package point

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetListHistoryPointExport(c echo.Context) error {

	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")

	var result interface{}

	createPointCSV, err := h.pointService.GetHistoryPointExport(start_date, end_date)
	if err != nil {
		h.logger.Errorf("[Handler] Get history point export has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500129,
			Message: "Get history point export has error",
		})
	}

	result = createPointCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
