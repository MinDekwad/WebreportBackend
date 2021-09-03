package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/configarea"
	"go-api-report2/model"
	"time"
)

func (r AmloService) GetProvinceNameTH() (interface{}, error) {

	provinceData, err := r.connection.Configarea.Query().
		Order(ent.Asc(configarea.FieldProvinceNameTH)).
		Select(
			configarea.FieldProvinceNameTH,
		).
		GroupBy(configarea.FieldProvinceNameTH).
		Strings(context.Background())
	if err != nil {
		return nil, err
	}
	return provinceData, nil

}

func (r AmloService) GetAllAreaListSearch(provinceNameTH string) (interface{}, error) {

	result, err := r.connection.Configarea.Query().
		Where(configarea.ProvinceNameTH(provinceNameTH)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r AmloService) GetDataArea(ID int) (interface{}, error) {
	data, err := r.connection.Configarea.Query().Where(configarea.IDEQ(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r AmloService) SaveEditConfigArea(c context.Context, ID int, input model.ConfigArea, updateDate string) (interface{}, error) {
	dateTime, err := time.Parse(dateTimeCustom, updateDate)
	if err != nil {
		return nil, err
	}

	res, err := r.connection.Configarea.UpdateOneID(ID).
		SetRankTmp(input.Rank).
		SetUpdateDate(dateTime).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r AmloService) SaveApproveConfigArea(c context.Context, ID int, userData string) (interface{}, error) {
	//var result interface{}
	rankTmp, err := r.connection.Configarea.Query().Select(configarea.FieldRankTmp).Where(configarea.IDEQ(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}

	// if rankTmp.RankTmp == "-" {
	// 	result = "NoTmp"
	// 	return result, err
	// } else {
	res, err := r.connection.Configarea.UpdateOneID(ID).
		SetRank(rankTmp.RankTmp).
		SetRankTmp("-").
		SetApproveDate(time.Now()).
		SetApproveBy(userData).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	//result = res
	//}
	return res, err
}

func (r AmloService) GetAllAreaList() (interface{}, error) {

	result, err := r.connection.Configarea.Query().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r AmloService) GetAreaListWaitingApprove() (interface{}, error) {

	result, err := r.connection.Configarea.Query().
		Where(configarea.RankTmpNEQ("-")).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return result, err
}
