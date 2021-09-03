package amlo

import (
	"context"
	"fmt"
	"go-api-report2/ent"
	"go-api-report2/ent/configdatecalculaterank"
	"go-api-report2/ent/configoccupation"
	"go-api-report2/ent/ranking"
	"strconv"
	"time"
)

func (r AmloService) CalculateRankByOccu(requestID int) error {
	now := time.Now()
	date := now.Format("2006-01-02")

	rankings, err := r.connection.Ranking.Query().
		Where(
			ranking.OccupationNameNEQ("Other"),
		).
		All(context.Background())
	if err != nil {
		return err
	}

	for _, rk := range rankings {
		cfOccupation, err := r.connection.Configoccupation.Query().
			Select(configoccupation.FieldRank, configoccupation.FieldID, configoccupation.FieldOccupationName).
			Where(
				configoccupation.OccupationName(rk.OccupationName),
			).
			First(context.Background())
		if err != nil {
			return err
		}

		rankOccu, err := strconv.Atoi(cfOccupation.Rank)
		if err != nil {
			return err
		}

		// sliceRank := []int{rankOccu}
		// sort.Ints(sliceRank)
		// rank := sliceRank[len(sliceRank)-1]

		nextDateCalRank, err := r.GetNextDateCal(fmt.Sprintf("%d", rankOccu), date)
		if err != nil {
			r.logger.Errorln("[Service] Get next date calculate ranking 3 error ", err)
			return err
		}

		if rk.CurrentRank <= rankOccu {
			stateCal := rk.StateCal + 1
			update := r.connection.Ranking.UpdateOneID(rk.ID).
				SetLastRank(rk.CurrentRank).
				SetCurrentRank(rankOccu).
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

		//var wg sync.WaitGroup

		//wg.Add(1)
		// go r.InsertOccupationHistoryByManual(&wg, rank, rk, date)
		//wg.Wait()
		go r.InsertOccupationHistoryByManual(rankOccu, rk, date)
		r.logger.Infoln("[Service] call fucntion")
	}

	_, err = r.connection.Fileinsert.UpdateOneID(requestID).
		SetIsSuccess(true).
		SetUpdateDate(time.Now()).
		Save(context.Background())
	r.logger.Infoln("[Service] call update file insert")
	return nil
}

// func (r AmloService) InsertOccupationHistoryByManual(wg *sync.WaitGroup, rank int, rk *ent.Ranking, date string) error {
func (r AmloService) InsertOccupationHistoryByManual(rank int, rk *ent.Ranking, date string) error {
	//defer wg.Done()
	_, err := r.connection.Occupationhistory.
		Create().
		SetWalletID(rk.WalletID).
		SetOccupationName(rk.OccupationName).
		SetRankOccupation(rank).
		SetDateCalRank(time.Now()).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r AmloService) GetNextDateCalOccupation(numRank, date string) (string, error) {
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
