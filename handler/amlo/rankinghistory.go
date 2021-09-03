package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetRankingHistory(c echo.Context) error {
	var result interface{}
	walletID := c.QueryParam("walletID")

	data, err := h.amloService.GetDataRankingHistory(walletID)
	if err != nil {
		h.logger.Errorf("[Handler] Get data ranking history has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500202,
			Message: "Get data ranking history has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
