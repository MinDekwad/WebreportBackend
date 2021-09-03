package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetDownloadAmloCustomer(c echo.Context) error {
	date := c.QueryParam("date")
	activedate := c.QueryParam("activedate")
	statusType := c.QueryParam("type")

	var result interface{}

	resData, err := h.amloService.DownloadAmloCustomer(date, statusType, activedate)
	if err != nil {
		h.logger.Errorf("[Handler] func download new customer amlo has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500180,
			Message: "Download download new customer amlo has error",
		})
	}

	result = resData
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
