package amlo

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetRankList(c echo.Context) error {

	p := c.QueryParam("page")
	l := c.QueryParam("limit")

	page, err := strconv.Atoi(p)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(l)
	if err != nil {
		return err
	}

	data, count, err := h.amloService.GetAllRankList(page, limit)
	if err != nil {
		h.logger.Errorf("[Handler] Get all rank list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500147,
			Message: "Get all rank list has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: data, Count: count})
}

func (h Handler) GetRankListSearch(c echo.Context) error {
	//var result interface{}

	walletID := c.QueryParam("walletID")
	name := c.QueryParam("name")
	taxID := c.QueryParam("taxID")

	p := c.QueryParam("page")
	l := c.QueryParam("limit")
	page, err := strconv.Atoi(p)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(l)
	if err != nil {
		return err
	}

	data, count, err := h.amloService.GetAllRankListSearch(walletID, name, taxID, page, limit)
	if err != nil {
		h.logger.Errorf("[Handler] Get search all rank list has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500148,
			Message: "Get search all rank list has error",
		})
	}
	//result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: data, Count: count})
}
