package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/configoccupation"
	"go-api-report2/model"
	"time"
)

func (r AmloService) ServiceCreateConfigOccupation(c context.Context, input model.CreateConfigOccupation, today string) (interface{}, error) {
	var result interface{}
	count, err := r.connection.Configoccupation.Query().
		Where(
			configoccupation.OccupationName(input.OccupationName),
		).Count(context.Background())
	if err != nil {
		return nil, err
	}
	if count > 0 {
		result = "Dup"
		return result, err
	}

	dateTime, err := time.Parse(dateTimeCustom, today)
	if err != nil {
		return nil, err
	}

	res, err := r.connection.Configoccupation.Create().
		SetOccupationName(input.OccupationName).
		SetRank("-").
		SetRankTmp(input.Rank).
		SetUpdateDate(dateTime).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	result = res
	return result, err
}

func (r AmloService) GetAllOccupationList() (interface{}, error) {
	result, err := r.connection.Configoccupation.Query().All(context.Background())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r AmloService) SaveDelConfigOccupation(c context.Context, ID int) (interface{}, error) {
	var result interface{}
	err := r.connection.Configoccupation.
		DeleteOneID(ID).
		Exec(context.Background())
	if err != nil {
		result = "Error"
		return nil, err
	}
	result = "Success"
	return result, nil
}

func (r AmloService) GetDataOccupation(ID int) (interface{}, error) {
	data, err := r.connection.Configoccupation.Query().Where(configoccupation.IDEQ(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r AmloService) SaveEditConfigOccupation(c context.Context, ID int, input model.CreateConfigOccupation, updateDate string) (interface{}, error) {
	dateTime, err := time.Parse(dateTimeCustom, updateDate)
	if err != nil {
		return nil, err
	}

	res, err := r.connection.Configoccupation.UpdateOneID(ID).
		SetOccupationName(input.OccupationName).
		SetRankTmp(input.Rank).
		SetUpdateDate(dateTime).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r AmloService) SaveApproveConfigOccupation(c context.Context, ID int, userData string) (interface{}, error) {

	rankTmp, err := r.connection.Configoccupation.Query().Select(configoccupation.FieldRankTmp).Where(configoccupation.IDEQ(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}

	// if rankTmp.RankTmp == "-" {
	// 	result = "NoTmp"
	// 	return result, err
	// } else {
	res, err := r.connection.Configoccupation.UpdateOneID(ID).
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

func (r AmloService) GetAllOccupationListSearch(occupationName, rank string) (interface{}, error) {
	var res interface{}

	if rank == "" {
		result, err := r.connection.Configoccupation.Query().
			Where(configoccupation.OccupationNameContains(occupationName)).
			Order(ent.Asc(configoccupation.FieldOccupationName)).
			All(context.Background())
		if err != nil {
			return nil, err
		}
		res = result
	} else {
		result, err := r.connection.Configoccupation.Query().
			Where(configoccupation.OccupationNameContains(occupationName), configoccupation.Rank(rank)).
			Order(ent.Asc(configoccupation.FieldOccupationName)).
			All(context.Background())
		if err != nil {
			return nil, err
		}
		res = result
	}

	// if rank == "" && occupationName != "" {
	// 	result, err := r.connection.Configoccupation.Query().Where(configoccupation.OccupationNameContains(occupationName)).All(context.Background())
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	res = result
	// }
	// if rank == "" && occupationName == "" {
	// 	result, err := r.connection.Configoccupation.Query().All(context.Background())
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	res = result
	// }
	// if rank != "" && occupationName != "" {
	// 	result, err := r.connection.Configoccupation.Query().
	// 		Where(configoccupation.OccupationName(occupationName), configoccupation.RankEQ(rank)).
	// 		All(context.Background())
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	res = result
	// }
	// if rank != "" && occupationName == "" {
	// 	result, err := r.connection.Configoccupation.Query().
	// 		Where(configoccupation.RankEQ(rank)).
	// 		All(context.Background())
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	res = result
	// }
	return res, nil
}

func (r AmloService) GetOccuListWaitingApprove() (interface{}, error) {

	result, err := r.connection.Configoccupation.Query().
		Where(configoccupation.RankTmpNEQ("-")).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return result, err
}
