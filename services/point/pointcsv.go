package point

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/limitpoint"
	"go-api-report2/ent/pointcsv"
	"go-api-report2/ent/pointtransaction"
	"go-api-report2/model"
	"time"
)

func (r PointService) CreatePointCSV(start_date, end_date, note string) ([]model.PointExportCsv, error) {

	if start_date == "" || end_date == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, start_date)
	if err != nil {
		return nil, err
	}
	enDate, err := time.Parse(dateTimeCustom, end_date)
	endDate := enDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	lp, err := r.connection.Limitpoint.Query().Select(limitpoint.FieldLimitPoint).First(context.Background())
	if err != nil {
		return nil, err
	}
	limitPoint := lp.LimitPoint

	var restotal []model.PointExportCsv
	err = r.connection.Pointtransaction.Query().
		Select(pointtransaction.FieldWalletID).
		Where(
			pointtransaction.DateGTE(stDate),
			pointtransaction.DateLTE(endDate),
		).
		GroupBy(pointtransaction.FieldWalletID).
		Aggregate(ent.Sum(pointtransaction.FieldPoint)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return nil, err
	}

	countPointCsv, err := r.connection.Pointcsv.Query().
		Where(
			pointcsv.PointTranDateGTE(stDate),
			pointcsv.PointTranDateLTE(endDate),
		).Count(context.Background())
	if err != nil {
		return nil, err
	}

	//new query
	if countPointCsv > 0 {
		_, err := r.connection.Pointcsv.
			Update().
			Where(
				pointcsv.PointTranDateGTE(stDate),
				pointcsv.PointTranDateLTE(endDate),
			).
			SetActionExport(0).
			Save(context.Background())

		if err != nil {
			return nil, err
		}
	}

	for _, pointTran := range restotal {
		if pointTran.Sum > limitPoint {
			pointTran.Sum = limitPoint
		}
		_, err := r.connection.Pointcsv.
			Create().
			SetWalletID(pointTran.WalletID).
			SetCreateDate(time.Now()).
			SetAdjustamount(pointTran.Sum).
			SetNillableNote(&note).
			SetPointTranDate(stDate).
			SetActionExport(1).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return nil, err
}

func (r PointService) ExportPointCSV(start_date, end_date string) (interface{}, error) {
	if start_date == "" || end_date == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, start_date)
	if err != nil {
		return nil, err
	}
	enDate, err := time.Parse(dateTimeCustom, end_date)
	endDate := enDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	pc, err := r.connection.Pointcsv.Query().
		Select(pointcsv.FieldWalletID, pointcsv.FieldAdjustamount, pointcsv.FieldNote).
		Where(
			pointcsv.PointTranDateGTE(stDate),
			pointcsv.PointTranDateLTE(endDate),
			pointcsv.ActionExport(1),
		).
		Order(ent.Desc(pointcsv.FieldAdjustamount)).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return pc, nil
}

func (r PointService) DownloadPointCSV(start_date, end_date string) (interface{}, error) {
	if start_date == "" || end_date == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, start_date)
	if err != nil {
		return nil, err
	}
	enDate, err := time.Parse(dateTimeCustom, end_date)
	endDate := enDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	pt, err := r.connection.Pointtransaction.Query().
		Where(
			pointtransaction.DateGTE(stDate),
			pointtransaction.DateLTE(endDate),
		).
		Order(ent.Desc(pointtransaction.FieldPoint)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return pt, nil
}
