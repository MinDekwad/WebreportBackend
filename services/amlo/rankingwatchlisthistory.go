package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/watchlisthistory"
)

func (r AmloService) GetDataRankingWatchlistHistory(walletID string) (interface{}, error) {
	rkTaxID, err := r.connection.Ranking.Query().Select(ranking.FieldTaxID).Where(ranking.WalletID(walletID)).All(context.Background())
	if err != nil {
		return nil, err
	}

	var wlData interface{}
	for _, rkt := range rkTaxID {
		wl, err := r.connection.Watchlisthistory.Query().Where(watchlisthistory.TaxID(rkt.TaxID)).Order(ent.Desc(watchlisthistory.FieldDateCalRank, watchlisthistory.FieldID)).All(context.Background())
		if err != nil {
			return nil, err
		}
		wlData = wl
	}

	return wlData, nil
}
