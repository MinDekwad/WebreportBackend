package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetInsertRanking(c echo.Context) error {
	var result interface{}
	//date := c.QueryParam("date")

	data, err := h.amloService.InsertRanking()
	if err != nil {
		h.logger.Errorf("[Handler] Insert ranking has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500155,
			Message: "Insert ranking has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
