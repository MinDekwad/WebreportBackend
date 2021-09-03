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

func (h Handler) GetUserinfo(c echo.Context) error {
	sslverify := h.conf.OAuth2.SSLVerify
	oauth2URI := h.conf.OAuth2.Endpoint
	//fmt.Println("oauth2URI : ", oauth2URI)

	client := resty.New()
	client.SetContentLength(true)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})
	token := c.Request().Header.Get("Authorization")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get(oauth2URI + "/oauth2/userinfo")
	if err != nil {
		h.logger.Errorf("[Handler] Get user info has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500999,
			Message: "Get user info has error",
		})
	}

	if !resp.IsSuccess() {
		return c.JSON(http.StatusBadRequest, model.Response{Message: fmt.Sprintf("login statusCode : %d", resp.StatusCode())})
	}

	data := make(map[string]interface{}, 0)
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		h.logger.Errorf("[Handler] Get user info json unmarshal error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500888,
			Message: "Error convert json of get user info",
		})
	}
	//fmt.Println(resp)
	//return nil
	return c.JSON(http.StatusOK, data)
}
