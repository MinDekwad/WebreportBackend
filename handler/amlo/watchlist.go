package amlo

import (
	"go-api-report2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetWatchList(c echo.Context) error {
	var result interface{}

	data, err := h.amloService.GetAllWatchList()
	if err != nil {
		h.logger.Errorf("[Handler] Get all watchlist has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500157,
			Message: "Get all watchlist has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) GetWatchListSearch(c echo.Context) error {
	var result interface{}

	name := c.QueryParam("name")
	taxID := c.QueryParam("taxID")

	data, err := h.amloService.GetAllWatchListSearch(name, taxID)
	if err != nil {
		h.logger.Errorf("[Handler] Get search watchlist by parameter has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500162,
			Message: "Get search watchlist by parameter has error",
		})
	}
	result = data
	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}

func (h Handler) DelWatchlist(c echo.Context) error {

	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)

	_, err = h.amloService.SaveDelWatchlist(c.Request().Context(), ID)
	if err != nil {
		h.logger.Errorf("[Handler] Delete watchlist has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500178,
			Message: "Delete watchlist has error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success"})
}
