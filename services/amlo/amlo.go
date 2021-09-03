package amlo

import (
	"go-api-report2/ent"
	"go-api-report2/log"
	"go-api-report2/model"
	"math"
)

var (
	channel        chan model.Reports
	triggerChannel chan bool
)

const (
	timeCustomLayout = "2006-01-02"
	layoutISO        = "2006-01-02 00:00:00"
	layoutISOA       = "2006-01-02 15:04:05"
	dateTimeCustom   = "2006-01-02"
)

type AmloService struct {
	connection *ent.Client
	logger     *log.Logger
}

func NewService(connection *ent.Client, logger *log.Logger) (AmloService, error) {

	channel = make(chan model.Reports, 0)
	triggerChannel = make(chan bool, 0)

	return AmloService{
		connection: connection,
		logger:     logger,
	}, nil
}

func (reportService AmloService) getPage(count int) (int, int) {
	perPage := 2000
	result := float64(count) / float64(perPage)
	page := int(math.Ceil(result))

	return page, perPage
}
