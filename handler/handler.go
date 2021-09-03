package handler

import (
	"crypto/tls"
	"go-api-report2/config"
	"go-api-report2/log"
	"go-api-report2/services"
	"go-api-report2/services/amlo"
	"go-api-report2/services/bulk"
	"go-api-report2/services/daily"
	"go-api-report2/services/ekyc"
	"go-api-report2/services/point"
	"go-api-report2/services/statement"

	syslog "log"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var (
	clientid, scope, redirecturi, oauth2uri string
)

type Handler struct {
	reportService    services.ReportService
	amloService      amlo.AmloService
	bulkService      bulk.BulkService
	statementService statement.StatementService
	dailyService     daily.DailyService
	ekycService      ekyc.EkycService
	pointService     point.PointService

	logger *log.Logger

	conf      *config.AppSettings
	SSLVerify bool
	Endpoint  string
	APIKey    string
	Restry    *resty.Client
}

type ResponseToken struct {
	AccessToken string `json:"access_token,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	IDToken     string `json:"id_token,omitempty"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		syslog.Fatal("Error loading .env file")
	}
}

func NewHandler(reportService services.ReportService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		reportService: reportService,
		conf:          conf,
		Restry:        client,
		logger:        logger,
	}, nil
}

func NewAmloHandler(amloService amlo.AmloService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		amloService: amloService,
		conf:        conf,
		Restry:      client,
		logger:      logger,
	}, nil
}

func NewBulkHandler(bulkService bulk.BulkService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		bulkService: bulkService,
		conf:        conf,
		Restry:      client,
		logger:      logger,
	}, nil
}

func NewStatementHandler(statementService statement.StatementService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		statementService: statementService,
		conf:             conf,
		Restry:           client,
		logger:           logger,
	}, nil
}

func NewDailyHandler(dailyService daily.DailyService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		dailyService: dailyService,
		conf:         conf,
		Restry:       client,
		logger:       logger,
	}, nil
}

func NewEkycHandler(ekycService ekyc.EkycService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		ekycService: ekycService,
		conf:        conf,
		Restry:      client,
		logger:      logger,
	}, nil
}

func NewPointHandler(pointService point.PointService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

	client := resty.New()
	sslverify := conf.Billpay.SSLVerify
	//
	if conf.Billpay.CACert != "" {
		client.SetRootCertificate(conf.Billpay.CACert)
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: !sslverify,
	})

	return Handler{
		pointService: pointService,
		conf:         conf,
		Restry:       client,
		logger:       logger,
	}, nil
}
