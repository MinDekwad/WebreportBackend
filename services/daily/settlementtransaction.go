package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/bulk"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/merchanttransaction"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetTcrbBillPaymentSettlementTransaction(date string) (model.TcrbBillPaymentSettlementTransaction, error) {
	var result model.TcrbBillPaymentSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(21 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)
	endDate := stDate.Add(20 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TcrbBillPaymentSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
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

	if len(restotal) == 0 {
		result.SumTcrbBillPaymentSettlementTransaction = 0
		result.CountTcrbBillPaymentSettlementTransaction = 0
		return result, nil
	}

	result.SumTcrbBillPaymentSettlementTransaction = restotal[0].Sum
	result.CountTcrbBillPaymentSettlementTransaction = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetOnlineBillPaymentSettlementTransaction(date string) (model.OnlineBillPaymentSettlementTransaction, error) {
	var result model.OnlineBillPaymentSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineBillPaymentSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
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
		result.SumOnlineBillPaymentSettlementTransaction = 0
		result.CountOnlineBillPaymentSettlementTransaction = 0
		return result, nil
	}

	result.SumOnlineBillPaymentSettlementTransaction = restotal[0].Sum
	result.CountOnlineBillPaymentSettlementTransaction = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptPayOutSettlementTransaction(date string) (model.PromptPayOutSettlementTransaction, error) {
	var result model.PromptPayOutSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayOutSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
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

	if len(restotal) == 0 {
		result.SumPromptPayOutSettlementTransaction = 0
		result.CountPromptPayOutSettlementTransaction = 0
		return result, nil
	}
	result.SumPromptPayOutSettlementTransaction = restotal[0].Sum
	result.CountPromptPayOutSettlementTransaction = restotal[0].Count
	return result, nil
}

//func (reportService DailyService) GetBulkCreditSameday(date string) (interface{}, error) {
func (reportService DailyService) GetBulkCreditSameday(date string) (model.BulkCreditSameday, error) {
	var result model.BulkCreditSameday

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

	restotal, err := reportService.connection.Bulk.Query().Select(bulk.FieldBulkCreditSameday, bulk.FieldTransfertobankaccount).
		Where(
			bulk.DateTimeGTE(startDateMinus),
			bulk.DateTimeLTE(endDate),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.BulkCreditSameday = 0
		result.Transfertobankaccount = 0
		return result, nil
	}

	result.BulkCreditSameday = restotal[0].BulkCreditSameday
	result.Transfertobankaccount = restotal[0].Transfertobankaccount
	return result, nil
}

func (reportService DailyService) GetBulkCreditSamedayFee(date string) (model.BulkCreditSamedayFee, error) {
	var result model.BulkCreditSamedayFee
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
		result.BulkCreditSamedayFee = 0
		result.TransfertobankaccountFee = 0
		return result, nil
	}

	result.BulkCreditSamedayFee = restotal[0].BulkCreditSamedayFee
	result.TransfertobankaccountFee = restotal[0].Transfertobankaccount
	return result, nil
}

func (reportService DailyService) GetTopupLoanDisbursementSettlementTransaction(date string) (model.GetTopupLoanDisbursementSettlementTransaction, error) {
	var result model.GetTopupLoanDisbursementSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	//startDateMinus := stDate.AddDate(0, 0, -1)
	if err != nil {
		return result, err
	}

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.GetTopupLoanDisbursementSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(stDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("DIRECT_TCRB"),
			consumer.PaymentType("LOAN_DISBURSE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.CountTopupLoanDisbursementSettlementTransaction = 0
		result.SumTopupLoanDisbursementSettlementTransaction = 0
		return result, nil
	}

	result.CountTopupLoanDisbursementSettlementTransaction = restotal[0].Count
	result.SumTopupLoanDisbursementSettlementTransaction = restotal[0].Sum
	return result, nil
}

func (reportService DailyService) GetOnlineLoanTopupSettlementTransaction(date string) (model.OnlineLoanTopupSettlementTransaction, error) {
	var result model.OnlineLoanTopupSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	startDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineLoanTopupSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("LOAN_DIRECT_TCRB"),
			consumer.PaymentType("TOPUP_ONLINE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.CountOnlineLoanTopupSettlementTransaction = 0
		result.SumOnlineLoanTopupSettlementTransaction = 0
		return result, nil
	}

	result.CountOnlineLoanTopupSettlementTransaction = restotal[0].Count
	result.SumOnlineLoanTopupSettlementTransaction = restotal[0].Sum
	return result, nil
}

func (reportService DailyService) GetPromptPayInOtherBankSettlementTransaction(date string) (model.PromptPayInOtherBankSettlementTransaction, error) {
	var result model.PromptPayInOtherBankSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

	var restotal []model.PromptPayInOtherBankSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
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

	if len(restotal) == 0 {
		result.CountPromptPayInOtherBankSettlementTransaction = 0
		result.SumPromptPayInOtherBankSettlementTransaction = 0
		return result, nil
	}

	result.CountPromptPayInOtherBankSettlementTransaction = restotal[0].Count
	result.SumPromptPayInOtherBankSettlementTransaction = restotal[0].Sum
	return result, nil
}

func (reportService DailyService) GetPromptPayInTagThirtySettlementTransaction(date string) (model.PromptPayInTagThirtySettlementTransaction, error) {
	var result model.PromptPayInTagThirtySettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayInTagThirtySettlementTransaction
	err = reportService.connection.MerchantTransaction.Query().
		Where(
			merchanttransaction.DateTimeGTE(startDateMinus),
			merchanttransaction.DateTimeLTE(endDate),
			merchanttransaction.PaymentType("PROMPTPAY_TAG30_ONLINE"),
			merchanttransaction.PaymentChannel("PROMPTPAY"),
			merchanttransaction.StatusIn("SETTLE", "APPROVED"),
		).
		GroupBy(merchanttransaction.FieldPaymentType).
		Aggregate(ent.Sum(merchanttransaction.FieldAmount), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumPromptPayInTagThirtySettlementTransaction = 0
		result.CountPromptPayInTagThirtySettlementTransaction = 0
		return result, nil
	}

	result.SumPromptPayInTagThirtySettlementTransaction = restotal[0].Sum
	result.CountPromptPayInTagThirtySettlementTransaction = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptPayOutTagThirtySettlementTransaction(date string) (model.PromptPayOutTagThirtySettlementTransaction, error) {
	var result model.PromptPayOutTagThirtySettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayOutTagThirtySettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("CASH_OUT_PROMPTPAY_TAG30_OFF_US"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldAmount), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumPromptPayOutTagThirtySettlementTransaction = 0
		result.CountPromptPayOutTagThirtySettlementTransaction = 0
		return result, nil
	}

	result.SumPromptPayOutTagThirtySettlementTransaction = restotal[0].Sum
	result.CountPromptPayOutTagThirtySettlementTransaction = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetRtpTcrbLoanTwentyOneSettlementTransaction(date string) (model.RtpTcrbLoanTwentyOneSettlementTransaction, error) {
	var result model.RtpTcrbLoanTwentyOneSettlementTransaction
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	// startDate := stDate.Add(23 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	// startDateMinus := startDate.AddDate(0, 0, -1)
	// endDate := stDate.Add(22 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	startDate := stDate.Add(21 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)
	endDate := stDate.Add(20 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.RtpTcrbLoanTwentyOneSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDateMinus),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("RTP_TCRB_LOAN"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.TransactionStatus("APPROVE"),
			consumer.TransactionType("TRANSFER"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldAmount), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumRtpTcrbLoanTwentyOneSettlementTransaction = 0
		result.CountRtpTcrbLoanTwentyOneSettlementTransaction = 0
		return result, nil
	}

	result.SumRtpTcrbLoanTwentyOneSettlementTransaction = restotal[0].Sum
	result.CountRtpTcrbLoanTwentyOneSettlementTransaction = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetBillPayMeaTwentyThreeSettlementTransaction(date string) (model.BillPayMeaTwentyThreeSettlementTransaction, error) {
	var result model.BillPayMeaTwentyThreeSettlementTransaction

	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.BillPayMeaTwentyThreeSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("BILLPAYMENT_MEA"),
			consumer.PaymentChannel("MEA"),
			consumer.TransactionType("PAYMENT"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumBillPayMeaTwentyThreeSettlementTransaction = 0
		result.CountBillPayMeaTwentyThreeSettlementTransaction = 0
		return result, nil
	}

	result.SumBillPayMeaTwentyThreeSettlementTransaction = restotal[0].Sum
	result.CountBillPayMeaTwentyThreeSettlementTransaction = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetBillPayPeaTwentyThreeSettlementTransaction(date string) (model.BillPayPeaTwentyThreeSettlementTransaction, error) {
	var result model.BillPayPeaTwentyThreeSettlementTransaction

	if date == "" {
		return result, errors.New(" Date is required ")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.BillPayPeaTwentyThreeSettlementTransaction
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("BILLPAYMENT_PEA"),
			consumer.PaymentChannel("PEA"),
			consumer.TransactionType("PAYMENT"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumBillPayPeaTwentyThreeSettlementTransaction = 0
		result.CountBillPayPeaTwentyThreeSettlementTransaction = 0
		return result, nil
	}

	result.SumBillPayPeaTwentyThreeSettlementTransaction = restotal[0].Sum
	result.CountBillPayPeaTwentyThreeSettlementTransaction = restotal[0].Count
	return result, nil
}
