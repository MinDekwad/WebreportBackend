package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/configoccupation"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/userwallet"
	"time"
)

func (r AmloService) UpdateRankingOccupation() (interface{}, error) {

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation) // delete this if test finish
	lastOfMonthDay := firstOfMonth.AddDate(0, 1, -1)
	lastOfMonth := lastOfMonthDay.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	rw, err := r.connection.Userwallet.Query().
		Order(ent.Asc(userwallet.FieldWalletid)).
		Select(userwallet.FieldOccupationId, userwallet.FieldWalletid).
		Where(
			userwallet.UpdateDateGTE(firstOfMonth),
			userwallet.UpdateDateLTE(lastOfMonth),
			userwallet.OccupationIdNEQ(0),
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	var StatusSuccess string
	StatusSuccess = "Error"
	for _, reportWallet := range rw {

		confOccu, err := r.connection.Configoccupation.Query().
			Order(ent.Asc(configoccupation.FieldID)).
			Select(configoccupation.FieldOccupationName).
			Where(
				configoccupation.ID(reportWallet.OccupationId),
			).
			First(context.Background())
		if err != nil {
			return nil, err
		}

		_, err = r.connection.Ranking.
			Update().
			Where(
				ranking.WalletID(reportWallet.Walletid),
			).
			SetOccupationName(confOccu.OccupationName).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		StatusSuccess = "Success"
	}

	return StatusSuccess, nil
}
