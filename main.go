package main

import (
	"fmt"
	"go-api-report2/config"
	"go-api-report2/db"
	"go-api-report2/handler"
	amloHandler "go-api-report2/handler/amlo"
	bulkHandler "go-api-report2/handler/bulk"
	dailyHandler "go-api-report2/handler/daily"
	ekycHandler "go-api-report2/handler/ekyc"
	pointHandler "go-api-report2/handler/point"
	statementHandler "go-api-report2/handler/statement"
	mid "go-api-report2/middleware"
	"go-api-report2/router"
	"net/http"
	"time"

	"go-api-report2/services"
	"go-api-report2/services/amlo"
	"go-api-report2/services/bulk"
	"go-api-report2/services/daily"
	"go-api-report2/services/ekyc"
	"go-api-report2/services/point"
	"go-api-report2/services/statement"

	"go-api-report2/log"
	syslog "log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port int
)

type Handler struct {
	logger *log.Logger
	conf   *config.AppSettings
}

func init() {
	err := godotenv.Load()
	if err != nil {
		syslog.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("REPORT_SERVER_PORT")
	p, err := strconv.Atoi(port)
	if err != nil {
		syslog.Fatal(err)
	}

	client, err := db.CreateConnection()
	if err != nil {
		syslog.Fatal(err)
	}
	defer client.Close()

	appConf := config.NewApplicationConfig()
	// New Logger
	logger := log.NewLogger(appConf)
	// New Service
	reportService, err := services.NewReportService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	dailyService, err := daily.NewService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	bulkService, err := bulk.NewService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	amloService, err := amlo.NewService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	statementService, err := statement.NewService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	ekycService, err := ekyc.NewService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	pointService, err := point.NewService(client, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	go reportService.InsertCSV()
	go reportService.UpdateReportWallet()

	h, err := handler.NewHandler(reportService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	dailyHandler, err := dailyHandler.NewHandler(dailyService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	bulkHandler, err := bulkHandler.NewHandler(bulkService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	statementHandler, err := statementHandler.NewHandler(statementService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	amloHandler, err := amloHandler.NewHandler(amloService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	ekycHandler, err := ekycHandler.NewHandler(ekycService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	pointHandler, err := pointHandler.NewHandler(pointService, appConf, logger)
	if err != nil {
		syslog.Fatal(err)
	}

	m := mid.NewMiddleware(appConf, logger, reportService)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/api/v1/login", h.GetLogin)
	e.POST("/api/v1/login", h.PostLogin)
	e.GET("/api/v1/permissions", h.GetPermission)
	e.GET("/api/v1/userinfo", h.GetUserinfo)
	// csvGroup := e.Group("/api/v1/reports/csv")
	// csvGroup.GET("/:slug", h.GetImportCsv, m.CheckAccessTokenPermission("importsCSV", "READ"), m.WriteLog)
	// csvGroup.POST("/upload/:slug", h.Upload, m.CheckAccessTokenPermission("importsCSV", "IMPORT"), m.WriteLog)
	e.GET("/api/v1/reports/csv/:slug", h.GetImportCsv)
	e.POST("/api/v1/reports/csv/upload/:slug", h.Upload)
	e.POST("/api/v1/reports/updateKyc", h.UpdateReportKYC, m.CheckAccessTokenPermission("updateReportKYC", "UPDATE"), m.WriteLog)

	// reportGroup := e.Group("/api/v1/reports")
	// reportGroup.GET("/:slug", h.GetReports, m.CheckAccessTokenPermission("dailyReport", "READ"), m.WriteLog)

	dailyRouter := router.DailyRouter{
		Server:  e,
		Handler: dailyHandler,
		Mid:     m,
	}
	dailyRouter.NewRouter()

	statementRouter := router.StatementRouter{
		Server:  e,
		Handler: statementHandler,
		Mid:     m,
	}
	statementRouter.NewRouter()

	bulkRouter := router.BulkRouter{
		Server:  e,
		Handler: bulkHandler,
		Mid:     m,
	}
	bulkRouter.NewRouter()

	ekycRouter := router.EkycRouter{
		Server:  e,
		Handler: ekycHandler,
		Mid:     m,
	}
	ekycRouter.NewRouter()

	billpayRouter := router.BillpayRouter{
		Server:  e,
		Handler: h,
		Mid:     m,
	}
	billpayRouter.NewRouter()

	// pointRouter := router.PointRouter{
	// 	Server:  e,
	// 	Handler: h,
	// 	Mid:     m,
	// }
	// pointRouter.NewRouter()

	pointRouter := router.PointRouter{
		Server:  e,
		Handler: pointHandler,
		Mid:     m,
	}
	pointRouter.NewRouter()

	amloRouter := router.AmloRouter{
		Server:  e,
		Handler: amloHandler,
		Mid:     m,
	}
	amloRouter.NewRouter()

	pointKYCLBRouter := router.PointKYCLBRouter{
		Server:  e,
		Handler: h,
		Mid:     m,
	}
	pointKYCLBRouter.NewRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", p),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	logger.Debugln("[Main] Started server", time.Now())
	e.Logger.Fatal(e.StartServer(s))
}
