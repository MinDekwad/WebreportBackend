package amlo

import (
	"context"
	"go-api-report2/ent/fileinsert"
	"go-api-report2/model"
)

func (reportService AmloService) GetAreaStatus(ID int) (model.Fileinsert, error) {
	var result model.Fileinsert

	res, err := reportService.connection.Fileinsert.Query().Select(fileinsert.FieldIsSuccess).
		Where(fileinsert.ID(ID)).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(res) == 0 {
		result.FileinsertStatus = ""
		//result.MerchantFileimportID = 0
		return result, nil
	}
	result.FileinsertStatus = res[0].IsSuccess
	return result, nil
}

func (reportService AmloService) GetOccuStatus(ID int) (model.Fileinsert, error) {
	var result model.Fileinsert

	res, err := reportService.connection.Fileinsert.Query().Select(fileinsert.FieldIsSuccess).
		Where(fileinsert.ID(ID)).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(res) == 0 {
		result.FileinsertStatus = ""
		//result.MerchantFileimportID = 0
		return result, nil
	}
	result.FileinsertStatus = res[0].IsSuccess
	return result, nil
}

func (reportService AmloService) GetWlStatus(ID int) (model.Fileinsert, error) {
	var result model.Fileinsert

	res, err := reportService.connection.Fileinsert.Query().Select(fileinsert.FieldIsSuccess).
		Where(fileinsert.ID(ID)).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(res) == 0 {
		result.FileinsertStatus = ""
		//result.MerchantFileimportID = 0
		return result, nil
	}
	result.FileinsertStatus = res[0].IsSuccess
	return result, nil
}

func (reportService AmloService) GetTFStatus(ID int) (model.Fileinsert, error) {
	var result model.Fileinsert

	res, err := reportService.connection.Fileinsert.Query().Select(fileinsert.FieldIsSuccess).
		Where(fileinsert.ID(ID)).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(res) == 0 {
		result.FileinsertStatus = ""
		//result.MerchantFileimportID = 0
		return result, nil
	}
	result.FileinsertStatus = res[0].IsSuccess
	return result, nil
}
