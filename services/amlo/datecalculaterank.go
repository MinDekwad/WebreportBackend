package amlo

import (
	"context"
	"go-api-report2/ent/configdatecalculaterank"
	"go-api-report2/model"
	"strconv"
	"time"
)

func (r AmloService) GetAllDateCalculateRankList() (interface{}, error) {
	result, err := r.connection.Configdatecalculaterank.Query().All(context.Background())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r AmloService) GetDataDateCalculateRank(ID int) (interface{}, error) {
	data, err := r.connection.Configdatecalculaterank.Query().Where(configdatecalculaterank.IDEQ(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r AmloService) SaveEditConfigDateCalculateRank(c context.Context, ID int, input model.ConfigDateCalculateRank, updateDate string) (interface{}, error) {
	dateTime, err := time.Parse(dateTimeCustom, updateDate)
	if err != nil {
		return nil, err
	}

	res, err := r.connection.Configdatecalculaterank.UpdateOneID(ID).
		SetNumDateCalculateRankTmp(input.NumDateCalculateRankTmp).
		SetUpdateDate(dateTime).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r AmloService) SaveApproveConfigDateCalculateRank(c context.Context, ID int) (interface{}, error) {

	numDateCalculateRankTmp, err := r.connection.Configdatecalculaterank.Query().Select(configdatecalculaterank.FieldNumDateCalculateRankTmp).Where(configdatecalculaterank.IDEQ(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}

	NumDateCalculateRank, err := strconv.Atoi(numDateCalculateRankTmp.NumDateCalculateRankTmp)
	if err != nil {
		return nil, err
	}

	// if rankTmp.RankTmp == "-" {
	// 	result = "NoTmp"
	// 	return result, err
	// } else {
	res, err := r.connection.Configdatecalculaterank.UpdateOneID(ID).
		SetNumDateCalculateRank(NumDateCalculateRank).
		SetNumDateCalculateRankTmp("-").
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	//result = res
	//}
	return res, err
}
