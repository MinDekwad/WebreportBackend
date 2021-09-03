package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/bulk"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/loanbinding"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetTransferToBankAccountIncome(date string) (model.TransferToBankAccountIncome, error) {
	var result model.TransferToBankAccountIncome
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(8 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(7 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	restotal, err := reportService.connection.Bulk.Query().Select(bulk.FieldBulkCreditSamedayFee, bulk.FieldTransfertobankaccount).
		Where(
			bulk.DateTimeGTE(startDateMinus),
			bulk.DateTimeLTE(endDate),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumTransferToBankAccountIncome = 0
		result.CountTransferToBankAccountIncome = 0
		return result, nil
	}
	result.SumTransferToBankAccountIncome = restotal[0].BulkCreditSamedayFee
	result.CountTransferToBankAccountIncome = restotal[0].Transfertobankaccount

	return result, nil
}

func (reportService DailyService) GetOnlineLoanTopupIncome(date string) (model.OnlineLoanTopupIncome, error) {
	var result model.OnlineLoanTopupIncome
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineLoanTopupIncome
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("LOAN_DIRECT_TCRB"),
			consumer.PaymentType("TOPUP_ONLINE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.CountOnlineLoanTopupIncome = 0
		result.SumOnlineLoanTopupIncome = 0
		return result, nil
	}

	result.CountOnlineLoanTopupIncome = restotal[0].Count
	result.SumOnlineLoanTopupIncome = float64(restotal[0].Count) * 5
	return result, nil
}

func (reportService DailyService) GetOnlineBillPaymentIncome(date string) (model.OnlineBillPaymentIncome, error) {
	var result model.OnlineBillPaymentIncome
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineBillPaymentIncome
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("REPAYMENT"),
			consumer.PaymentChannel("LOAN_DIRECT_TCRB"),
			consumer.PaymentType("TCRB_ONLINE_BILL"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.CountOnlineBillPaymentIncome = 0
		result.SumOnlineBillPaymentIncome = 0
		return result, nil
	}

	result.CountOnlineBillPaymentIncome = restotal[0].Count
	result.SumOnlineBillPaymentIncome = float64(restotal[0].Count) * 5
	return result, nil
}

func (reportService DailyService) GetNewBindingRevolvingIncome(date string) (model.NewBindingRevolvingIncome, error) {
	var result model.NewBindingRevolvingIncome
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.NewBindingRevolvingIncome
	err = reportService.connection.Loanbinding.Query().
		Where(
			loanbinding.DateTimeGTE(startDate),
			loanbinding.DateTimeLTE(endDate),
			loanbinding.Status("APPROVED"),
		).
		GroupBy(loanbinding.FieldStatus).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}
	if len(restotal) == 0 {
		result.CountNewBindingRevolvingIncome = 0
		result.SumNewBindingRevolvingIncome = 0
		return result, nil
	}

	result.CountNewBindingRevolvingIncome = restotal[0].Count
	result.SumNewBindingRevolvingIncome = float64(restotal[0].Count) * 100
	return result, nil
}

func (reportService DailyService) GetRtpTcrbLoanIncome(date string) (model.RtpTcrbLoanIncome, error) {
	var result model.RtpTcrbLoanIncome
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.RtpTcrbLoanIncome
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("RTP_TCRB_LOAN"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.TransactionType("TRANSFER"),
			//consumer.TransactionStatus("APPROVED"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}
	if len(restotal) == 0 {
		result.CountRtpTcrbLoanIncome = 0
		result.SumRtpTcrbLoanIncome = 0
		return result, nil
	}

	result.CountRtpTcrbLoanIncome = restotal[0].Count
	result.SumRtpTcrbLoanIncome = float64(restotal[0].Count) * 5
	return result, nil
}

func (reportService DailyService) GetRtpThaiPaiboonIncome(date string) (model.RtpThaiPaiboonIncome, error) {
	var result model.RtpThaiPaiboonIncome
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.RtpThaiPaiboonIncome
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.ToAccount("105011100200003"),
		).
		GroupBy(consumer.FieldToAccount).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}
	if len(restotal) == 0 {
		result.CountRtpThaiPaiboonIncome = 0
		result.SumRtpThaiPaiboonIncome = 0
		return result, nil
	}

	result.CountRtpThaiPaiboonIncome = restotal[0].Count
	result.SumRtpThaiPaiboonIncome = float64(restotal[0].Count) * 5
	return result, nil
}
