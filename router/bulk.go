package router

import (
	// "go-api-report2/handler"
	"go-api-report2/handler/bulk"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	BulkRouter struct {
		Server *echo.Echo
		// Handler handler.Handler
		Handler bulk.Handler
		Mid     middleware.IMiddleware
	}
)

func (p BulkRouter) NewRouter() {
	h := p.Handler
	m := p.Mid

	bulkTransactionGroup := p.Server.Group("/api/v1/reports/bulk")
	bulkTransactionGroup.GET("/listBulkTransaction", h.GetListBulkTransaction, m.CheckAccessTokenPermission("bulkTransaction", "READ"), m.WriteLog)
	bulkTransactionGroup.POST("/createBulkTransaction", h.CreateReportBulkTransaction, m.CheckAccessTokenPermission("bulkTransaction", "CREATE"), m.WriteLog)
	bulkTransactionGroup.GET("/bulktransaction", h.GetBulkTransaction, m.CheckAccessTokenPermission("bulkTransaction", "READ"), m.WriteLog)
	bulkTransactionGroup.PUT("/saveEditBulkTransaction", h.UpdateBulkTransaction, m.CheckAccessTokenPermission("bulkTransaction", "UPDATE"), m.WriteLog)
}
