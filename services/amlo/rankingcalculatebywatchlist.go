package amlo

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/configdatecalculaterank"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/watchlist"
	"time"
)

func (r AmloService) CalculateRankByWl(requestID int) error {
	now := time.Now()
	date := now.Format("2006-01-02")

	wl, err := r.connection.Watchlist.Query().
		Where(watchlist.IsDeleted(false)).
		WithRelated().
		All(context.Background())
	if err != nil {
		return err
	}

	for _, watchlist := range wl {

		if len(watchlist.TaxID) != 13 {
			continue
		}

		rk, err := r.connection.Ranking.Query().
			Where(
				ranking.TaxID(watchlist.TaxID),
			).
			First(context.Background())
			//
		var notFoundErr *ent.NotFoundError
		if err != nil && !errors.As(err, &notFoundErr) {
			return err
		}

		if rk == nil {
			continue
		}

		nextDateCalRank, err := r.GetNextDateCalWatchlist("3", date)
		if err != nil {
			r.logger.Errorln("[Service] Get next date calcuate ranking 3 error ", err)
			return err
		}

		if rk.CurrentRank <= watchlist.RankWatchlist {
			stateCal := rk.StateCal + 1
			update := r.connection.Ranking.UpdateOneID(rk.ID).
				SetLastRank(rk.CurrentRank).
				SetCurrentRank(3).
				SetLastDateCalRank(date).
				SetNextDateCalRank(nextDateCalRank).
				SetStateCal(stateCal)
			if stateCal >= 3 {
				update.SetStatusRanking("Old")
			}

			_, err = update.Save(context.Background())
			if err != nil {
				return err
			}
		}
		r.InsertWatchlistHistoryByManual(3, rk, date, watchlist.Edges.Related.TypeName, false)
		r.logger.Infoln("[Service] call fucntion")
	}

	_, err = r.connection.Fileinsert.UpdateOneID(requestID).
		SetIsSuccess(true).
		SetUpdateDate(time.Now()).
		Save(context.Background())
	r.logger.Infoln("[Service] call update file insert")

	return nil
}

func (r AmloService) InsertWatchlistHistoryByManual(rank int, rk *ent.Ranking, date string, wltName string, isDel bool) error {

	var statusDel int
	statusDel = 1
	if isDel == true {
		statusDel = 2 // ถูกลบ
	}

	_, err := r.connection.Watchlisthistory.
		Create().
		SetName(rk.Name).
		SetTaxID(rk.TaxID).
		SetTypeName(wltName).
		SetRankWatchlist(3).
		SetDateCalRank(time.Now()).
		SetStatusDel(statusDel).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r AmloService) GetNextDateCalWatchlist(numRank, date string) (string, error) {
	dateCalRank, err := r.connection.Configdatecalculaterank.Query().
		Select(configdatecalculaterank.FieldNumDateCalculateRank).
		Where(
			configdatecalculaterank.Rank(numRank),
		).
		First(context.Background())
	if err != nil {
		return "", err
	}

	day, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return "", err
	}

	dayPlus := day.AddDate(0, 0, dateCalRank.NumDateCalculateRank)
	nextDateCalRank := dayPlus.Format("2006-01-02")

	return nextDateCalRank, err
}
