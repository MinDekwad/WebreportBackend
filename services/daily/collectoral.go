package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/reportwallet"
	"go-api-report2/model"
	"time"
)

// func (reportService ReportService) GetCollectoralIn(date string) (*model.CollectoralIn, error) {
func (reportService DailyService) GetCollectoralIn(date string) (*model.CollectoralIn, error) {
	if date == "" {
		return nil, errors.New("Date is required")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return nil, err
	}
	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	var restotal []model.CollectoralIn
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("ADJUST_IN"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("ADJUST_COL_IN"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return nil, err
	}

	if len(restotal) == 0 {
		return &model.CollectoralIn{}, nil
	}
	result := restotal[0]
	result.SumCollectoralIn = result.Sum
	result.CountCollectoralIn = result.Count

	return &result, nil
}

func (reportService DailyService) GetCollectoralOut(date string) (*model.CollectoralOut, error) {

	if date == "" {
		return nil, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return nil, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.CollectoralOut
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("ADJUST_OUT"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("ADJUST_COL_OUT"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)

	if err != nil {
		return nil, err
	}

	if len(restotal) == 0 {
		return &model.CollectoralOut{}, nil
	}
	result := restotal[0]
	result.SumCollectoralOut = result.Sum
	result.CountCollectoralOut = result.Count
	return &result, nil
}

func (reportService DailyService) GetCollectoralBalance(date string) (interface{}, error) {
	if date == "" {
		return nil, errors.New("Date is required")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return nil, err
	}
	startDatePlus := startDate.AddDate(0, 0, 1)
	res, err := reportService.connection.ReportWallet.Query().Select(reportwallet.FieldBalance).
		Where(
			reportwallet.DateTime(startDatePlus),
			reportwallet.WalletTypeName("COLLATERAL"),
		).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return 0, err
	}
	result := res[0].Balance
	return result, nil
}
