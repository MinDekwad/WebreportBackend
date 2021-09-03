package point

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/configpoint"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/limitpoint"
	"go-api-report2/ent/pointtransaction"
	"go-api-report2/ent/predicate"
	"go-api-report2/model"
	"log"
	"math"
	"time"
)

const (
	dateTimeCustom = "2006-01-02"
)

func (r PointService) CalculatePoint(date string) (interface{}, error) {
	var PointTran = "Error"
	var result model.TransactionPoint

	if date == "" {
		return result, errors.New("Date is required")
	}

	dateTime, err := time.Parse(dateTimeCustom, date)
	if err != nil {
		return result, err
	}
	endDate := dateTime.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	// นับ count consumer transaction table
	countConsumerDate, err := r.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(dateTime),
			consumer.DateTimeLTE(endDate),
		).Count(context.Background())
	if err != nil {
		return nil, err
	}

	// ถ้าไม่มีข้อมูล
	if countConsumerDate == 0 {
		PointTran = "ConsumerNoData"
		return PointTran, err
	}

	// select limit point
	lp, err := r.connection.Limitpoint.Query().Select(limitpoint.FieldLimitPoint).First(context.Background())
	if err != nil {
		return result, err
	}
	limitPoint := lp.LimitPoint

	// นับ count point transaction table
	count, err := r.connection.Pointtransaction.Query().
		Where(
			pointtransaction.Date(dateTime),
		).Count(context.Background())
	if err != nil {
		return result, err
	}

	// ถ้ามีข้อมูล
	if count > 0 {
		PointTran = "Duplicate"
		return PointTran, err
	}

	// select config point ที่ status = active
	configs, err := r.connection.Configpoint.Query().
		Where(
			configpoint.StatusTransaction("Active"),
		).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	// start calculate
	var countPoint int

	for _, config := range configs {

		options := []predicate.Consumer{
			consumer.TransactionType(*config.TransactionType),
			consumer.PaymentChannel(*config.PaymentChannel),
			consumer.PaymentType(*config.PaymentType),
			consumer.DateTimeGTE(dateTime),
			consumer.DateTimeLTE(endDate),
			consumer.TotalGTE(float64(config.Amount)),
		}

		if *config.DummyWallet != "" {
			options = append(options, consumer.ToAccount(*config.DummyWallet))
		}

		consumeTotal, err := r.connection.Consumer.Query().
			Select(consumer.FieldTotal, consumer.FieldFromReference).
			Where(
				options...,
			// consumer.TransactionType(*config.TransactionType),
			// consumer.PaymentChannel(*config.PaymentChannel),
			// consumer.PaymentType(*config.PaymentType),
			// consumer.ToAccount(*config.DummyWallet),
			// consumer.DateTimeGTE(dateTime),
			// consumer.DateTimeLTE(endDate),
			// consumer.TotalGTE(float64(config.Amount)),
			).All(context.Background())
		if err != nil {
			return nil, err
		}
		convertRate := float64(config.Amount)
		//var totalPoint float64

		if consumeTotal != nil {
			for _, consume := range consumeTotal {
				var total float64 = 0
				if consume.Total != nil {
					total = *consume.Total
				}
				point := total / convertRate
				// totalPoint += math.Floor(point)
				totalPoint := math.Floor(point)
				totalPoints := int(totalPoint)

				confPoint := config.Point // query point from db point config
				resPoint := totalPoints * confPoint

				if resPoint > limitPoint {
					resPoint = limitPoint
				}

				countPoint = countPoint + totalPoints

				_, err := r.connection.Pointtransaction.
					Create().
					SetDate(dateTime).
					SetWalletID(*consume.FromReference).
					SetTransactionName(config.TransactionName).
					SetPoint(resPoint).
					SetNillableReference(nil).
					Save(context.Background())
				if err != nil {
					log.Println("Insert point transaction has error ", err)
					return err, nil
				}
				PointTran = "Success"
			}
		}
	}
	return PointTran, err
}

func (r PointService) GetHistoryPoint(startDate, endDate string) (interface{}, error) {

	if startDate == "" || endDate == "" {
		return nil, errors.New("Date is required")
	}

	stDate, err := time.Parse(dateTimeCustom, startDate)
	if err != nil {
		return nil, err
	}
	enDate, err := time.Parse(dateTimeCustom, endDate)

	pt, err := r.connection.Pointtransaction.Query().
		Where(pointtransaction.DateGTE(stDate), pointtransaction.DateLTE(enDate)).
		Order(ent.Asc(pointtransaction.FieldWalletID)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return pt, nil
}
