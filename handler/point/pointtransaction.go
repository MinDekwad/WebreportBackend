package point

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetListtransactionType(c echo.Context) error {
	var result interface{}

	transactionType, err := h.pointService.GetTransactionType()
	if err != nil {
		h.logger.Errorf("[Handler] func get transaction type has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500112,
			Message: "Get transaction type has error",
		})
	}

	result = transactionType
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetListPaymentChannel(c echo.Context) error {
	var result interface{}

	transactionType, err := h.pointService.GetPaymentChannel()
	if err != nil {
		h.logger.Errorf("[Handler] func get payment channel has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500113,
			Message: "Get payment channel has error",
		})
	}

	result = transactionType
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetListPaymentType(c echo.Context) error {
	var result interface{}

	paymentType, err := h.pointService.GetPaymentType()
	if err != nil {
		h.logger.Errorf("[Handler] func get payment type has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500114,
			Message: "Get payment type has error",
		})
	}

	result = paymentType
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) CreateReportConfigPoint(c echo.Context) error {

	var result interface{}
	var req model.CreateConfigPoint
	today := c.QueryParam("date")
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}

	data, err := h.pointService.CreateConfigPoint(c.Request().Context(), req, today)
	if err != nil {
		h.logger.Errorf("[Handler] func create config point has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500115,
			Message: "Create config point has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetPointTransactionList(c echo.Context) error {
	var result interface{}

	data, err := h.pointService.GetAllPointTransactionList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all point transaction list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500116,
			Message: "Get all point transaction list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetPointTransaction(c echo.Context) error {
	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	data, err := h.pointService.GetDataPointTransaction(Uid)
	if err != nil {
		h.logger.Errorf("[Handler] Get data point transaction has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500117,
			Message: "Get data point transaction has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateConfigPoint(c echo.Context) error {
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	updateDate := c.QueryParam("date")

	var req model.EditConfigPoint
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}

	_, err = h.pointService.SaveEditConfigPoint(c.Request().Context(), Uid, req, updateDate)
	if err != nil {
		h.logger.Errorf("[Handler] Edit config point has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500118,
			Message: "Edit config point has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}
