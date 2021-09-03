package services

import (
	"go-api-report2/ent"
	"go-api-report2/log"
	"go-api-report2/model"
)

var (
	channel        chan model.Reports
	triggerChannel chan bool
	//pointTransactionChannel chan model.TransactionPoint // new
)

type ReportService struct {
	connection *ent.Client
	logger     *log.Logger
}

func NewReportService(connection *ent.Client, logger *log.Logger) (ReportService, error) {

	channel = make(chan model.Reports, 0)
	triggerChannel = make(chan bool, 0)
	//pointTransactionChannel = make(chan model.TransactionPoint, 0) // new

	return ReportService{
		connection: connection,
		logger:     logger,
	}, nil
}
