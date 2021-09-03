package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/consumer"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetAdjustFromWallet(date string) (model.AdjustFromWallet, error) {
	var result model.AdjustFromWallet
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AdjustFromWallet
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("ADJUST_WALLET_Out"),
		).
		GroupBy(consumer.FieldPaymentChannel).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumAdjustFromWalletTranOut = 0
		result.CountAdjustFromWalletTranOut = 0
		return result, nil
	}

	result.SumAdjustFromWalletTranOut = restotal[0].Sum
	result.CountAdjustFromWalletTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptPayOutOtherBank(date string) (model.PromptPayOutOtherBank, error) {
	var result model.PromptPayOutOtherBank
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayOutOtherBank
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

	if len(restotal) == 0 {
		result.SumPromptPayOutOtherBankTranOut = 0
		result.CountPromptPayOutOtherBankTranOut = 0
		return result, nil
	}

	result.SumPromptPayOutOtherBankTranOut = restotal[0].Sum
	result.CountPromptPayOutOtherBankTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptPayOutTCRB(date string) (model.PromptPayOutTCRB, error) {
	var result model.PromptPayOutTCRB
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayOutTCRB
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

	if len(restotal) == 0 {
		result.SumPromptPayOutTCRBTranOut = 0
		result.CountPromptPayOutTCRBTranOut = 0
		return result, nil
	}

	result.SumPromptPayOutTCRBTranOut = restotal[0].Sum
	result.CountPromptPayOutTCRBTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTcrbBillPayment(date string) (model.TcrbBillPayment, error) {
	var result model.TcrbBillPayment
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TcrbBillPayment
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

	if len(restotal) == 0 {
		result.SumTcrbBillPaymentTranOut = 0
		result.CountTcrbBillPaymentTranOut = 0
		return result, nil
	}

	result.SumTcrbBillPaymentTranOut = restotal[0].Sum
	result.CountTcrbBillPaymentTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTransferToBankAccountTxn(date string) (model.TransferToBankAccountTxn, error) {
	var result model.TransferToBankAccountTxn
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TransferToBankAccountTxn
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("TCRB_BULK_PAYMENT"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldAmount), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumTransferToBankAccountTxnTranOut = 0
		result.CountTransferToBankAccountTxnTranOut = 0
		return result, nil
	}

	result.SumTransferToBankAccountTxnTranOut = restotal[0].Sum
	result.CountTransferToBankAccountTxnTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTransferToBankAccountFee(date string) (model.TransferToBankAccountFee, error) {
	var result model.TransferToBankAccountFee
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TransferToBankAccountFee
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.PaymentType("TCRB_BULK_PAYMENT"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldFee), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumTransferToBankAccountFeeTranOut = 0
		result.CountTransferToBankAccountFeeTranOut = 0
		return result, nil
	}

	result.SumTransferToBankAccountFeeTranOut = restotal[0].Sum
	result.CountTransferToBankAccountFeeTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetOnlineBillPayment(date string) (model.OnlineBillPayment, error) {
	var result model.OnlineBillPayment
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineBillPayment
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
		result.SumOnlineBillPaymentTranOut = 0
		result.CountOnlineBillPaymentTranOut = 0
		return result, nil
	}

	result.SumOnlineBillPaymentTranOut = restotal[0].Sum
	result.CountOnlineBillPaymentTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptpayOutTagThirty(date string) (model.PromptpayOutTagThirty, error) {
	var result model.PromptpayOutTagThirty
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayOutTagThirty
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
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
		result.SumPromptpayOutTagThirtyTranOut = 0
		result.CountPromptpayOutTagThirtyTranOut = 0
		return result, nil
	}

	result.SumPromptpayOutTagThirtyTranOut = restotal[0].Sum
	result.CountPromptpayOutTagThirtyTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetRtpTCRBLoanTranOut(date string) (model.TCRBLoanTranOut, error) {
	var result model.TCRBLoanTranOut
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TCRBLoanTranOut
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
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
		result.SumTCRBLoanTranOut = 0
		result.CountTCRBLoanTranOut = 0
		return result, nil
	}

	result.SumTCRBLoanTranOut = restotal[0].Sum
	result.CountTCRBLoanTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetRtpThaiHealthTranOut(date string) (model.ThaiHealthTranOut, error) {
	var result model.ThaiHealthTranOut
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.ThaiHealthTranOut
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.ToAccount("105011100200002"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumThaiHealthTranOut = 0
		result.CountThaiHealthTranOut = 0
		return result, nil
	}

	result.SumThaiHealthTranOut = restotal[0].Sum
	result.CountThaiHealthTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetRtpThaiPaiboonTranOut(date string) (model.ThaiPaiboonTranOut, error) {
	var result model.ThaiPaiboonTranOut
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.ThaiPaiboonTranOut
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.ToAccount("105011100200003"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumThaiPaiboonTranOut = 0
		result.CountThaiPaiboonTranOut = 0
		return result, nil
	}

	result.SumThaiPaiboonTranOut = restotal[0].Sum
	result.CountThaiPaiboonTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetBillPayMeaTranOut(date string) (model.BillPayMeaTranOut, error) {
	var result model.BillPayMeaTranOut

	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.BillPayMeaTranOut
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentType("BILLPAYMENT_MEA"),
			consumer.PaymentChannel("MEA"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumBillPayMeaTranOut = 0
		result.CountBillPayMeaTranOut = 0
		return result, nil
	}

	result.SumBillPayMeaTranOut = restotal[0].Sum
	result.CountBillPayMeaTranOut = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetBillPayPeaTranOut(date string) (model.BillPayPeaTranOut, error) {
	var result model.BillPayPeaTranOut

	if date == "" {
		return result, errors.New(" Date is required ")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.BillPayPeaTranOut
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentType("BILLPAYMENT_PEA"),
			consumer.PaymentChannel("PEA"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumBillPayPeaTranOut = 0
		result.CountBillPayPeaTranOut = 0
		return result, nil
	}

	result.SumBillPayPeaTranOut = restotal[0].Sum
	result.CountBillPayPeaTranOut = restotal[0].Count
	return result, nil
}
