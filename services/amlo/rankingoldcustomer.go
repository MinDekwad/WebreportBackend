package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/ranking"
)

func (r AmloService) GetAllOldCustomerList() (interface{}, error) {
	//var result interface{}
	var result []interface{}
	rk, err := r.connection.Ranking.Query().
		Where(
			ranking.StatusRanking("Old"),
		).
		Order(ent.Asc(ranking.FieldWalletID)).
		All(context.Background())
	if err != nil {
		return "", err
	}

	i := 0
	for _, ranking := range rk {
		ak, err := r.connection.Agentkyc.Query().
			Where(agentkyc.Consumerwalletid(ranking.WalletID)).
			All(context.Background())
		if err != nil {
			return nil, err
		}
		for _, agentk := range ak {
			//fmt.Println(ranking.WalletID, *agentk.KYCDate)
			result = append(result, ranking.WalletID, *agentk.KYCDate)
			//result =
		}
		i++
	}

	// rk, err := r.connection.Ranking.Query().
	// 	Order(func(s *sql.Selector) {
	// 		t := sql.Table(agentkyc.Table)
	// 		s.Join(t).On(s.C(ranking.WalletID), t.C(agentkyc.FieldConsumerwalletid))
	// 		s.OrderBy(s.C(ranking.FieldWalletID))
	// 	}).
	// 	Select(agentkyc.FieldKYCDate).
	// 	Strings(context.Background())
	// if err != nil {
	// 	return "", err
	// }

	return result, nil
}
