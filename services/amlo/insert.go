package amlo

import (
	"context"
	"go-api-report2/ent"
	"time"
)

func (reportService AmloService) CreateFileInsert(name string) (*ent.Fileinsert, error) {

	return reportService.connection.Fileinsert.Create().
		SetName(name).
		SetImportdate(time.Now()).
		SetIsSuccess(false).
		Save(context.Background())

}
