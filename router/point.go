package router

import (
	"go-api-report2/handler/point"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	PointRouter struct {
		Server *echo.Echo
		//Handler handler.Handler
		Handler point.Handler
		Mid     middleware.IMiddleware
	}
)

func (p PointRouter) NewRouter() {
	h := p.Handler
	m := p.Mid
	pointGroup := p.Server.Group("/api/v1/reports/point")
	pointGroup.GET("/transactionType", h.GetListtransactionType)
	pointGroup.GET("/paymentChannel", h.GetListPaymentChannel)
	pointGroup.GET("/paymentType", h.GetListPaymentType)
	pointGroup.POST("/createConfigPoint", h.CreateReportConfigPoint, m.CheckAccessTokenPermission("pointTransactionCreate", "CREATE"), m.WriteLog)
	pointGroup.GET("/pointTransactionList", h.GetPointTransactionList, m.CheckAccessTokenPermission("pointTransactionList", "READ"), m.WriteLog)
	pointGroup.GET("/pointTransaction", h.GetPointTransaction, m.CheckAccessTokenPermission("pointTransaction", "READ"), m.WriteLog)
	pointGroup.PUT("/editConfigPoint", h.UpdateConfigPoint, m.CheckAccessTokenPermission("pointTransactionEdit", "UPDATE"), m.WriteLog)
	pointGroup.GET("/limitPoint", h.GetLimitPoint, m.CheckAccessTokenPermission("pointLimit", "READ"), m.WriteLog)
	pointGroup.PUT("/editLimitPoint", h.UpdateLimitPoint, m.CheckAccessTokenPermission("pointLimit", "UPDATE"), m.WriteLog)
	pointGroup.POST("/calculatePoint", h.CreatePoint, m.CheckAccessTokenPermission("pointCalculate", "CREATE"), m.WriteLog)
	pointGroup.GET("/historyPoint", h.GetListHistoryPoint, m.CheckAccessTokenPermission("pointHistory", "READ"), m.WriteLog)
	pointGroup.GET("/createPointCSV", h.CreateReportPointCSV, m.CheckAccessTokenPermission("exportPointCSV", "EXPORT"), m.WriteLogExportPoint)
	pointGroup.GET("/exportPointCSV", h.GetExportPointCsv)
	pointGroup.GET("/historyPointExport", h.GetListHistoryPointExport, m.CheckAccessTokenPermission("pointHistoryExport", "READ"), m.WriteLog)
	pointGroup.GET("/downPointCSV", h.GetDownloadPointCsv)
}
