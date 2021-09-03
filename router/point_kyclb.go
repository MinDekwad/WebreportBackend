package router

import (
	"go-api-report2/handler"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	PointKYCLBRouter struct {
		Server  *echo.Echo
		Handler handler.Handler
		Mid     middleware.IMiddleware
	}
)

func (p PointKYCLBRouter) NewRouter() {
	h := p.Handler
	//m := p.Mid
	// TODO : Add middleware check permission and writelog
	pointGroup := p.Server.Group("/api/v1/reports/pointkyclb")
	pointGroup.POST("/calculatePointKYC", h.CreateCalculatePointKYC)
	pointGroup.POST("/calculatePointLB", h.CreateCalculatePointLB)
	pointGroup.GET("/createPointTransacKYCCSV", h.CreateReportPointTransacKCYCSV)
	pointGroup.GET("/exportPointKycCsv", h.GetExportPointKycCsv)
	pointGroup.GET("/createPointTransacLBCSV", h.CreateReportPointTransacLBCSV)
	pointGroup.GET("/exportPointLbCsv", h.GetExportPointLbCsv)
	pointGroup.GET("/pointPendingKYCList", h.GetPendingKycList)
	pointGroup.GET("/pointPendingKYCListSearch", h.GetPendingKYCListSearch)
	pointGroup.GET("/fetchDataPendingLBList", h.GetPendingLbList)
	pointGroup.GET("/pointPendingLBListSearch", h.GetPendingLBListSearch)
}
