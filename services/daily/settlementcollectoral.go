package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/consumer"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetTCRBBillPaymentSettleCollectoralTwentyOneCutOff(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	startDate := stDate.Add(00 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TCRBBillPaymentSettleCollectoralTwentyOneCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("SETTLE_LOAN"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumTCRBBillPaymentSettle := restotal[0].Sum
	return sumTCRBBillPaymentSettle, nil
}

func (reportService DailyService) GetOnlineBillPaymentSettleCollectoralTwentyThreeCutOff(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	var restotal []model.OnlineBillPaymentSettleCollectoralTwentyThreeCutOff
	err := reportService.SumSettlementByPaymentType(date, "SETTLE_LOAN_REPAYMENT", &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}
	sumOnlineBillPaymentSettle := restotal[0].Sum
	return sumOnlineBillPaymentSettle, nil
}

func (reportService DailyService) GetPromptPayOutSettleCollectoralTwentyThreeCutOff(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	var restotal []model.PromptPayOutSettleCollectoralTwentyThreeCutOff
	err := reportService.SumSettlementByPaymentType(date, "SETTLE_PROMPTPAY", &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumPromptPayOutSettle := restotal[0].Sum
	return sumPromptPayOutSettle, nil
}

func (reportService DailyService) GetPromptPayOutTagThirtySettleCollectoralTwentyThreeCutOff(date string) (float64, error) {

	if date == "" {
		return 0, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}
	// startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	// endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	startDate := stDate.Add(00 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayOutTagThirySettleCollectoralTwentyThreeCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionStatus("APPROVE"),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("SETTLE_PROMPTPAY_TAG30"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}
	sumPromptpayOutTagThirtySettle := restotal[0].Sum
	return sumPromptpayOutTagThirtySettle, nil
}

func (reportService DailyService) SumSettlementByPaymentType(date string, status string, result interface{}) error {
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return err
	}
	// start := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	// end := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	start := stDate.Add(00 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	end := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(start),
			consumer.DateTimeLTE(end),
			consumer.PaymentType(status),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), result)
	if err != nil {
		return err
	}
	return nil
}

func (reportService DailyService) GetRtpTcrbLoanSettleCollectoralTwentyOneCutOff(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	// startDate := stDate.Add(21 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDate := stDate.Add(00 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	//startDateMinus := startDate.AddDate(0, 0, -1)

	// endDate := stDate.Add(20 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayOutTagThirySettleCollectoralTwentyThreeCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionStatus("APPROVE"),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("SETTLE_RTP_TCRB"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}
	sumRtpTcrbLoanTwentyOneSettle := restotal[0].Sum
	return sumRtpTcrbLoanTwentyOneSettle, nil
}

func (reportService DailyService) GetBillPayMeaSettleCollectoralTwentyThree(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	//startDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	//startDateMinus := startDate.AddDate(0,0,-1)
	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.BillPayMeaSettleCollectoralTwentyThree
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_BILLPAY_MEA"),
			consumer.PaymentChannel("CLOSELOOP"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumBillPayMeaSettleCollectoralTwentyThree := restotal[0].Sum
	return sumBillPayMeaSettleCollectoralTwentyThree, nil
}

func (reportService DailyService) GetBillPayPeaSettleCollectoralTwentyThree(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	//startDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	//startDateMinus := startDate.AddDate(0,0,-1)
	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.BillPayPeaSettleCollectoralTwentyThree
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_BILLPAY_PEA"),
			consumer.PaymentChannel("CLOSELOOP"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumBillPayMeaSettleCollectoralTwentyThree := restotal[0].Sum
	return sumBillPayMeaSettleCollectoralTwentyThree, nil
}

func (reportService DailyService) GetAisTopupSettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisTopupSettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_BILLPAY_AIS"),
			consumer.PaymentChannel("AIS"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumAisTopupSettleCollectoral := restotal[0].Sum
	return sumAisTopupSettleCollectoral, nil
}

func (reportService DailyService) GetAisPackageSettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisPackageSettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_PACKAGE_AIS"),
			consumer.PaymentChannel("AIS"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumAisPackageSettleCollectoral := restotal[0].Sum
	return sumAisPackageSettleCollectoral, nil
}

func (reportService DailyService) GetAisBillpaySettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisBillpaySettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_BILL_AIS"),
			consumer.PaymentChannel("AIS"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumAisBillpaySettleCollectoral := restotal[0].Sum
	return sumAisBillpaySettleCollectoral, nil
}

func (reportService DailyService) GetAisFiberSettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisFiberSettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_FIBER_AIS"),
			consumer.PaymentChannel("AIS"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumAisFiberSettleCollectoral := restotal[0].Sum
	return sumAisFiberSettleCollectoral, nil
}

func (reportService DailyService) GetDtacTopupSettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.DtacTopupSettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_BILLPAY_DTAC"),
			consumer.PaymentChannel("DTAC"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumDtacTopupSettleCollectoral := restotal[0].Sum
	return sumDtacTopupSettleCollectoral, nil
}

func (reportService DailyService) GetTrueTopupSettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TrueTopupSettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_BILLPAY_TRUEMOVE"),
			consumer.PaymentChannel("TRUEMOVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumTrueTopupSettleCollectoral := restotal[0].Sum
	return sumTrueTopupSettleCollectoral, nil
}

func (reportService DailyService) GetTruePackageSettleCollectoral(date string) (float64, error) {
	if date == "" {
		return 0, errors.New(" Date is required ")
	}
	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TruePackageSettleCollectoral
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentType("SETTLE_PAYMENT_PACKAGE_TRUEMOVE"),
			consumer.PaymentChannel("TRUEMOVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, err
	}

	sumTrueTopupSettleCollectoral := restotal[0].Sum
	return sumTrueTopupSettleCollectoral, nil
}
