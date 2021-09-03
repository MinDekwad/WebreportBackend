package router

import (
	//"go-api-report2/handler"
	"go-api-report2/handler/statement"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	StatementRouter struct {
		Server  *echo.Echo
		Handler statement.Handler
		Mid     middleware.IMiddleware
	}
)

func (p StatementRouter) NewRouter() {
	h := p.Handler
	m := p.Mid

	statementEndingGroup := p.Server.Group("/api/v1/reports/statementending")
	statementEndingGroup.GET("/listStatementEndingBalance", h.GetListStatementendingBalance, m.CheckAccessTokenPermission("statementEndingBalance", "READ"), m.WriteLog)
	statementEndingGroup.POST("/createStatementEndingBalance", h.CreateReportStatementEndingBalance, m.CheckAccessTokenPermission("statementEndingBalance", "CREATE"), m.WriteLog)
	statementEndingGroup.GET("/statementendingbalance", h.GetStatementendingBalance, m.CheckAccessTokenPermission("statementEndingBalance", "READ"), m.WriteLog)
	statementEndingGroup.PUT("/saveEditStatementEndingBalance", h.UpdateStatementendingBalance, m.CheckAccessTokenPermission("statementEndingBalance", "UPDATE"), m.WriteLog)
}
