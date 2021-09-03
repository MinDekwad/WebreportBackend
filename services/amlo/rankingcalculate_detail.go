package amlo

import (
	"context"
	"errors"
	"fmt"
	"go-api-report2/ent"
	"go-api-report2/ent/configarea"
	"go-api-report2/ent/configdatecalculaterank"
	"go-api-report2/ent/configoccupation"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/transactionfactor"
	"go-api-report2/ent/transactionfactoritem"
	"go-api-report2/ent/watchlist"
	"go-api-report2/model"
	"sort"
	"strconv"
	"sync"
	"time"
)

func (r AmloService) CalculateRankDetail(rankings []*ent.Ranking, date string) error {

	for _, rk := range rankings {
		cfArea, err := r.connection.Configarea.Query().
			Select(configarea.FieldRank, configarea.FieldID).
			Where(
				//configarea.ZipCode(*rk.ZipCode),
				configarea.ProvinceNameTH(rk.ProvinceNameTH),
			).
			First(context.Background())
		if err != nil {
			return err
		}

		cfOccupation, err := r.connection.Configoccupation.Query().
			Select(configoccupation.FieldRank, configoccupation.FieldID, configoccupation.FieldOccupationName).
			Where(
				configoccupation.OccupationName(rk.OccupationName),
			).
			First(context.Background())
		if err != nil {
			return err
		}

		wl, err := r.connection.Watchlist.Query().
			Where(watchlist.TaxID(rk.TaxID)).WithRelated().
			First(context.Background())

		var notFoundErr *ent.NotFoundError
		if err != nil && !errors.As(err, &notFoundErr) {
			return err
		}

		rankArea, err := strconv.Atoi(cfArea.Rank)
		if err != nil {
			return err
		}

		rankOccu, err := strconv.Atoi(cfOccupation.Rank)
		if err != nil {
			return err
		}

		var rankWatchlist int
		//var watchlistTypeID int
		var watchlistTypeName string
		var isDel bool

		if wl != nil {
			rankWatchlist = wl.RankWatchlist
			//watchlistTypeID = wl.Edges.Related.TypeID
			watchlistTypeName = wl.Edges.Related.TypeName
			isDel = wl.IsDeleted
		}

		sliceRank := []int{rankArea, rankOccu, rankWatchlist}
		sort.Ints(sliceRank)
		rank := sliceRank[len(sliceRank)-1]

		nextDateCalRank, err := r.GetNextDateCal(fmt.Sprintf("%d", rank), date)
		if err != nil {
			r.logger.Errorln("[Service] Get next date calculate ranking error ", err)
			return err
		}

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

		var wg sync.WaitGroup
		if rankWatchlist == 3 {
			wg.Add(1)
			go r.InsertWatchlistHistory(&wg, rank, rk, date, watchlistTypeName, isDel)
		}

		wg.Add(1)
		go r.InsertAreaHistory(&wg, rank, rk, date)
		wg.Add(1)
		go r.InsertOccupationHistory(&wg, rank, rk, date)
		wg.Wait()
		r.logger.Infoln("[Service] call fucntion")
	}

	// transaction factor
	tf, err := r.connection.Transactionfactor.Query().
		Where(transactionfactor.StatusApprove("Approve")).
		All(context.Background())
	if err != nil {
		return err
	}

	for _, tranFactor := range tf {
		numDate := tranFactor.NumDay

		tn := time.Now()

		s := tn.Format("2006-01-02")
		st, err := time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}
		stDay := st.AddDate(0, 0, -numDate)
		stDays := stDay.Format("2006-01-02")
		stDayss, err := time.Parse("2006-01-02", stDays)
		if err != nil {
			return err
		}
		startDate := stDayss.Add(00 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)

		etn := tn.Format("2006-01-02")
		end, err := time.Parse("2006-01-02", etn)
		if err != nil {
			return err
		}
		endDate := end.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

		var restotal []model.ConsumerTranFactor
		err = r.connection.Consumer.Query().
			Select(consumer.FieldTotal, consumer.FieldFromReference).
			Where(
				consumer.TransactionType(tranFactor.TransactionType),
				consumer.PaymentChannel(tranFactor.PaymentChannel),
				consumer.PaymentType(tranFactor.PaymentType),
				consumer.DateTimeGTE(startDate),
				consumer.DateTimeLTE(endDate),
				consumer.TransactionStatus("APPROVE"),
			).
			GroupBy(consumer.FieldFromReference).
			Aggregate(ent.Sum(consumer.FieldTotal)).
			Scan(context.Background(), &restotal)
		if err != nil {
			return err
		}

		for _, consume := range restotal {
			tfi, err := r.connection.Transactionfactoritem.Query().
				Select(transactionfactoritem.FieldRank, transactionfactoritem.FieldTransactionFactorID).
				Where(
					transactionfactoritem.MinGTE(0),
					transactionfactoritem.MinLTE(consume.Sum),
					transactionfactoritem.TransactionFactorID(tranFactor.ID),
				).
				Limit(1).
				Order(ent.Desc(transactionfactoritem.FieldMax)).
				First(context.Background())
			if err != nil {
				return err
			}

			nextDateCalRank, err := r.GetNextDateCal(fmt.Sprintf("%d", tfi.Rank), date)
			//nextDateCalRank, err := r.GetNextDateCalTransaction(fmt.Sprintf("%d", tfi.Rank), date)
			if err != nil {
				r.logger.Errorln("[Service] Get next date calculate ranking transaction factor error ", err)
				return err
			}

			go r.UpdateRankingTransaction(tfi.TransactionFactorID, tfi.Rank, consume.FromReference, nextDateCalRank, date)

			r.logger.Infoln("[Service] call function cal transaction")
		}
	}
	// transaction factor

	return nil
}
func (r AmloService) InsertAreaHistory(wg *sync.WaitGroup, rank int, rk *ent.Ranking, date string) error {
	defer wg.Done()
	_, err := r.connection.Areahistory.
		Create().
		SetWalletID(rk.WalletID).
		SetProvinceNameTH(rk.ProvinceNameTH).
		SetDistrictNameTH(rk.DistrictNameTH).
		SetSubDistrict(rk.SubDistrict).
		SetRankArea(rank).
		SetDateCalRank(time.Now()).
		Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}
func (r AmloService) InsertOccupationHistory(wg *sync.WaitGroup, rank int, rk *ent.Ranking, date string) error {
	defer wg.Done()
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
func (r AmloService) InsertWatchlistHistory(wg *sync.WaitGroup, rank int, rk *ent.Ranking, date string, wltName string, isDel bool) error {
	defer wg.Done()

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

// func (r AmloService) InsertTransactionHistory(wg1 *sync.WaitGroup, tfi *ent.Transactionfactoritem, walletID string, date string, tranFactorID int) error {
// 	defer wg1.Done()
// 	_, err := r.connection.Transactionfactorhistory.
// 		Create().
// 		SetWalletID(walletID).
// 		SetTransactionfactorID(tranFactorID).
// 		SetRankTransactionFactor(tfi.Rank).
// 		SetDateCalRank(time.Now()).
// 		Save(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r AmloService) UpdateRanking(wg *sync.WaitGroup, tfi *ent.Transactionfactoritem, walletID string) error {
	defer wg.Done()
	_, err := r.connection.Ranking.
		Update().
		Where(
			ranking.WalletID(walletID),
			ranking.StatusRanking("Old"),
		).
		SetTransactionFactorRank(tfi.Rank).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r AmloService) GetNextDateCal(numRank, date string) (string, error) {
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

func (r AmloService) UpdateRankingTransaction(TransactionFactorID, Rank int, walletID, nextDateCalRank, date string) error {

	qrk, err := r.connection.Ranking.Query().Where(ranking.WalletID(walletID)).All(context.Background())
	if err != nil {
		return err
	}

	if len(qrk) > 0 {
		for _, rk := range qrk {
			_, err := r.connection.Ranking.
				Update().
				Where(
					ranking.WalletID(rk.WalletID),
				).
				SetLastRank(rk.CurrentRank).
				SetTransactionFactorRank(Rank).
				SetLastDateCalRank(date).
				SetNextDateCalRank(nextDateCalRank).
				Save(context.Background())
			if err != nil {
				return err
			}

			_, err = r.connection.Transactionfactorhistory.
				Create().
				SetWalletID(rk.WalletID).
				SetTransactionfactorID(TransactionFactorID).
				SetRankTransactionFactor(Rank).
				SetDateCalRank(time.Now()).
				Save(context.Background())
			if err != nil {
				return err
			}

		}
	}

	return nil
}
