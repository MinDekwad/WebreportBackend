package ekyc

import (
	"crypto/tls"
	"go-api-report2/config"
	"go-api-report2/log"
	"go-api-report2/services/ekyc"

	syslog "log"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var (
	clientid, scope, redirecturi, oauth2uri string
)

type Handler struct {
	ekycService ekyc.EkycService

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

func NewHandler(ekycService ekyc.EkycService, conf *config.AppSettings, logger *log.Logger) (Handler, error) {

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
