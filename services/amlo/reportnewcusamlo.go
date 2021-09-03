package amlo

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/ranking"
)

func (r AmloService) DownloadAmloCustomer(date, statusType, activedate string) (interface{}, error) {

	var result interface{}

	if date == "" && activedate == "" {
		return nil, errors.New("Date is required")
	}

	if date != "" {
		rk, err := r.connection.Ranking.Query().
			Where(
				ranking.NextDateCalRank(date),
				ranking.StatusRanking(statusType)).
			Order(ent.Asc(ranking.FieldWalletID)).
			All(context.Background())
		if err != nil {
			return nil, err
		}
		result = rk
	}

	if activedate != "" {

		rk, err := r.connection.Ranking.Query().
			Where(
				ranking.LastDateCalRank(activedate),
				ranking.StatusRanking(statusType)).
			Order(ent.Asc(ranking.FieldWalletID)).
			All(context.Background())
		if err != nil {
			return nil, err
		}
		result = rk
	}
	return result, nil
}
