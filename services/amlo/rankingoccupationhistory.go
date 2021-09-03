package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/occupationhistory"
)

func (r AmloService) GetDataRankingOccupationHistory(walletID string) (interface{}, error) {
	oh, err := r.connection.Occupationhistory.Query().Where(occupationhistory.WalletID(walletID)).Order(ent.Desc(occupationhistory.FieldDateCalRank, occupationhistory.FieldID)).All(context.Background())
	if err != nil {
		return nil, err
	}
	return oh, nil
}
