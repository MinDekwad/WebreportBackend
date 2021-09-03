package statement

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetListStatementendingBalance(c echo.Context) error {
	var result interface{}

	dataStatementEndingBalance, err := h.statementService.GetAllStatementEndingBalance()
	if err != nil {
		h.logger.Errorf("[Handler] Get report list statement ending balance has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500082,
			Message: "List statement ending balance has error",
		})
	}
	result = dataStatementEndingBalance
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) CreateReportStatementEndingBalance(c echo.Context) error {

	var result interface{}

	var req model.CreateStatementEndingBalanc
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
	}
	data, err := h.statementService.CreateStatementEndingBalanc(c.Request().Context(), req)
	if err != nil {
		h.logger.Errorf("[Handler] func Create statement ending balance has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500008,
			Message: "Create statement ending balance has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})

}

func (h Handler) GetStatementendingBalance(c echo.Context) error {

	var result interface{}
	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	dataEditStatementEndingBalance, err := h.statementService.GetEditStatementEndingBalance(Uid)
	if err != nil {
		h.logger.Errorf("[Handler] Get statement ending balance has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500103,
			Message: "Get statement ending balance has error",
		})
	}
	result = dataEditStatementEndingBalance

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) UpdateStatementendingBalance(c echo.Context) error {

	uid := c.QueryParam("id")
	Uid, err := strconv.Atoi(uid)

	var req model.EditStatementEndingBalanc
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Request body is invalid1"})
	}
	_, err = h.statementService.SaveEditStatementEndingBalance(c.Request().Context(), Uid, req)
	if err != nil {
		h.logger.Errorf("[Handler] Edit statement ending balance has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500105,
			Message: "Edit statement ending balance has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})

}
