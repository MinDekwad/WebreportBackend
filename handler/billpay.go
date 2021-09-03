package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api-report2/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	timeCustomLayout = "2006-01-02"
	layoutISO        = "2006-01-02 00:00:00"
	layoutISOA       = "2006-01-02 15:04:05"
)

func (h Handler) GetBillpayTransaction(c echo.Context) error {
	req := new(model.RequestBillpay)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "request is invalid"})
	}

	apiKey := h.conf.Billpay.APIKey
	endpoint := h.conf.Billpay.Endpoint

	resp, err := h.Restry.R().
		SetQueryParams(map[string]string{
			"start_date":   req.StartDate,
			"end_date":     req.EndDate,
			"limit":        fmt.Sprintf("%d", req.Limit),
			"status":       req.Status,
			"channel_id":   fmt.Sprintf("%d", req.ChannelID),
			"reference_id": req.ReferenceID,
			"page":         fmt.Sprintf("%d", req.Page),
		}).
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", apiKey).
		Get(endpoint + "/api/v1/admins/transactions")
	if err != nil {
		h.logger.Errorf("[Handle] Get billpay transaction error : %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500004,
			Message: "Set query parameter billpay transaction has error",
		})
	}

	if !resp.IsSuccess() {
		return c.JSON(http.StatusBadRequest, model.Response{Message: fmt.Sprintf("Data has error : %d", resp.StatusCode())})
	}

	data := make(map[string]interface{}, 0)
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		h.logger.Errorf("[Handler] Get billpay transaction map string body has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500005,
			Message: "Billpay transaction map string body has error",
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h Handler) GetBillpaySummary(c echo.Context) error {
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")
	field := c.QueryParam("field")

	if start_date == "" {
		return errors.New("Start Date is required")
	}

	if end_date == "" {
		return errors.New("End Date is required")
	}

	st, err := time.Parse(timeCustomLayout, start_date)
	if err != nil {
		return err
	}
	stdate := st.Format(layoutISO)

	e, err := time.Parse(timeCustomLayout, end_date)
	if err != nil {
		return err
	}
	endDate := e.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	enddate := endDate.Format(layoutISOA)

	req := new(model.RequestBillpay)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{Message: "request is invalid"})
	}

	apiKey := h.conf.Billpay.APIKey
	endpoint := h.conf.Billpay.Endpoint

	resp, err := h.Restry.R().
		SetQueryParams(map[string]string{
			"start_date": stdate,
			"end_date":   enddate,
			"field":      field,
		}).
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", apiKey).
		Get(endpoint + "/api/v1/admins/summaries")
	if err != nil {
		h.logger.Errorf("[Handler] Get billpay summary error : %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500006,
			Message: "Set query parameter billpay summary has error",
		})
	}

	if !resp.IsSuccess() {
		return c.JSON(http.StatusBadRequest, model.Response{Message: fmt.Sprintf("Data has error : %d", resp.StatusCode())})
	}

	data := make(map[string]interface{}, 0)
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		h.logger.Errorf("[Handler] Get billpay summary map string body has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500007,
			Message: "Billpay summary map string body has error",
		})
	}

	return c.JSON(http.StatusOK, data)
}
