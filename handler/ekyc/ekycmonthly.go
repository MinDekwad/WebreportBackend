package ekyc

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GeteKycMonthly(c echo.Context) error {

	var result interface{}
	date := c.QueryParam("date")

	dataEKycMonthly, err := h.ekycService.GetEKycMonthly(date)
	if err != nil {
		h.logger.Errorf("[Handler] Get report eKyc monthly has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500084,
			Message: "EKyc monthly has error",
		})
	}
	result = dataEKycMonthly

	return c.JSON(http.StatusOK, model.Response{Message: "Success", Data: result})
}

func (h Handler) GeteKycMonthlyTmdsAll(c echo.Context) error {

	var result interface{}
	date := c.QueryParam("date")

	dataEKycMonthlyTmdsAll, err := h.ekycService.GetEKycMonthlyTmdsAll(date)
	if err != nil {
		h.logger.Errorf("[Handler] Get report eKyc monthly tmds all has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500085,
			Message: "EKyc monthly tmds all has error",
		})
	}
	result = dataEKycMonthlyTmdsAll

	return c.JSON(http.StatusOK, model.Response{Message: "Success", Data: result})
}

func (h Handler) GeteKycMonthlyTcrbAll(c echo.Context) error {

	var result interface{}
	date := c.QueryParam("date")

	dataEKycMonthlyTcrbAll, err := h.ekycService.GetEKycMonthlyTcrbAll(date)
	if err != nil {
		h.logger.Errorf("[Handler] Get report eKyc monthly tcrb all has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500086,
			Message: "EKyc monthly tcrb all has error",
		})
	}

	result = dataEKycMonthlyTcrbAll
	return c.JSON(http.StatusOK, model.Response{Message: "Success", Data: result})
}
