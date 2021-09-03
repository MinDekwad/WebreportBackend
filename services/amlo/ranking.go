package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/ranking"
)

func (r AmloService) GetAllRankList(page int, limit int) (interface{}, int, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 20
	}

	offset := (page - 1) * limit
	result, err := r.connection.Ranking.Query().
		Order(ent.Asc(ranking.FieldName)).
		Limit(limit).Offset(offset).
		All(context.Background())

	if err != nil {
		return "", 0, err
	}

	count, err := r.connection.Ranking.Query().
		Count(context.Background())
		//
	if err != nil {
		return "", 0, err
	}

	return result, count, nil
}

func (r AmloService) GetAllRankListSearch(walletID, name, taxID string, page int, limit int) (interface{}, int, error) {

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 20
	}
	offset := (page - 1) * limit

	result, err := r.connection.Ranking.Query().
		Where(ranking.WalletIDContains(walletID), ranking.NameContains(name), ranking.TaxIDContains(taxID)).
		Order(ent.Asc(ranking.FieldName)).
		Limit(limit).Offset(offset).
		All(context.Background())
	if err != nil {
		return "", 0, err
	}

	count, err := r.connection.Ranking.Query().
		Count(context.Background())
	if err != nil {
		return "", 0, err
	}

	return result, count, nil
}
