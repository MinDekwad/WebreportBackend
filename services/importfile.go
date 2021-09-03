package services

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/fileimport"
	"time"
)

func (reportService ReportService) CreateFileImport(filename string, filetype string) (*ent.Fileimport, error) {

	return reportService.connection.Fileimport.Create().
		SetFilename(filename).
		SetFiletype(filetype).
		SetImportdate(time.Now()).
		SetStatus("Pending").
		Save(context.Background())

}

func (reportService ReportService) GetWalletFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountWalleteFilename := v[0].Count
	return CountWalleteFilename, nil
}

func (reportService ReportService) GetConsumerFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountConsumerFileName := v[0].Count
	return CountConsumerFileName, nil
}

func (reportService ReportService) GetMerchantFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountMerchantFileName := v[0].Count
	return CountMerchantFileName, nil
}

func (reportService ReportService) GetAgentKycFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountAgentKycFileName := v[0].Count
	return CountAgentKycFileName, nil
}

func (reportService ReportService) GetLoanbindingFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountLoanbindingFileName := v[0].Count
	return CountLoanbindingFileName, nil
}

// new
func (reportService ReportService) GetUserProfileAmloFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountUserProfileAmloFileName := v[0].Count
	return CountUserProfileAmloFileName, nil
}

func (reportService ReportService) GetKycPendingFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountKycPendingFileName := v[0].Count
	return CountKycPendingFileName, nil
}

func (reportService ReportService) GetLbPendingFileName(filename string) (int, error) {
	var v []struct {
		Filename string `json:"filename"`
		Count    int    `json:"count"`
		Filetype string `json:"filetype"`
	}
	err := reportService.connection.Fileimport.Query().
		Where(
			fileimport.Filename(filename),
		).
		GroupBy(fileimport.FieldFiletype).
		Aggregate(ent.Count()).
		Scan(context.Background(), &v)
	if err != nil {
		return 500, nil
	}

	if v == nil || len(v) < 1 {
		return 0, err
	}

	CountLbPendingFileName := v[0].Count
	return CountLbPendingFileName, nil
}
