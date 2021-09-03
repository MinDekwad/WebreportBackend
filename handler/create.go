package handler

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreateReports(c echo.Context) error {

	var result interface{}
	slug := c.Param("slug")
	filename := c.QueryParam("filename")

	switch slug {

	case "createWalletFileName":
		dataCreateFileImportWalletdaily, err := h.reportService.CreateFileImport(filename, "Wallet_daily")
		if err != nil {
			h.logger.Errorf("[Handler] create file import wallet daily has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500010,
				Message: "Create file import wallet daily has error",
			})
		}
		result = dataCreateFileImportWalletdaily

	case "createConsumerFileName":
		dataCreateFileImportConsumer, err := h.reportService.CreateFileImport(filename, "consumer")
		if err != nil {
			h.logger.Errorf("[Handler] create file import consumer transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500011,
				Message: "Create file import consumer transaction has error",
			})
		}
		result = dataCreateFileImportConsumer

	default:
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Error", Data: "This report type is not supported"})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
