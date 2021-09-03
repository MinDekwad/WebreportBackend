package amlo

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) PutCreateTransactionFactorItemTmp(c echo.Context) error {

	var result interface{}
	var req model.TransactionFactorItem
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}

	data, err := h.amloService.CreateTransactionFactorItemTmp(c.Request().Context(), req)
	if err != nil {
		h.logger.Errorf("[Handler] func create transaction factor item tmp has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500163,
			Message: "Create transaction factor item tmp has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetdataTrasacItemTmp(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllTrasacItemTmp()
	if err != nil {
		h.logger.Errorf("[Handler] Get all transaction factor tmp has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500164,
			Message: "Get all transaction factor tmp has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) PostCreateTransactionFactor(c echo.Context) error {

	var result interface{}
	var req model.TransactionFactor
	//today := c.QueryParam("date")
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}

	data, err := h.amloService.CreateTransactionFactor(c.Request().Context(), req)
	if err != nil {
		h.logger.Errorf("[Handler] func create transaction factor has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500165,
			Message: "Create transaction factor has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetTransactionfactorList(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllTransactionfactorList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all transaction factor list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500166,
			Message: "Get all transaction factor list has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetTransactionfactor(c echo.Context) error {
	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	data, err := h.amloService.GetDataTransactionfactor(Uid)
	if err != nil {
		h.logger.Errorf("[Handler] Get data transaction factor has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500167,
			Message: "Get data transaction factor  has error",
		})
	}
	result = data

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetdataTrasacItem(c echo.Context) error {
	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	data, err := h.amloService.GetAllTrasacItem(Uid)
	if err != nil {
		h.logger.Errorf("[Handler] Get all transaction factor has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500168,
			Message: "Get all transaction factor has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) PutcreateTransactionFactorItem(c echo.Context) error {

	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)
	var req model.TransactionFactorItem
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}

	data, err := h.amloService.CreateTransactionFactorItem(c.Request().Context(), req, Uid)
	if err != nil {
		h.logger.Errorf("[Handler] func create transaction factor item has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500169,
			Message: "Create transaction factor item has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateTransactionFactor(c echo.Context) error {
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	updateDate := c.QueryParam("date")

	var req model.TransactionFactor
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}

	_, err = h.amloService.SaveEditTransactionFactor(c.Request().Context(), Uid, req, updateDate)
	if err != nil {
		h.logger.Errorf("[Handler] Edit transaction factor has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500170,
			Message: "Edit transaction factor has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) GetTransactionItemPer(c echo.Context) error {
	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)
	data, err := h.amloService.GetTransactionItemPer(Uid)
	if err != nil {
		h.logger.Errorf("[Handler] Get transaction item per has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500171,
			Message: "Get transaction item per has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateTransactionFactorItemPer(c echo.Context) error {
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	var req model.TransactionFactorItem
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}

	_, err = h.amloService.SaveEditTransactionFactorItemPer(c.Request().Context(), Uid, req)
	if err != nil {
		h.logger.Errorf("[Handler] Edit transactin factor item per has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500172,
			Message: "Edit transactin factor item per has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) DelTransactionFactorItemPer(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveDelTransactionFactorItemPer(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Delete transaction factor item per has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500173,
			Message: "Delete transaction factor item per has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) DelTransactionFactor(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveDelTransactionFactor(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Delete transaction factor has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500174,
			Message: "Delete transaction factor has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) DelTransactionFactorItemTmpPer(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveDelTransactionFactorItemTmpPer(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Delete transaction factor item tmp per has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500175,
			Message: "Delete transaction factor item tmp per has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) ApproveTransaction(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveApproveTransaction(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Approve transaction has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500182,
			Message: "Approve transaction has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}

func (h Handler) GetTransactionWaitingApprove(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetTransactionListWaitingApprove()
	if err != nil {
		h.logger.Errorf("[Handler] Get transaction list waiting approve has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500204,
			Message: "Get transaction list waiting approve has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
