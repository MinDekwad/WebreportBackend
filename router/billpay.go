package router

import (
	"go-api-report2/handler"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	BillpayRouter struct {
		Server  *echo.Echo
		Handler handler.Handler
		Mid     middleware.IMiddleware
	}
)

func (p BillpayRouter) NewRouter() {
	h := p.Handler
	m := p.Mid
	billpayGroup := p.Server.Group("/api/v1/billpay")
	billpayGroup.GET("/transaction", h.GetBillpayTransaction, m.CheckAccessTokenPermission("billpaymentTransaction", "READ"), m.WriteLog)
	billpayGroup.GET("/summary", h.GetBillpaySummary, m.CheckAccessTokenPermission("billpaymentSummary", "READ"), m.WriteLog)
}
