package handler

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-api-report2/model"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetPermission(c echo.Context) error {

	sslverify := h.conf.OAuth2.SSLVerify
	name := h.conf.OAuth2.Name
	oauth2URI := h.conf.OAuth2.Endpoint

	client := resty.New()
	client.SetContentLength(true)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	token := c.Request().Header.Get("Authorization")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		SetQueryParam("service", name).
		Get(oauth2URI + "/permissions/verifies")

	if err != nil {
		h.logger.Errorf("[Handler] Get permission Verifies has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500012,
			Message: "Verifies permission has error",
		})
	}

	if !resp.IsSuccess() {
		return c.JSON(http.StatusBadRequest, model.Response{Message: fmt.Sprintf("login statusCode : %d", resp.StatusCode())})
	}

	data := make(map[string]interface{}, 0)
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		h.logger.Errorf("[Handler] Get permisson json unmarshal error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500013,
			Message: "Error convert json of get permission",
		})
	}

	return c.JSON(http.StatusOK, data)
}
