package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var RootConfig *AppSettings

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	RootConfig = NewApplicationConfig()
}

// AppSettings struct of root app config
type AppSettings struct {
	Name     string       `mapstructure:"name"`
	Server   ServerDetail `mapstructure:"server"`
	Database ServerDetail `mapstructure:"database"`
	OAuth2   Detail       `mapstructure:"oauth2"`
	Admin    Admin        `mapstructure:"admin"`
	Billpay  Billpay      `mapstructure:"billpay"`
	Log      Log          `mapstructure:"log"`
}

// Log struct of log instance config
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// ServerDetail struct of
type ServerDetail struct {
	Port        int    `mapstructure:"port"`
	SSL         bool   `mapstructure:"ssl"`
	Key         string `mapstructure:"key"`
	Cert        string `mapstructure:"cert"`
	APIKey      string `mapstructure:"apikey"`
	RootCA      string `mapstructure:"rootca"`
	Endpoint    string `mapstructure:"endpoint"`
	Path        string `mapstructure:"path"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	Driver      string `mapstructure:"driver"`
	AutoMigrate bool   `mapstructure:"auto-migrate"`
	EndpointWeb string `mapstructure:"endpoint_web"`
}

// Detail struct of service detail
type Detail struct {
	EndpointWeb string `mapstructure:"endpoint_web"`
	Endpoint    string `mapstructure:"endpoint"`
	Path        string `mapstructure:"path"`
	Protocol    string `mapstructure:"protocol"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	APIKey      string `mapstructure:"apikey"`
	AutoMigrate bool   `mapstructure:"auto-migrate"`
	RedirectURI string `mapstructure:"redirect-uri"`
	Scope       string `mapstructure:"scope"`
	SSLVerify   bool   `mapstructure:"ssl-verify"`
}

type Admin struct {
	APIKey string `mapstructure:"apikey"`
}

type Billpay struct {
	Endpoint  string `mapstructure:"endpoint"`
	APIKey    string `mapstructure:"apikey"`
	CACert    string `mapstructure:"ca-cert"`
	SSLVerify bool   `mapstructure:"ssl-verify"`
}

// NewApplicationConfig new instance of App config
func NewApplicationConfig() *AppSettings {
	conf := AppSettings{}

	p := os.Getenv("REPORT_DB_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		panic(err)
	}

	conf.Database = ServerDetail{
		Endpoint: os.Getenv("REPORT_DB_HOST"),
		Port:     port,
		Username: os.Getenv("REPORT_DB_USER"),
		Password: os.Getenv("REPORT_DB_PASSWORD"),
		Name:     os.Getenv("REPORT_DB_NAME"),
	}

	b := os.Getenv("REPORT_OAUTH2_VERIFY_SSL")
	verify := false
	if b == "true" || b == "t" || b == "" {
		verify = true
	}

	conf.OAuth2 = Detail{
		Endpoint:    os.Getenv("REPORT_OAUTH2_BACKEND_URI"),
		EndpointWeb: os.Getenv("REPORT_OAUTH2_WEB_URI"),
		Username:    os.Getenv("REPORT_CLIENT_ID"),
		Password:    os.Getenv("REPORT_CLIENT_SECRET"),
		Scope:       os.Getenv("REPORT_OAUTH2_SCOPE"),
		RedirectURI: os.Getenv("REPORT_REDIRECT_URI"),
		Name:        os.Getenv("REPORT_OAUTH2_SERVICE_NAME"),
		SSLVerify:   verify,
	}

	conf.Admin = Admin{
		APIKey: os.Getenv("REPORT_OAUTH2_API"),
	}

	bo := os.Getenv("REPORT_BILLPAY_VERIFY_SSL")
	verify2 := false
	if bo == "true" || bo == "t" || bo == "" {
		verify2 = true
	}

	conf.Billpay = Billpay{
		Endpoint:  os.Getenv("REPORT_BILLPAY_URI"),
		APIKey:    os.Getenv("REPORT_BILLPAY_APIKEY"),
		CACert:    os.Getenv("REPORT_BILLPAY_CACERT"),
		SSLVerify: verify2,
	}

	// Log
	conf.Log = Log{
		Level:  os.Getenv("REPORT_LOG_LEVEL"),
		Format: os.Getenv("REPORT_LOG_FORMAT"),
	}

	return &conf
}
