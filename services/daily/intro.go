package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/bankdetail"
	"go-api-report2/ent/reportwallet"
	"go-api-report2/ent/statementendingbalanc"
	"go-api-report2/model"
	"time"
)

type CountAgentKYC struct {
	Count     int    `json:"count"`
	AgentID   string `json:"agentid"`
	KYCStatus string `json:"kycstatus"`
}

const (
	timeCustomLayout = "2006-01-02"
)

var KycStatus = []string{"APPROVED", "SUSPEND"}
var MerchantStatus = []string{"APPROVED", "SETTLE"}
var AgentID = []string{"Agent1", "KycTablet1", "KycTablet2", "Tassana.kyc"}

func (reportService DailyService) GetNewWallet(date string) (int, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}
	_, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}
	var newWallet []model.NewWallet
	err = reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.KYCDate(date),
			agentkyc.KYCStatus("APPROVED"),
		).
		GroupBy(agentkyc.FieldKYCStatus).
		Aggregate(ent.Count()).
		Scan(context.Background(), &newWallet)
	if err != nil {
		return 0, err
	}

	if newWallet == nil || len(newWallet) < 1 {
		return 0, err
	}

	return newWallet[0].Count, nil
}

func (reportService DailyService) GetNoOfWallet(date string) (int, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}
	startDatePlus := startDate.AddDate(0, 0, 1)

	endDate := startDatePlus.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.NoOfWallet
	err = reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTimeGTE(startDatePlus),
			reportwallet.DateTimeLTE(endDate),
			reportwallet.StatusIn(KycStatus...),
		).
		GroupBy(reportwallet.FieldWalletTypeName).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if restotal == nil || len(restotal) < 1 {
		return 0, err
	}

	CountNoOfWallet := restotal[0].Count
	return CountNoOfWallet, nil
}

func (reportService DailyService) GetWalletBalance(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	startDatePlus := startDate.AddDate(0, 0, 1)
	endDate := startDatePlus.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	restotal := make([]model.WalletBalance, 0)
	err = reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTimeGTE(startDatePlus),
			reportwallet.DateTimeLTE(endDate),
			reportwallet.StatusIn(KycStatus...),
		).
		GroupBy(reportwallet.FieldWalletTypeName).
		Aggregate(ent.Sum(reportwallet.FieldBalance)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if len(restotal) == 0 {
		return 0, nil
	}

	SumWalletBalance := restotal[0].Sum
	return SumWalletBalance, nil
}

func (reportService DailyService) GetDiffWalletBalance(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}

	startDatePlus := startDate.AddDate(0, 0, 1)
	endDate := startDatePlus.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	restotal1 := make([]model.WalletBalance, 0)
	err = reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTimeGTE(startDatePlus),
			reportwallet.DateTimeLTE(endDate),
			reportwallet.StatusIn(KycStatus...),
		).
		GroupBy(reportwallet.FieldWalletTypeName).
		Aggregate(ent.Sum(reportwallet.FieldBalance)).
		Scan(context.Background(), &restotal1)
	if err != nil {
		return 0, err
	}

	if len(restotal1) == 0 {
		return 0, nil
	}

	endDate2 := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	restotal2 := make([]model.WalletBalance, 0)
	err = reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTimeGTE(startDate),
			reportwallet.DateTimeLTE(endDate2),
			reportwallet.StatusIn(KycStatus...),
		).
		GroupBy(reportwallet.FieldWalletTypeName).
		Aggregate(ent.Sum(reportwallet.FieldBalance)).
		Scan(context.Background(), &restotal2)
	if err != nil {
		return 0, err
	}

	if len(restotal2) == 0 {
		return 0, nil
	}

	SumDiffWalletBalance := restotal1[0].Sum - restotal2[0].Sum
	return SumDiffWalletBalance, nil
}

func (reportService DailyService) SumWalletBalance(startDate time.Time) (model.DiffWalletBalance, error) {
	var results []model.DiffWalletBalance
	var endDate time.Time

	y, m, d := startDate.Date()
	endDate = time.Date(y, m, d, 23, 59, 59, 0, startDate.Location())

	err := reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTimeGTE(startDate),
			reportwallet.DateTimeLTE(endDate),
			reportwallet.WalletTypeName("CUSTOMER"),
			reportwallet.StatusIn(KycStatus...),
		).
		GroupBy(reportwallet.FieldWalletTypeName).
		Aggregate(ent.Sum(reportwallet.FieldBalance)).
		Scan(context.Background(), &results)
	if err != nil {
		return model.DiffWalletBalance{}, err
	}
	if results == nil || len(results) == 0 {
		return model.DiffWalletBalance{}, nil
	}
	count, err := reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTimeGTE(startDate),
			reportwallet.DateTimeLTE(endDate),
			reportwallet.WalletTypeName("CUSTOMER"),
			reportwallet.StatusIn(KycStatus...),
		).Count(context.Background())
	results[0].Total = count
	return results[0], nil
}

func (reportService DailyService) GetTotalBalanceInMicroPay(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}
	startDatePlus := startDate.AddDate(0, 0, 1)

	var restotal []model.DiffWalletBalance
	err = reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.DateTime(startDatePlus),
		).
		GroupBy(reportwallet.FieldWalletTypeName).
		Aggregate(ent.Sum(reportwallet.FieldBalance)).
		Scan(context.Background(), &restotal)
	if restotal == nil || len(restotal) < 1 {
		return 0, nil
	}

	var total float64

	for _, s1 := range restotal {
		total = total + s1.Sum
	}

	return total, nil
}

func (reportService DailyService) GetStatementEndingBalance(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}
	y, m, d := stDate.Date()

	startDate := time.Date(y, m, d, 0, 0, 0, 0, stDate.Location())
	reportService.logger.Infof("stateDate %v", startDate)

	endDate := time.Date(y, m, d, 23, 59, 59, 0, stDate.Location())
	reportService.logger.Infof("endDate %v", endDate)

	var restotal []model.StatementEndingBalanc
	err = reportService.connection.StatementEndingBalanc.Query().
		Where(
			statementendingbalanc.StatementDateGTE(startDate),
			statementendingbalanc.StatementDateLTE(endDate),
			statementendingbalanc.HasBankWith(bankdetail.ID(2)),
		).
		GroupBy(statementendingbalanc.FieldID).
		Aggregate(ent.Sum(statementendingbalanc.FieldStatementBalance)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if restotal == nil || len(restotal) == 0 {
		return 0, err
	}

	SumStatementEndingBalance := restotal[0].Sum
	return SumStatementEndingBalance, nil
}

func (reportService DailyService) GetStatementIncomingBalance(date string) (float64, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return 0, err
	}
	startDate := stDate.Add(00 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.StatementIncomingBalance
	err = reportService.connection.StatementEndingBalanc.Query().
		Where(
			statementendingbalanc.StatementDateGTE(startDate),
			statementendingbalanc.StatementDateLTE(endDate),
			statementendingbalanc.HasBankWith(bankdetail.ID(4)),
			// statementendingbalanc.BankUID(4),
		).
		GroupBy(statementendingbalanc.FieldID).
		Aggregate(ent.Sum(statementendingbalanc.FieldStatementBalance)).
		Scan(context.Background(), &restotal)
	if err != nil {
		return 0, err
	}

	if restotal == nil || len(restotal) == 0 {
		return 0, nil
	}

	SumStatementIncomingBalance := restotal[0].Sum
	return SumStatementIncomingBalance, nil
}
