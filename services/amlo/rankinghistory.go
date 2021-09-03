package amlo

import (
	"context"
	"go-api-report2/ent/ranking"
)

func (r AmloService) GetDataRankingHistory(walletID string) (interface{}, error) {

	rk, err := r.connection.Ranking.Query().Where(ranking.WalletID(walletID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return rk, nil
}
