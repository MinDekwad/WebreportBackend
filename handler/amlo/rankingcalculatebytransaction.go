package amlo

// func (h Handler) GetCalculateRankingByTransaction(c echo.Context) error {
// 	var result interface{}

// 	data, err := h.reportService.CalculateRankByTransaction()

// 	if err != nil {
// 		h.logger.Errorf("[Handler] Calculate rank by transaction has error %s", err)
// 		return c.JSON(http.StatusInternalServerError, model.Response{
// 			Code:    500179,
// 			Message: "Calculate rank by transaction has error",
// 		})
// 	}
// 	result = data
// 	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
// }
