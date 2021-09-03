package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) PostCalculateRank(c echo.Context) error {
	var result interface{}
	dates := c.QueryParam("date")
	provincename := c.QueryParam("provincename")
	date := dates[:10]

	data, err := h.amloService.CalculateRank(date, provincename)
	if err != nil {
		h.logger.Errorf("[Handler] Post calculate rank has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500148,
			Message: "Post calculate rank has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

// delete this if test finish
/*
func (h Handler) GetUpdateOccu(c echo.Context) error {

	var result interface{}

	data, err := h.reportService.UpdateOccu()
	if err != nil {
		h.logger.Errorf("[Handler] UpdateOccu rank has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500149,
			Message: "UpdateOccu rank has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})

}
*/
