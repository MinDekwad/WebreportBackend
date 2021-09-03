package point

import (
	"context"
	"go-api-report2/ent/configpoint"
	"go-api-report2/ent/consumer"
	"go-api-report2/model"
	"time"
)

// const (
// 	dateTimeCustom = "2006-01-02"
// )

func (r PointService) GetTransactionType() (interface{}, error) {

	transactionType, err := r.connection.Consumer.Query().
		Select(consumer.FieldTransactionType).
		GroupBy(consumer.FieldTransactionType).
		Strings(context.Background())
	if err != nil {
		return nil, err
	}
	return transactionType, nil

}

func (r PointService) GetPaymentChannel() (interface{}, error) {

	paymentChannel, err := r.connection.Consumer.Query().
		Select(consumer.FieldPaymentChannel).
		GroupBy(consumer.FieldPaymentChannel).
		Strings(context.Background())
	if err != nil {
		return nil, err
	}
	return paymentChannel, nil

}

func (r PointService) GetPaymentType() (interface{}, error) {

	paymentChannel, err := r.connection.Consumer.Query().
		Select(consumer.FieldPaymentType).
		GroupBy(consumer.FieldPaymentType).
		Strings(context.Background())
	if err != nil {
		return nil, err
	}
	return paymentChannel, nil

}

func (reportService PointService) CreateConfigPoint(c context.Context, input model.CreateConfigPoint, today string) (interface{}, error) {
	var result interface{}
	count, err := reportService.connection.Configpoint.Query().
		Where(
			configpoint.TransactionName(input.TransactionName),
			configpoint.TransactionType(input.TransactionType),
			configpoint.PaymentChannel(input.PaymentChannel),
			configpoint.PaymentType(input.PaymentType),
		).Count(context.Background())
	if err != nil {
		return nil, err
	}

	if count > 0 {
		result = "Dup"
		return result, err
	}

	dateTime, err := time.Parse(dateTimeCustom, today)
	if err != nil {
		return nil, err
	}

	var dateExpire time.Time
	numExpire := input.Expire
	if numExpire == 0 {
		dateExpire, err = time.Parse(dateTimeCustom, "2001-01-01")
		if err != nil {
			return nil, err
		}
	} else {
		dateExpire = dateTime.AddDate(0, 0, numExpire)
	}

	res, err := reportService.connection.Configpoint.Create().
		SetNillableTransactionName(&input.TransactionName).
		SetNillableTransactionType(&input.TransactionType).
		SetNillablePaymentChannel(&input.PaymentChannel).
		SetNillablePaymentType(&input.PaymentType).
		SetNillableDummyWallet(&input.DummyWallet).
		SetNillableAmount(&input.Amount).
		SetNillablePoint(&input.Point).
		SetNillableExpire(&input.Expire).
		SetUpdateDate(dateTime).
		SetExpireDate(dateExpire).
		SetStatusTransaction(input.StatusTransaction).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	result = res
	return result, err
}

func (r PointService) GetAllPointTransactionList() (interface{}, error) {
	result, err := r.connection.Configpoint.Query().All(context.Background())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r PointService) GetDataPointTransaction(Uid int) (interface{}, error) {
	data, err := r.connection.Configpoint.Query().Where(configpoint.ID(Uid)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r PointService) SaveEditConfigPoint(c context.Context, Uid int, input model.EditConfigPoint, updateDate string) (interface{}, error) {

	dateTime, err := time.Parse(dateTimeCustom, updateDate)
	if err != nil {
		return nil, err
	}

	var dateExpire time.Time
	numExpire := input.Expire
	if numExpire == 0 {
		dateExpire, err = time.Parse(dateTimeCustom, "2001-01-01")
		if err != nil {
			return nil, err
		}
	} else {
		dateExpire = dateTime.AddDate(0, 0, numExpire)
	}

	res, err := r.connection.Configpoint.UpdateOneID(Uid).
		SetNillableTransactionName(&input.TransactionName).
		SetNillableTransactionType(&input.TransactionType).
		SetNillablePaymentChannel(&input.PaymentChannel).
		SetNillablePaymentType(&input.PaymentType).
		SetNillableDummyWallet(&input.DummyWallet).
		SetNillableAmount(&input.Amount).
		SetNillablePoint(&input.Point).
		SetNillableExpire(&input.Expire).
		SetUpdateDate(dateTime).
		SetExpireDate(dateExpire).
		SetStatusTransaction(input.StatusTransaction).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}
