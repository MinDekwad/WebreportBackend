package handler

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) UpdateReportKYC(c echo.Context) error {
	date := c.QueryParam("date")
	var result interface{}
	dataUpdateReportKYC, err := h.reportService.PostUpdateReportKYC(date)
	if err != nil {
		h.logger.Errorf("[Handler] Update report kyc has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500114,
			Message: "Update report kyc has error",
		})
	}
	result = dataUpdateReportKYC
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})

	// date := c.QueryParam("date")
	// var result interface{}
	// err := h.reportService.PostUpdateReportKYC(date)
	// if err != nil {
	// 	h.logger.Errorf("[Handler] Update report kyc has error %s", err)
	// }
	// result = err
	// return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
	//return nil
}
