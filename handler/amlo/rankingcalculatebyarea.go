package amlo

// func (h Handler) GetCalculateRankingByArea(c echo.Context) error {
// 	var result interface{}

// 	name := c.QueryParam("name")

// 	data, err := h.reportService.CalculateRankByArea(name)

// 	if err != nil {
// 		h.logger.Errorf("[Handler] Calculate rank by area has error %s", err)
// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			Code:    500160,
// 			Message: "Calculate rank by area has error",
// 		})
// 	}
// 	result = data
// 	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
// }
