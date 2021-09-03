package handler

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreateCalculatePointKYC(c echo.Context) error {
	date := c.QueryParam("date")
	//var result interface{}

	calculatePointKYC, err := h.reportService.CalculatePointKYC(date)
	if err != nil {
		h.logger.Errorf("[Handler] Calculate point kyc has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500197,
			Message: "Calculate point kyc has rrror",
		})
	}

	//result = calculatePointKYC
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: calculatePointKYC})
}

func (h Handler) CreateCalculatePointLB(c echo.Context) error {
	date := c.QueryParam("date")
	//var result interface{}

	calculatePointLB, err := h.reportService.CalculatePointLB(date)
	if err != nil {
		h.logger.Errorf("[Handler] Calculate point loanbinding has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500198,
			Message: "Calculate point loanbinding has rrror",
		})
	}

	//result = calculatePointKYC
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: calculatePointLB})
}

func (h Handler) CreateReportPointTransacKCYCSV(c echo.Context) error {

	date := c.QueryParam("date")
	note := c.QueryParam("note")

	var result interface{}

	createPointTransacKYCCSV, err := h.reportService.CreatePointTransacKYCCSV(date, note)
	//createPointKYCCSV, err := h.reportService.CreatePointKYCCSV(date)
	if err != nil {
		h.logger.Errorf("[Handler] func create point kyc csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500123,
			Message: "Create point kyc csv has error",
		})
	}

	result = createPointTransacKYCCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetExportPointKycCsv(c echo.Context) error {
	date := c.QueryParam("date")

	var result interface{}

	exportPointKycCSV, err := h.reportService.ExportPointKycCSV(date)
	if err != nil {
		h.logger.Errorf("[Handler] func export point kyc csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500124,
			Message: "Export point kyc csv has error",
		})
	}

	result = exportPointKycCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) CreateReportPointTransacLBCSV(c echo.Context) error {

	date := c.QueryParam("date")
	note := c.QueryParam("note")

	var result interface{}

	createPointTransacLBCSV, err := h.reportService.CreatePointTransacLBCSV(date, note)
	//createPointKYCCSV, err := h.reportService.CreatePointKYCCSV(date)
	if err != nil {
		h.logger.Errorf("[Handler] func create point loanbinding csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500123,
			Message: "Create point loanbinding csv has error",
		})
	}

	result = createPointTransacLBCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetExportPointLbCsv(c echo.Context) error {
	date := c.QueryParam("date")

	var result interface{}

	exportPointLbCSV, err := h.reportService.ExportPointLbCSV(date)
	if err != nil {
		h.logger.Errorf("[Handler] func export point loanbinding csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500124,
			Message: "Export point loanbinding csv has error",
		})
	}

	result = exportPointLbCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetPendingKycList(c echo.Context) error {

	data, err := h.reportService.GetAllPendingKycList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all pending kyc list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500199,
			Message: "Get all pending kyc  list has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: data})
}

func (h Handler) GetPendingKYCListSearch(c echo.Context) error {

	walletID := c.QueryParam("walletID")
	dateGen := c.QueryParam("dategen")

	data, err := h.reportService.GetAllPendingKYCListSearch(walletID, dateGen)
	if err != nil {
		h.logger.Errorf("[Handler] Get search all pending kyc list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500200,
			Message: "Get search all pending kyc list has error",
		})
	}
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: data})
}

func (h Handler) GetPendingLbList(c echo.Context) error {

	data, err := h.reportService.GetAllPendingLbList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all pending loan binding list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500201,
			Message: "Get all pending loan binding list has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: data})
}

func (h Handler) GetPendingLBListSearch(c echo.Context) error {

	walletID := c.QueryParam("walletID")
	dateGen := c.QueryParam("dategen")

	data, err := h.reportService.GetAllPendingLBListSearch(walletID, dateGen)
	if err != nil {
		h.logger.Errorf("[Handler] Get search all pending loanbinding list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500201,
			Message: "Get search all pending loanbinding list has error",
		})
	}
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: data})
}
