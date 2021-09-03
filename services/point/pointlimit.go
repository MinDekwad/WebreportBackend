package point

import (
	"context"
	"go-api-report2/ent/limitpoint"
	"go-api-report2/model"
)

func (r PointService) GetAllLimitPoint() (interface{}, error) {
	result, err := r.connection.Limitpoint.Query().Select(limitpoint.FieldLimitPoint).First(context.Background())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r PointService) SaveEditLimitPoint(c context.Context, Uid int, input model.LimitPoint) (interface{}, error) {
	res, err := r.connection.Limitpoint.UpdateOneID(Uid).
		SetLimitPoint(input.LimitPoint).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}
