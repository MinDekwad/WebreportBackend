package log

import (
	"go-api-report2/config"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Logger struct of logrus
type Logger struct {
	*logrus.Logger
}

// NewLogger create log instance
func NewLogger(conf *config.AppSettings) *Logger {
	log := logrus.New()

	switch strings.ToLower(conf.Log.Format) {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
		})
	default:
		log.SetFormatter(&logrus.TextFormatter{FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		}})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example

	log.SetOutput(os.Stdout)

	switch strings.ToLower(conf.Log.Level) {
	case "debug", "":
		log.SetLevel(logrus.DebugLevel)
		log.SetReportCaller(true)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	}

	// Only log the warning severity or above.

	return &Logger{log}
}
