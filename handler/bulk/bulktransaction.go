package bulk

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetListBulkTransaction(c echo.Context) error {
	var result interface{}
	dataBulkTransaction, err := h.bulkService.GetAllBulkTransaction()
	if err != nil {
		h.logger.Errorf("[Handler] Get report all bulk transaction has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500083,
			Message: "All bulk transaction has error",
		})
	}
	result = dataBulkTransaction
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) CreateReportBulkTransaction(c echo.Context) error {

	var result interface{}
	var req model.CreateBulkTransaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}
	data, err := h.bulkService.CreateBulkTransaction(c.Request().Context(), req)
	if err != nil {
		h.logger.Errorf("[Handler] create bulk transaction has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500009,
			Message: "Create bulk transaction has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})

}

func (h Handler) GetBulkTransaction(c echo.Context) error {

	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	dataEditBulkTransaction, err := h.bulkService.GetEditBulkTransaction(Uid)
	if err != nil {
		h.logger.Errorf("[Handler] Get bulk transaction has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500104,
			Message: "Get bulk transaction has error",
		})
	}
	result = dataEditBulkTransaction

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})

}

func (h Handler) UpdateBulkTransaction(c echo.Context) error {

	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	var req model.EditBulkTransaction
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}
	_, err = h.bulkService.SaveEditBulkTransaction(c.Request().Context(), Uid, req)
	if err != nil {
		h.logger.Errorf("[Handler] Edit bulk transaction has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500106,
			Message: "Edit bulk transaction has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})

}
