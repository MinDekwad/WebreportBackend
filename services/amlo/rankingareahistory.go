package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/areahistory"
	"go-api-report2/ent/ranking"
)

func (r AmloService) GetDataRankingAreaHistory(walletID string) (interface{}, error) {

	ah, err := r.connection.Areahistory.Query().Where(areahistory.WalletID(walletID)).Order(ent.Desc(areahistory.FieldDateCalRank, areahistory.FieldID)).All(context.Background())
	if err != nil {
		return nil, err
	}
	return ah, nil
}

func (r AmloService) ExportRankingCSV() (interface{}, error) {

	rk, err := r.connection.Ranking.Query().
		Where(
			ranking.CurrentRank(3), // use this if test finish
			//ranking.Rank(2), // delete this if test finish
		).
		Order(ent.Desc(ranking.FieldName)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return rk, nil
}

func (r AmloService) ExportRanking() (interface{}, error) {

	rk, err := r.connection.Ranking.Query().
		Where(
			ranking.CurrentRank(3),
		).
		Order(ent.Asc(ranking.FieldName)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return rk, nil
}
