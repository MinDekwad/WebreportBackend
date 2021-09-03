package router

import (
	"go-api-report2/handler/amlo"
	"go-api-report2/middleware"

	"github.com/labstack/echo/v4"
)

type (
	AmloRouter struct {
		Server  *echo.Echo
		Handler amlo.Handler
		Mid     middleware.IMiddleware
	}
)

func (p AmloRouter) NewRouter() {
	h := p.Handler
	//m := p.Mid
	// TODO : Add middleware check permission and writelog
	amloGroup := p.Server.Group("/api/v1/reports/amlo")
	amloGroup.GET("/occupationList", h.GetOccupationList)
	amloGroup.GET("/occupationListSearch", h.GetOccupationListSearch)
	amloGroup.POST("/delOccupation", h.DelConfigOccupation)
	amloGroup.GET("/occupation", h.GetOccupation)
	amloGroup.POST("/createOccupation", h.CreateConfigOccupation)
	amloGroup.PUT("/editOccupation", h.UpdateConfigOccupation)
	amloGroup.POST("/approveOccupation", h.ApproveConfigOccupation)
	amloGroup.GET("/provinceNameTH", h.GetListProvinceNameTH)
	amloGroup.GET("/areaListSearch", h.GetAreaListSearch)
	amloGroup.GET("/areaList", h.GetAreaList)
	amloGroup.GET("/area", h.GetArea)
	amloGroup.PUT("/editArea", h.UpdateConfigArea)
	amloGroup.POST("/approveArea", h.ApproveConfigArea)
	amloGroup.GET("/dateCalculateRankList", h.GetDateCalculateRankList)
	amloGroup.GET("/dateCalculateRank", h.GetDateCalculateRank)
	amloGroup.PUT("/editDateCalculateRank", h.UpdateConfigDateCalculateRank)
	amloGroup.POST("/approveDateCalculateRank", h.ApproveConfigDateCalculateRank)
	amloGroup.GET("/rankList", h.GetRankList)
	amloGroup.GET("/rankListSearch", h.GetRankListSearch)
	amloGroup.POST("/calculateRank", h.PostCalculateRank) // run auto all day by system, run this second (2) after run GetInsertRanking (1)
	amloGroup.GET("/rankingAreaHistory", h.GetRankingAreaHistory)
	amloGroup.GET("/rankingOccupationHistory", h.GetRankingOccupationHistory)
	amloGroup.GET("/rankingWatchlistHistory", h.GetRankingWatchlistHistory)
	amloGroup.GET("/rankingTransactionHistory", h.GetRankingTransactionHistory)
	amloGroup.GET("/downRankingCSV", h.GetExportRankingCsv)
	amloGroup.GET("/oldcustomerList", h.GetOldCustomerList)
	amloGroup.GET("/insertRanking", h.GetInsertRanking) // run auto all day by system, run this first (1)
	amloGroup.GET("/updateRankingOccupation", h.GetUpdateRankingOccupation)
	amloGroup.GET("/watchlist", h.GetWatchList)
	amloGroup.GET("/watchlistSearch", h.GetWatchListSearch)
	amloGroup.PUT("/createTransactionFactorItemTmp", h.PutCreateTransactionFactorItemTmp)
	amloGroup.GET("/dataTrasacItemTmp", h.GetdataTrasacItemTmp)
	amloGroup.POST("/createTransactionFactor", h.PostCreateTransactionFactor)
	amloGroup.GET("/transactionfactorList", h.GetTransactionfactorList)
	amloGroup.GET("/transactionFactor", h.GetTransactionfactor)
	amloGroup.GET("/dataTrasacItem", h.GetdataTrasacItem)
	amloGroup.PUT("/createTransactionFactorItem", h.PutcreateTransactionFactorItem)
	amloGroup.PUT("/editTransactionFactor", h.UpdateTransactionFactor)
	amloGroup.GET("/transactionItemPer", h.GetTransactionItemPer)
	amloGroup.PUT("/editTransactionFactorItemPer", h.UpdateTransactionFactorItemPer)
	amloGroup.POST("/delTransactionFactorItemPer", h.DelTransactionFactorItemPer)
	amloGroup.POST("/delTransactionFactor", h.DelTransactionFactor)
	amloGroup.POST("/delTransactionFactorItemTmpPer", h.DelTransactionFactorItemTmpPer)
	amloGroup.POST("/delWatchlist", h.DelWatchlist)
	amloGroup.GET("/downloadAmloCustomer", h.GetDownloadAmloCustomer)
	amloGroup.POST("/calculate/:slug", h.Calculate)
	amloGroup.GET("/status/:slug", h.CheckStatusCalculate)
	amloGroup.POST("/approveTransaction", h.ApproveTransaction)
	//amloGroup.GET("/api/v1/reports/amlo/rankingHistory", h.GetRankingHistory)
	amloGroup.GET("/rankingHistory", h.GetRankingHistory)
	amloGroup.GET("/areaListWaitingApprove", h.GetAreaWaitingApprove)
	amloGroup.GET("/occuListWaitingApprove", h.GetOccuWaitingApprove)
	amloGroup.GET("/transactionListWaitingApprove", h.GetTransactionWaitingApprove)
	amloGroup.GET("/downloadRanking", h.GetExportRanking)
}
