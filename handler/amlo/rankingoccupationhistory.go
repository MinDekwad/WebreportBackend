package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetRankingOccupationHistory(c echo.Context) error {
	var result interface{}
	walletID := c.QueryParam("walletID")

	data, err := h.amloService.GetDataRankingOccupationHistory(walletID)
	if err != nil {
		h.logger.Errorf("[Handler] Get data ranking occupation history has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500150,
			Message: "Get data ranking occupation history has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
