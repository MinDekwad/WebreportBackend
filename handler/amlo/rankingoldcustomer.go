package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetOldCustomerList(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllOldCustomerList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all old customer list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500154,
			Message: "Get all old customer list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
