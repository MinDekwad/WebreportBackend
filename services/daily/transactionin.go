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

func (reportService DailyService) GetWalletToWallet(date string) (model.WalletToWallet, error) {
	var result model.WalletToWallet
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.WalletToWallet
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("WALLET"),
		).
		GroupBy(consumer.FieldPaymentChannel).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumWalletToWallet = 0
		result.CountWalletToWallet = 0
		return result, nil
	}

	result.SumWalletToWallet = restotal[0].Sum
	result.CountWalletToWallet = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetSettleMerchantOnline(date string) (model.SettleMerchantOnline, error) {
	var result model.SettleMerchantOnline
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.SettleMerchantOnline
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("MERCHANT"),
			consumer.PaymentType("SETTLE_MERCHANT"),
		).
		GroupBy(consumer.FieldPaymentChannel).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.SumSettleMerchantOnline = 0
		result.CountSettleMerchantOnline = 0
		return result, nil
	}

	result.SumSettleMerchantOnline = restotal[0].Sum
	result.CountSettleMerchantOnline = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetAdjustToWallet(date string) (model.AdjustToWallet, error) {
	var result model.AdjustToWallet
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AdjustToWallet
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("CLOSELOOP"),
			consumer.PaymentType("ADJUST_WALLET_IN"),
		).
		GroupBy(consumer.FieldPaymentChannel).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumAdjustToWallet = 0
		result.CountAdjustToWallet = 0
		return result, nil
	}

	result.SumAdjustToWallet = restotal[0].Sum
	result.CountAdjustToWallet = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptPayInOtherBank(date string) (model.PromptPayInOtherBank, error) {
	var result model.PromptPayInOtherBank
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayInOtherBank
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

	if len(restotal) == 0 {
		result.SumPromptPayInOtherBankTranIn = 0
		result.CountPromptPayInOtherBankTranIn = 0
		return result, nil
	}

	result.SumPromptPayInOtherBankTranIn = restotal[0].Sum
	result.CountPromptPayInOtherBankTranIn = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptPayInTCRB(date string) (model.PromptPayInTCRB, error) {
	var result model.PromptPayInTCRB
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptPayInTCRB
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY_IN"),
			consumer.BankCodeEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 { // code to check array. use len()
		result.SumPromptPayInTCRBTranIn = 0
		result.CountPromptPayInTCRBTranIn = 0
		return result, nil
	}

	result.SumPromptPayInTCRBTranIn = restotal[0].Sum
	result.CountPromptPayInTCRBTranIn = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetPromptpayInTagThirty(date string) (model.PromptpayInTagThirty, error) {
	var result model.PromptpayInTagThirty
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.PromptpayInTagThirty
	err = reportService.connection.MerchantTransaction.Query().
		Where(
			merchanttransaction.DateTimeGTE(startDate),
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

	if len(restotal) == 0 { // code to check array. use len()
		result.SumPromptpayInTagThirtyTranIn = 0
		result.CountPromptpayInTagThirtyTranIn = 0
		return result, nil
	}

	result.SumPromptpayInTagThirtyTranIn = restotal[0].Sum
	result.CountPromptpayInTagThirtyTranIn = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTopupLoanDisbursement(date string) (model.TopupLoanDisbursement, error) {
	var result model.TopupLoanDisbursement
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TopupLoanDisbursement
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
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
		result.SumTopupLoanDisbursementTranIn = 0
		result.CountTopupLoanDisbursementTranIn = 0
		return result, nil
	}

	result.SumTopupLoanDisbursementTranIn = restotal[0].Sum
	result.CountTopupLoanDisbursementTranIn = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTopupDirectDebit(date string) (model.TopupDirectDebit, error) {
	var result model.TopupDirectDebit
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TopupDirectDebit
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("DIRECT_DEBIT"),
			consumer.PaymentType("TCRB_DEPOSIT"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.CountTopupDirectDebitTranIn = 0
		result.SumTopupDirectDebitTranIn = 0
		return result, nil
	}

	result.CountTopupDirectDebitTranIn = restotal[0].Count
	result.SumTopupDirectDebitTranIn = restotal[0].Sum
	return result, nil
}

func (reportService DailyService) GetTopupPayRoll(date string) (model.TopupPayRoll, error) {
	var result model.TopupPayRoll
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TopupPayRoll
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("PAYROLL"),
			consumer.PaymentType("TCRB_PAY_ROLL"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumTopupPayRollTranIn = 0
		result.CountTopupPayRollTranIn = 0
		return result, nil
	}

	result.SumTopupPayRollTranIn = restotal[0].Sum
	result.CountTopupPayRollTranIn = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetOnlineLoanTopup(date string) (model.OnlineLoanTopup, error) {
	var result model.OnlineLoanTopup
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineLoanTopup
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
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
		result.SumOnlineLoanTopupTranIn = 0
		result.CountOnlineLoanTopupTranIn = 0
		return result, nil
	}

	result.SumOnlineLoanTopupTranIn = restotal[0].Sum
	result.CountOnlineLoanTopupTranIn = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetCashBack(date string) (model.CashBack, error) {
	var result model.CashBack
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.OnlineLoanTopup
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("PROMOTION"),
			consumer.PaymentType("CASH_BACK"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumCashBackTranIn = 0
		result.CountCashBackTranIn = 0
		return result, nil
	}

	result.SumCashBackTranIn = restotal[0].Sum
	result.CountCashBackTranIn = restotal[0].Count
	return result, nil
}
