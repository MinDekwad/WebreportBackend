package point

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreateReportPointCSV(c echo.Context) error {

	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")
	note := c.QueryParam("note")

	var result interface{}

	createPointCSV, err := h.pointService.CreatePointCSV(start_date, end_date, note)
	if err != nil {
		h.logger.Errorf("[Handler] func create point csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500123,
			Message: "Create point csv has error",
		})
	}

	result = createPointCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetExportPointCsv(c echo.Context) error {
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")

	var result interface{}

	exportPointCSV, err := h.pointService.ExportPointCSV(start_date, end_date)
	if err != nil {
		h.logger.Errorf("[Handler] func export point csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500124,
			Message: "Export point csv has error",
		})
	}

	result = exportPointCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetDownloadPointCsv(c echo.Context) error {
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")

	var result interface{}

	downloadPointCSV, err := h.pointService.DownloadPointCSV(start_date, end_date)
	if err != nil {
		h.logger.Errorf("[Handler] func download point csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500130,
			Message: "Download point csv has error",
		})
	}

	result = downloadPointCSV
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
