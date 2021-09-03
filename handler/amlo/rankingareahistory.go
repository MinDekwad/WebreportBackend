package amlo

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetRankingAreaHistory(c echo.Context) error {
	var result interface{}
	walletID := c.QueryParam("walletID")

	data, err := h.amloService.GetDataRankingAreaHistory(walletID)
	if err != nil {
		h.logger.Errorf("[Handler] Get data ranking area history has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500149,
			Message: "Get data ranking area history has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetExportRankingCsv(c echo.Context) error {

	var result interface{}

	downloadPointCSV, err := h.amloService.ExportRankingCSV()
	if err != nil {
		h.logger.Errorf("[Handler] func export ranking csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500151,
			Message: "Export ranking csv has error",
		})
	}

	result = downloadPointCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetExportRanking(c echo.Context) error {

	var result interface{}

	download, err := h.amloService.ExportRanking()
	if err != nil {
		h.logger.Errorf("[Handler] func export ranking has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500205,
			Message: "Export ranking has error",
		})
	}

	result = download
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
