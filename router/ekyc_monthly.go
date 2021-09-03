package router

import (
	//"go-api-report2/handler"
	"go-api-report2/handler/ekyc"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	EkycRouter struct {
		Server  *echo.Echo
		Handler ekyc.Handler
		Mid     middleware.IMiddleware
	}
)

func (p EkycRouter) NewRouter() {
	h := p.Handler
	m := p.Mid

	eKycMonthlyGroup := p.Server.Group("/api/v1/reports/ekyc")
	eKycMonthlyGroup.GET("/eKycMonthly", h.GeteKycMonthly, m.CheckAccessTokenPermission("eKycMonthly", "READ"), m.WriteLog)
	eKycMonthlyGroup.GET("/eKycMonthlyTmdsAll", h.GeteKycMonthlyTmdsAll, m.CheckAccessTokenPermission("eKycMonthly", "READ"), m.WriteLog)
	eKycMonthlyGroup.GET("/eKycMonthlyTcrbAll", h.GeteKycMonthlyTcrbAll, m.CheckAccessTokenPermission("eKycMonthly", "READ"), m.WriteLog)
}
