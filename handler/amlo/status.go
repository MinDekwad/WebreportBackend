package amlo

import (
	//"io"
	//"strconv"
	//"strings"

	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) Calculate(c echo.Context) error {
	var result interface{}

	slug := c.Param("slug")
	name := c.QueryParam("name")

	switch slug {
	case "area":

		requestID, err := h.amloService.CreateFileInsert("area")
		if err != nil {
			return err
		}

		go func() {
			err = h.amloService.CalculateRankByArea(name, requestID.ID)
			if err != nil {
				h.logger.Errorf("[Handler] Report %s insert data error", err)
			}
		}()
		return c.JSON(http.StatusOK, model.Response{Message: "success", Data: requestID})

	case "occu":

		requestID, err := h.amloService.CreateFileInsert("occupation")
		if err != nil {
			return err
		}

		go func() {
			err = h.amloService.CalculateRankByOccu(requestID.ID)
			if err != nil {
				h.logger.Errorf("[Handler] Report %s insert data error", err)
			}
		}()
		return c.JSON(http.StatusOK, model.Response{Message: "success", Data: requestID})

	case "wl":

		requestID, err := h.amloService.CreateFileInsert("watchlist")
		if err != nil {
			return err
		}

		go func() {
			err = h.amloService.CalculateRankByWl(requestID.ID)
			if err != nil {
				h.logger.Errorf("[Handler] Report %s insert data error", err)
			}
		}()
		return c.JSON(http.StatusOK, model.Response{Message: "success", Data: requestID})

	case "transaction":
		requestID, err := h.amloService.CreateFileInsert("transaction")
		if err != nil {
			return err
		}

		go func() {
			err = h.amloService.CalculateRankByTF(requestID.ID)
			if err != nil {
				h.logger.Errorf("[Handler] Report %s insert data error", err)
			}
		}()
		return c.JSON(http.StatusOK, model.Response{Message: "success", Data: requestID})

	}
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) CheckStatusCalculate(c echo.Context) error {
	var result interface{}
	slug := c.Param("slug")
	id := c.QueryParam("request_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	switch slug {
	case "area":
		dataAreaInsertStatus, err := h.amloService.GetAreaStatus(ID)
		if err != nil {
			h.logger.Errorf("[Handler] Get status area has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500098,
				Message: "Status area has error",
			})
		}
		result = dataAreaInsertStatus

	case "occu":
		dataOccuInsertStatus, err := h.amloService.GetOccuStatus(ID)
		if err != nil {
			h.logger.Errorf("[Handler] Get status occupation has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500098,
				Message: "Status occupation has error",
			})
		}
		result = dataOccuInsertStatus

	case "wl":
		dataWlInsertStatus, err := h.amloService.GetWlStatus(ID)
		if err != nil {
			h.logger.Errorf("[Handler] Get status watchlist has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500098,
				Message: "Status watchlist has error",
			})
		}
		result = dataWlInsertStatus

	case "tf":
		dataWlInsertStatus, err := h.amloService.GetTFStatus(ID)
		if err != nil {
			h.logger.Errorf("[Handler] Get status transaction factor has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500098,
				Message: "Status transaction factor has error",
			})
		}
		result = dataWlInsertStatus

	default:
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Error", Data: "This import csv has error"})
	}
	return c.JSON(http.StatusOK, model.Response{Message: "Success", Data: result})
}
