package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetUpdateRankingOccupation(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.UpdateRankingOccupation()
	if err != nil {
		h.logger.Errorf("[Handler] Update ranking has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500156,
			Message: "Update ranking has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
