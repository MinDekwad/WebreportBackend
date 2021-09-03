package point

import (
	"context"
	"errors"
	"go-api-report2/ent/logexport"
	"time"
)

func (r PointService) GetHistoryPointExport(start_date, end_date string) (interface{}, error) {
	if start_date == "" || end_date == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, start_date)
	if err != nil {
		return nil, err
	}
	enDate, err := time.Parse(dateTimeCustom, end_date)
	endDate := enDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	lxp, err := r.connection.Logexport.Query().
		Where(
			logexport.ExportDateGTE(stDate),
			logexport.ExportDateLTE(endDate),
		).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	if len(lxp) == 0 {
		return nil, err
	}

	return lxp, err
}
