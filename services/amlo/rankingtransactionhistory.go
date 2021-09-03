package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/transactionfactorhistory"
)

func (r AmloService) GetDataRankingTransactionHistory(walletID string) (interface{}, error) {

	// rkTaxID, err := r.connection.Ranking.Query().Select(ranking.FieldTaxID).Where(ranking.WalletID(walletID)).All(context.Background())
	// if err != nil {
	// 	return nil, err
	// }

	tfh, err := r.connection.Transactionfactorhistory.Query().
		Where(transactionfactorhistory.WalletID(walletID)).
		WithTransactionfactor().
		Order(ent.Desc(transactionfactorhistory.FieldDateCalRank, transactionfactorhistory.FieldID)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return tfh, nil
}
