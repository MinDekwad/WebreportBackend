package amlo

import (
	"context"
	"fmt"
	"go-api-report2/ent"
	"go-api-report2/ent/configdatecalculaterank"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/transactionfactor"
	"go-api-report2/ent/transactionfactoritem"
	"go-api-report2/model"
	"time"
)

func (r AmloService) CalculateRankByTF(requestID int) error {
	// func (r AmloService) CalculateRankByTransaction() (string, error) {

	now := time.Now()
	date := now.Format("2006-01-02")

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

		// fmt.Println("statdate : ", startDate)
		// fmt.Println("enddate : ", endDate)

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

			nextDateCalRank, err := r.GetNextDateCalTransaction(fmt.Sprintf("%d", tfi.Rank), date)
			if err != nil {
				r.logger.Errorln("[Service] Get next date calculate ranking 3 error ", err)
				return err
			}

			go r.UpdateRankingByTransactionManual(tfi.TransactionFactorID, tfi.Rank, consume.FromReference, nextDateCalRank, date)

			r.logger.Infoln("[Service] call function cal transaction")
		}
	}
	_, err = r.connection.Fileinsert.UpdateOneID(requestID).
		SetIsSuccess(true).
		SetUpdateDate(time.Now()).
		Save(context.Background())
	r.logger.Infoln("[Service] call update file insert")
	return nil
}

func (r AmloService) UpdateRankingByTransactionManual(TransactionFactorID, Rank int, walletID, nextDateCalRank, date string) error {

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

func (r AmloService) GetNextDateCalTransaction(numRank, date string) (string, error) {
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
