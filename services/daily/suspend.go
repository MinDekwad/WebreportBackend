package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/consumer"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetPromptpayInOtherBankTwentyThreeCutOff(date string) (model.PromptpayInOtherBankTwentyThreeCutOff, error) {
	var result model.PromptpayInOtherBankTwentyThreeCutOff
	if date == "" {
		return result, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayInOtherBankTwentyThreeCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY_IN"),
			consumer.BankCodeNEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.SumPromptpayInOtherBankSuspend = 0
		result.CountPromptpayInOtherBankSuspend = 0
		return result, nil
	}

	result.SumPromptpayInOtherBankSuspend = restotal[0].Sum
	result.CountPromptpayInOtherBankSuspend = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptpayOutOtherBankTwentyThreeCutOff(date string) (model.PromptpayOutOtherBankTwentyThreeCutOff, error) {
	var result model.PromptpayOutOtherBankTwentyThreeCutOff
	if date == "" {
		return result, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayOutOtherBankTwentyThreeCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY29_CASH_OUT"),
			consumer.BankCodeNEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)

	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.CountPromptpayOutOtherBankSuspend = 0
		result.SumPromptpayOutOtherBankSuspend = 0
		return result, nil
	}

	result.CountPromptpayOutOtherBankSuspend = restotal[0].Count
	result.SumPromptpayOutOtherBankSuspend = restotal[0].Sum
	return result, nil
}

func (reportService DailyService) GetPromptpayOutTCRBTwentyThreeCutOff(date string) (model.PromptpayOutTCRBTwentyThreeCutOff, error) {
	var result model.PromptpayOutTCRBTwentyThreeCutOff
	if date == "" {
		return result, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayOutTCRBTwentyThreeCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY29_CASH_OUT"),
			consumer.BankCodeEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)

	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.CountPromptpayOutTCRBSuspend = 0
		result.SumPromptpayOutTCRBSuspend = 0
		return result, nil
	}

	result.CountPromptpayOutTCRBSuspend = restotal[0].Count
	result.SumPromptpayOutTCRBSuspend = restotal[0].Sum
	return result, nil
}

func (reportService DailyService) GetTCRBBillPaymentTwentyOneCutOff(date string) (model.TCRBBillPaymentTwentyOneCutOff, error) {
	var result model.TCRBBillPaymentTwentyOneCutOff
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(21 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	// startDate := stDate.Add(21 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	// startDateMinus := startDate.AddDate(0, 0, -1)
	// endDate := stDate.Add(20 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TCRBBillPaymentTwentyOneCutOff
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("TCRB_LOAN"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.CountTCRBBillPaymentSuspend = 0
		result.SumTCRBBillPaymentSuspend = 0
		return result, nil
	}

	result.CountTCRBBillPaymentSuspend = restotal[0].Count
	result.SumTCRBBillPaymentSuspend = restotal[0].Sum
	return result, nil
}
