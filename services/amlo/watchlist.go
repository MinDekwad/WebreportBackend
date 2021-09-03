package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/watchlist"
	"go-api-report2/ent/watchlisthistory"
)

func (r AmloService) GetAllWatchList() (interface{}, error) {

	wl, err := r.connection.Watchlist.Query().Where(watchlist.IsDeleted(false)).WithRelated().Order(ent.Asc(watchlist.FieldName)).All(context.Background())
	if err != nil {
		return "", err
	}
	return wl, nil
}

func (r AmloService) GetAllWatchListSearch(name, taxID string) (interface{}, error) {

	result, err := r.connection.Watchlist.Query().
		Where(watchlist.IsDeleted(false), watchlist.NameContains(name), watchlist.TaxIDContains(taxID)).WithRelated().
		Order(ent.Asc(watchlist.FieldName)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r AmloService) SaveDelWatchlist(c context.Context, ID int) (interface{}, error) {

	wl, err := r.connection.Watchlist.Query().Where(watchlist.ID(ID)).First(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := r.connection.Watchlist.UpdateOneID(ID).
		SetIsDeleted(true).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = r.connection.Watchlisthistory.
		Update().
		Where(
			watchlisthistory.TaxID(wl.TaxID),
		).
		SetStatusDel(2).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	rk, err := r.connection.Ranking.Query().Where(ranking.TaxID(wl.TaxID)).First(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = r.connection.Ranking.
		Update().
		Where(
			ranking.TaxID(wl.TaxID),
		).
		SetCurrentRank(rk.LastRank).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return res, nil
}
