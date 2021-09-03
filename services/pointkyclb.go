package services

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/loanbinding"
	"go-api-report2/ent/pendingkyc"
	"go-api-report2/ent/pendingloanbinding"
	"go-api-report2/ent/pointpendingkyctransaction"
	"go-api-report2/ent/pointpendinglbtransaction"
	"go-api-report2/ent/predicate"
	"time"
)

const (
	timeCustomLayout = "2006-01-02"
	layoutISO        = "2006-01-02 00:00:00"
	layoutISOA       = "2006-01-02 15:04:05"
	dateTimeCustom   = "2006-01-02"
)

func (r ReportService) CalculatePointKYC(date string) (string, error) {
	var result string
	result = "Error"
	if date == "" {
		return "", errors.New("Date is required")
	}

	count, err := r.connection.Pendingkyc.Query().
		Where(
			pendingkyc.KYCDate(date),
		).Count(context.Background())
	if err != nil {
		return result, err
	}

	// ถ้ามีข้อมูล
	if count > 0 {
		result = "Duplicate"
		return result, err
	}

	agkyc, err := r.connection.Agentkyc.Query().Where(agentkyc.KYCDate(date), agentkyc.KYCStatus("APPROVED")).All(context.Background())
	if err != nil {
		return "", err
	}

	for _, agentKYC := range agkyc {
		pkyc, err := r.connection.Pendingkyc.Query().
			Where(pendingkyc.WalletID(*agentKYC.Consumerwalletid)).All(context.Background())
		if err != nil {
			return "", err
		}
		if len(pkyc) != 0 {

			_, err := r.connection.Pendingkyc.
				Update().
				Where(
					pendingkyc.WalletID(*agentKYC.Consumerwalletid),
				).
				SetAgentID(*agentKYC.AgentID).
				SetAgentNameLastname(*agentKYC.AgentNameLastname).
				SetKYCDate(*agentKYC.KYCDate).
				SetDateGen(time.Now()).
				SetStatusGen(true).
				SetPoint(50).
				Save(context.Background())
			if err != nil {
				return "", err
			}
			result = "Success"
		}
	}
	return result, err
}

func (r ReportService) CalculatePointLB(date string) (string, error) {
	var result string
	result = "Error"
	if date == "" {
		return "", errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	// นับ count point pending kyc table
	count, err := r.connection.Pendingloanbinding.Query().
		Where(
			pendingloanbinding.DateTimeGTE(startDate),
			pendingloanbinding.DateTimeLTE(endDate),
		).Count(context.Background())
	if err != nil {
		return result, err
	}

	// ถ้ามีข้อมูล
	if count > 0 {
		result = "Duplicate"
		return result, err
	}

	loanbind, err := r.connection.Loanbinding.Query().
		Where(
			loanbinding.DateTimeGTE(startDate),
			loanbinding.DateTimeLTE(endDate),
			loanbinding.Status("APPROVED"),
		).
		All(context.Background())
	if err != nil {
		return "", err
	}

	for _, lb := range loanbind {
		plb, err := r.connection.Pendingloanbinding.Query().
			Where(pendingloanbinding.WalletID(*lb.WalletId)).All(context.Background())
		if err != nil {
			return "", err
		}
		if len(plb) != 0 {
			// 		//fmt.Println(pkyc)
			_, err := r.connection.Pendingloanbinding.
				Update().
				Where(
					pendingloanbinding.WalletID(*lb.WalletId),
				).
				SetDateGenLB(time.Now()).
				SetDateTime(*lb.DateTime).
				SetStatusGenLB(true).
				SetPointLB(50).
				Save(context.Background())
			if err != nil {
				return "", err
			}
			result = "Success"
		}
	}
	return result, err
}

func (r ReportService) CreatePointTransacKYCCSV(date string, note string) (interface{}, error) {
	if date == "" {
		return nil, errors.New("Date is required")
	}

	_, err := r.connection.Pointpendingkyctransaction.
		Delete().
		Where(pointpendingkyctransaction.KYCDate(date)).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}

	pendingkyc, err := r.connection.Pendingkyc.Query().
		Where(
			pendingkyc.KYCDate(date),
			// pendingkyc.DateGenGTE(stDate),
			// pendingkyc.DateGenLTE(endDate),
		).All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, pkyc := range pendingkyc {
		if pkyc != nil {

			_, err := r.connection.Pointpendingkyctransaction.Create().
				SetWalletId(pkyc.WalletID).
				SetDateExport(time.Now()).
				SetDateGen(*pkyc.DateGen).
				SetNote(note).
				SetPoint(pkyc.Point).
				SetStatusExport(true).
				SetKYCDate(*pkyc.KYCDate).
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (r ReportService) ExportPointKycCSV(date string) (interface{}, error) {
	if date == "" {
		return nil, errors.New("Date is required")
	}

	pkyc, err := r.connection.Pointpendingkyctransaction.Query().
		Select(pointpendingkyctransaction.FieldWalletId, pointpendingkyctransaction.FieldPoint, pointpendingkyctransaction.FieldDateExport, pointpendingkyctransaction.FieldNote).
		Where(
			pointpendingkyctransaction.KYCDate(date),
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return pkyc, nil
}

func (r ReportService) CreatePointTransacLBCSV(date string, note string) (interface{}, error) {
	if date == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, date)
	if err != nil {
		return nil, err
	}
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	_, err = r.connection.Pointpendinglbtransaction.
		Delete().
		Where(pointpendinglbtransaction.LBDateGTE(stDate), pointpendinglbtransaction.LBDateLTE(endDate)).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}

	pendinglb, err := r.connection.Pendingloanbinding.Query().
		Where(
			pendingloanbinding.DateTimeGTE(stDate),
			pendingloanbinding.DateTimeLTE(endDate),
		).All(context.Background())
	if err != nil {
		return nil, err
	}
	for _, plb := range pendinglb {
		if plb != nil {

			_, err := r.connection.Pointpendinglbtransaction.Create().
				SetWalletID(plb.WalletID).
				SetDateExportLB(time.Now()).
				SetDateGenLB(*plb.DateGenLB).
				SetNoteLB(note).
				SetPointLB(plb.PointLB).
				SetStatusExportLB(true).
				SetLBDate(*plb.DateTime).
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (r ReportService) ExportPointLbCSV(date string) (interface{}, error) {
	if date == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, date)
	if err != nil {
		return nil, err
	}
	endDate := stDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	plb, err := r.connection.Pointpendinglbtransaction.Query().
		Select(pointpendinglbtransaction.FieldWalletID, pointpendinglbtransaction.FieldPointLB, pointpendinglbtransaction.FieldDateExportLB, pointpendinglbtransaction.FieldNoteLB).
		Where(
			pointpendinglbtransaction.LBDateGTE(stDate),
			pointpendinglbtransaction.LBDateLTE(endDate),
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return plb, nil
}

func (r ReportService) GetAllPendingKycList() (interface{}, error) {

	result, err := r.connection.Pendingkyc.Query().
		Order(ent.Asc(pendingkyc.FieldWalletID)).
		All(context.Background())
		//
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

func (r ReportService) GetAllPendingKYCListSearch(walletID, dateGen string) (interface{}, error) {

	options := []predicate.Pendingkyc{
		pendingkyc.WalletIDContains(walletID),
	}

	if dateGen != "" {
		startDate, err := time.Parse(timeCustomLayout, dateGen)
		if err != nil {
			return nil, err
		}
		endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
		options = append(options, pendingkyc.DateGenGTE(startDate), pendingkyc.DateGenLTE(endDate))
	}

	if walletID != "" {
		options = append(options, pendingkyc.WalletIDContains(walletID))
	}

	result, err := r.connection.Pendingkyc.Query().
		//Where(pendingkyc.WalletIDContains(walletID)).
		Where(options...).
		Order(ent.Asc(pendingkyc.FieldWalletID)).
		All(context.Background())
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

func (r ReportService) GetAllPendingLbList() (interface{}, error) {

	result, err := r.connection.Pendingloanbinding.Query().
		Order(ent.Asc(pendingloanbinding.FieldWalletID)).
		All(context.Background())

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

func (r ReportService) GetAllPendingLBListSearch(walletID, dateGen string) (interface{}, error) {

	options := []predicate.Pendingloanbinding{
		pendingloanbinding.WalletIDContains(walletID),
	}

	if dateGen != "" {
		startDate, err := time.Parse(timeCustomLayout, dateGen)
		if err != nil {
			return nil, err
		}
		endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
		options = append(options, pendingloanbinding.DateGenLBGTE(startDate), pendingloanbinding.DateGenLBLTE(endDate))
	}

	if walletID != "" {
		options = append(options, pendingloanbinding.WalletIDContains(walletID))
	}

	result, err := r.connection.Pendingloanbinding.Query().
		//Where(pendingloanbinding.WalletIDContains(walletID)).
		Where(options...).
		Order(ent.Asc(pendingloanbinding.FieldWalletID)).
		All(context.Background())
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return result, nil
}
