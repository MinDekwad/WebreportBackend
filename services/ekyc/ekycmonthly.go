package ekyc

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
)

type CountAgentKYC struct {
	Count     int    `json:"count"`
	AgentID   string `json:"agentid"`
	KYCStatus string `json:"kycstatus"`
}

var KycStatus = []string{"APPROVED", "SUSPEND"}
var MerchantStatus = []string{"APPROVED", "SETTLE"}
var AgentID = []string{"Agent1", "KycTablet1", "KycTablet2", "Tassana.kyc"}

func (reportService EkycService) GetEKycMonthly(date string) (interface{}, error) {
	if date == "" {
		return nil, errors.New("Date is required")

	}

	if date == "undefined-01" {
		return nil, errors.New("Date format invalid")
	}

	start, end, err := reportService.GetStartAndEndOfMonth(date)
	if err != nil {
		return nil, err
	}

	restotal, err := reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.KYCDateGTE(start),
			agentkyc.KYCDateLT(end),
			agentkyc.KYCStatusIn(KycStatus...),
		).
		Order(ent.Desc(agentkyc.FieldID)).
		All(context.Background())

	if err != nil {
		return restotal, err
	}

	return restotal, nil
}

func (reportService EkycService) GetEKycMonthlyTmdsAll(date string) (int, error) {
	if date == "" {
		return 0, nil
	}
	start, end, err := reportService.GetStartAndEndOfMonth(date)
	if err != nil {
		return 0, err
	}

	var result []CountAgentKYC
	err = reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.KYCDateGT(start),
			agentkyc.KYCDateLT(end),
			agentkyc.KYCStatusIn(KycStatus...),
			agentkyc.AgentIDIn(AgentID...),
		).
		GroupBy(agentkyc.FieldKYCStatus).
		Aggregate(ent.Count()).
		Scan(context.Background(), &result)
	if err != nil {
		reportService.logger.Errorln("[Service] Get agentkyc monthly tmds error ", err)
		return 0, err
	}

	if result == nil || len(result) < 1 {
		reportService.logger.Debugln("[Service] Get agentkyc monthly tmds is null or empty ")
		return 0, nil
	}

	CountEkycTMDSALL := result[0].Count
	return CountEkycTMDSALL, nil
}

func (reportService EkycService) GetEKycMonthlyTcrbAll(date string) (int, error) {
	if date == "" {
		return 0, errors.New("Date is required")
	}
	start, end, err := reportService.GetStartAndEndOfMonth(date)
	if err != nil {
		return 0, err
	}

	var result []CountAgentKYC
	err = reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.KYCDateGTE(start),
			agentkyc.KYCDateLT(end),
			agentkyc.AgentIDNotIn(AgentID...),
			agentkyc.KYCStatusIn(KycStatus...),
			//agentkyc.AgentIDNotIn("Agent1", "KycTablet1", "KycTablet2", "Tassana.kyc"),
			// agentkyc.KYCStatusIn("APPROVED", "SUSPEND"),
		).
		GroupBy(agentkyc.FieldKYCStatus).
		Aggregate(ent.Count()).
		Scan(context.Background(), &result)
	if err != nil {
		return 0, nil
	}
	CountEkycTCRBALL := result[0].Count
	return CountEkycTCRBALL, nil
}
