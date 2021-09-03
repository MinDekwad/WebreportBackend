package statement

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

type StatementService struct {
	connection *ent.Client
	logger     *log.Logger
}

func NewService(connection *ent.Client, logger *log.Logger) (StatementService, error) {

	channel = make(chan model.Reports, 0)
	triggerChannel = make(chan bool, 0)

	return StatementService{
		connection: connection,
		logger:     logger,
	}, nil
}
