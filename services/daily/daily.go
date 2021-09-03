package daily

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

type DailyService struct {
	connection *ent.Client
	logger     *log.Logger
}

func NewService(connection *ent.Client, logger *log.Logger) (DailyService, error) {

	channel = make(chan model.Reports, 0)
	triggerChannel = make(chan bool, 0)

	return DailyService{
		connection: connection,
		logger:     logger,
	}, nil
}
