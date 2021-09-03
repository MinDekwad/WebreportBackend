package middleware

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"go-api-report2/config"
	"go-api-report2/log"
	"go-api-report2/model"
	"go-api-report2/services"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	layoutISO    = "2006-01-02"
	layoutExport = "2006_01_02"
)

type ScopeConfig struct {
	Skipper middleware.Skipper
	Service string
	Scope   string
}

// IMiddleware interface of middleware
type IMiddleware interface {
	CheckAccessTokenPermission(resource, action string) echo.MiddlewareFunc
	WriteLog(next echo.HandlerFunc) echo.HandlerFunc
	WriteLogExportPoint(next echo.HandlerFunc) echo.HandlerFunc
}

// Middleware implementor of interface middleware
type Middleware struct {
	logger      *log.Logger
	oauthConfig config.Detail
	services    services.ReportService
}

// NewMiddleware create instance of IMiddleware
func NewMiddleware(conf *config.AppSettings, logger *log.Logger, services services.ReportService) IMiddleware {
	return Middleware{
		oauthConfig: conf.OAuth2,
		services:    services,
		logger:      logger,
	}
}

func (mid Middleware) WriteLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc, ok := c.(*model.CustomContext)
		if !ok {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500014,
				Message: "Convert custom context error",
			})
		}
		// mid.logger.Debugln("[Middleware] Get adminId from custom context", cc.AdminID, cc.PermissionResponse.Resource, cc.PermissionResponse.Action)
		err := mid.services.CreateLog(cc.AdminID, cc.PermissionResponse.Resource, cc.PermissionResponse.Action)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500108,
				Message: "Insert log error",
			})
		}

		return next(c)
	}
}

func (mid Middleware) WriteLogExportPoint(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc, ok := c.(*model.CustomContext)
		if !ok {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500014,
				Message: "Convert custom context error",
			})
		}
		mid.logger.Debugln("[Middleware] Get adminId from custom context", cc.AdminID, cc.PermissionResponse.Resource, cc.PermissionResponse.Action)

		stdate := c.QueryParam("start_date")
		end_date := c.QueryParam("end_date")
		user := c.QueryParam("user")
		t, _ := time.Parse(layoutISO, stdate)
		start_date := t.Format(layoutExport)
		// err := mid.services.CreateLogExportPointCSV(cc.AdminID, cc.PermissionResponse.Resource, cc.PermissionResponse.Action, start_date, end_date)
		err := mid.services.CreateLogExportPointCSV(user, cc.AdminID, cc.PermissionResponse.Resource, cc.PermissionResponse.Action, start_date, end_date)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500128,
				Message: "Insert log export point error",
			})
		}

		return next(c)
	}
}

// CheckAccessTokenPermission do
func (mid Middleware) CheckAccessTokenPermission(resource, action string) echo.MiddlewareFunc {
	config := ScopeConfig{}
	if config.Skipper == nil {
		config.Skipper = func(echo.Context) bool {
			return false
		}
	}

	if config.Skipper == nil {
		config.Skipper = func(echo.Context) bool {
			return false
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return c.JSON(http.StatusUnauthorized, model.Response{Message: "authorize is required"})
			}

			slice := strings.Split(auth, "Bearer ")
			if len(slice) < 1 || slice[len(slice)-1] == "" {
				return c.JSON(http.StatusUnauthorized, model.Response{Message: "authorize is required."})
			}

			accessToken := slice[len(slice)-1]

			reqcheck := model.OAuthPermission{
				Service:  mid.oauthConfig.Name,
				Resource: resource,
				Action:   action,
			}

			payload, err := json.Marshal(reqcheck)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
			}

			permission, err := mid.checkPermission(accessToken, payload)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
			}

			if !permission.Allow {
				return c.JSON(http.StatusUnauthorized, model.Response{Message: "permission denied"})
			}

			cc := &model.CustomContext{
				Context:            c,
				PermissionResponse: permission,
				AdminID:            permission.AdminID,
			}

			return next(cc)
		}
	}
}

func (mid Middleware) checkPermission(token string, payload interface{}) (*model.OAuthPermission, error) {
	client := resty.New()
	client.SetContentLength(true)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !mid.oauthConfig.SSLVerify,
	})
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		SetBody(payload).
		Post(mid.oauthConfig.Endpoint + "/permissions/verifies")

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		if resp.StatusCode() == http.StatusUnauthorized {
			return &model.OAuthPermission{Allow: false}, nil
		}

		return nil, errors.New("check active permission response status " + resp.Status())
	}

	//fmt.Println("[Middleware] debug permission resp", string(resp.Body()))

	data := model.ResponsePermission{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	oauthPermission := data.Data
	oauthPermission.Allow = true

	return &oauthPermission, nil
}
