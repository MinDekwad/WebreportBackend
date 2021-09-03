package router

import (
	//"go-api-report2/handler"
	"go-api-report2/handler/daily"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	DailyRouter struct {
		Server  *echo.Echo
		Handler daily.Handler
		Mid     middleware.IMiddleware
	}
)

func (p DailyRouter) NewRouter() {
	h := p.Handler
	m := p.Mid

	reportGroup := p.Server.Group("/api/v1/reports")
	reportGroup.GET("/:slug", h.GetReports, m.CheckAccessTokenPermission("dailyReport", "READ"), m.WriteLog)
}
