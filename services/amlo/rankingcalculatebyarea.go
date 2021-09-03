package amlo

import (
	"context"
	"fmt"
	"go-api-report2/ent"
	"go-api-report2/ent/configarea"
	"go-api-report2/ent/configdatecalculaterank"
	"go-api-report2/ent/ranking"
	"strconv"

	// "go-api-report2/ent/configoccupation"
	// "go-api-report2/ent/ranking"
	// "strconv"
	"time"
)

func (r AmloService) CalculateRankByArea(name string, requestID int) error {

	now := time.Now()
	date := now.Format("2006-01-02")

	rankings, err := r.connection.Ranking.Query().
		Where(
			ranking.ProvinceNameTH(name),
		).
		All(context.Background())

	if err != nil {
		return err
	}

	areaMap := map[string]int{}
	areas, err := r.connection.Configarea.Query().
		Where(
			configarea.ProvinceNameTH(name),
		).
		All(context.Background())
	if err != nil {
		return err
	}

	for _, a := range areas {
		rank, err := strconv.Atoi(a.Rank)
		if err != nil {
			return err
		}
		key := fmt.Sprintf("%s:%s:%s", a.ProvinceNameTH, a.DistrictNameTH, a.SubDistrictNameTH)
		areaMap[key] = rank
	}

	for _, rk := range rankings {
		key := fmt.Sprintf("%s:%s:%s", rk.ProvinceNameTH, rk.DistrictNameTH, rk.SubDistrict)
		rank, ok := areaMap[key]
		if !ok {
			r.logger.Errorln("Area not found", rk.ID, "key", key)
			continue
		}

		nextDateCalRank, err := r.GetNextDateCal(fmt.Sprint(rank), date)
		if err != nil {
			r.logger.Errorln("[Service] Get next date calculate ranking 3 error ", err)
			return err
		}

		if rk.CurrentRank <= rank {
			stateCal := rk.StateCal + 1
			update := r.connection.Ranking.UpdateOneID(rk.ID).
				SetLastRank(rk.CurrentRank).
				SetCurrentRank(rank).
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

		r.InsertAreaHistoryByManual(rank, rk, date)
	}
	// end loop

	_, err = r.connection.Fileinsert.UpdateOneID(requestID).
		SetIsSuccess(true).
		SetUpdateDate(time.Now()).
		Save(context.Background())

	return nil
}

func (r AmloService) InsertAreaHistoryByManual(rank int, rk *ent.Ranking, date string) error {

	//fmt.Println(rk.WalletID)

	_, err := r.connection.Areahistory.
		Create().
		SetWalletID(rk.WalletID).
		SetProvinceNameTH(rk.ProvinceNameTH).
		SetDistrictNameTH(rk.DistrictNameTH).
		//SetDistrictNameEN(rk.DistrictNameEN).
		SetSubDistrict(rk.SubDistrict).
		SetRankArea(rank).
		SetDateCalRank(time.Now()).
		Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r AmloService) GetNextDateCalArea(numRank, date string) (string, error) {
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
