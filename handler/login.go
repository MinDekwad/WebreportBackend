package handler

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-api-report2/model"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetLogin(c echo.Context) error {
	state := c.QueryParam("state")
	conf := h.conf.OAuth2
	clientID := conf.Username
	scope := conf.Scope
	redirectURI := conf.RedirectURI
	endpointWeb := conf.EndpointWeb

	u := url.Values{}
	u.Set("scope", scope)
	u.Set("redirect_uri", redirectURI)
	u.Set("state", state)
	u.Set("client_id", clientID)
	u.Set("response_type", "code")

	url := endpointWeb + "/oauth2/authorize?" + u.Encode()

	return c.JSON(http.StatusOK, model.Response{Message: "success", Url: url})
}

func (h Handler) PostLogin(c echo.Context) error {
	m := make(map[string]string, 0)
	if err := c.Bind(&m); err != nil {
		h.logger.Errorf("[Handler] Post login binding error : %s", err)
		return c.JSON(http.StatusBadRequest, model.BadRequest)
	}

	conf := h.conf.OAuth2

	clientID := conf.Username
	clientSecret := conf.Password
	oauth2URI := conf.Endpoint
	sslverify := conf.SSLVerify
	redirectURI := conf.RedirectURI

	m["grant_type"] = "authorization_code"
	m["redirect_uri"] = redirectURI
	m["client_id"] = clientID

	bytes, err := json.Marshal(m)
	if err != nil {
		h.logger.Errorf("[Handler] Post login json marchal error : %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500001,
			Message: "Error convert json",
		})
	}

	client := resty.New()
	client.SetContentLength(true)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(clientID, clientSecret).
		SetBody(bytes).
		Post(oauth2URI + "/oauth2/token")

	if err != nil {
		h.logger.Errorf("[Handler] Post login exchange toker has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500002,
			Message: "Exchange token has error",
		})
	}
	if !resp.IsSuccess() {
		return c.JSON(http.StatusBadRequest, model.Response{Message: fmt.Sprintf("login statusCode : %d", resp.StatusCode())})
	}

	data := ResponseToken{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		h.logger.Error("[Handler] Post login reponse token unmarshal has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500003,
			Message: "Error ResponseToken",
		})
	}
	return c.JSON(http.StatusOK, data)
}
