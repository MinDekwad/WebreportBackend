package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/merchanttransaction"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetPromptpayInTagThirtySuspendMerchant(date string) (model.PromptpayInTagThirtySuspendMerchant, error) {
	var result model.PromptpayInTagThirtySuspendMerchant
	if date == "" {
		return result, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayInTagThirtySuspendMerchant
	err = reportService.connection.MerchantTransaction.Query().
		Where(
			merchanttransaction.DateTimeGTE(startDate),
			merchanttransaction.DateTimeLTE(endDate),
			merchanttransaction.PaymentChannel("PROMPTPAY"),
			merchanttransaction.PaymentType("PROMPTPAY_TAG30_ONLINE"),
			merchanttransaction.StatusIn(MerchantStatus...),
		).
		GroupBy(merchanttransaction.FieldPaymentType).
		Aggregate(ent.Sum(merchanttransaction.FieldAmount), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.SumPromptpayInTagThirtySuspendMerchant = 0
		result.CountPromptpayInTagThirtySuspendMerchant = 0
		return result, nil
	}

	result.SumPromptpayInTagThirtySuspendMerchant = restotal[0].Sum
	result.CountPromptpayInTagThirtySuspendMerchant = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptpayOutTagThirtySuspendMerchant(date string) (model.PromptpayOutTagThirtySuspendMerchant, error) {

	var result model.PromptpayOutTagThirtySuspendMerchant
	if date == "" {
		return result, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayOutTagThirtySuspendMerchant
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionStatus("APPROVE"),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("CASH_OUT_PROMPTPAY_TAG30_OFF_US"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumPromptpayOutTagThirtySuspendMerchant = 0
		result.CountPromptpayOutTagThirtySuspendMerchant = 0
		return result, nil
	}

	result.SumPromptpayOutTagThirtySuspendMerchant = restotal[0].Sum
	result.CountPromptpayOutTagThirtySuspendMerchant = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetRtpTcrbLoanSuspendMerchant(date string) (model.RtpTcrbLoanSuspendMerchant, error) {

	var result model.RtpTcrbLoanSuspendMerchant
	if date == "" {
		return result, errors.New("Date is required")
	}
	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(21 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.RtpTcrbLoanSuspendMerchant
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionStatus("APPROVE"),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("RTP_TCRB_LOAN"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumRtpTcrbLoanSuspendMerchant = 0
		result.CountRtpTcrbLoanSuspendMerchant = 0
		return result, nil
	}

	result.SumRtpTcrbLoanSuspendMerchant = restotal[0].Sum
	result.CountRtpTcrbLoanSuspendMerchant = restotal[0].Count
	return result, nil
}
