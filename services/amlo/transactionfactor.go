package amlo

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/transactionfactor"
	"go-api-report2/ent/transactionfactoritem"
	"go-api-report2/model"
	"time"
)

func (reportService AmloService) CreateTransactionFactorItemTmp(c context.Context, input model.TransactionFactorItem) (interface{}, error) {
	var result interface{}

	res, err := reportService.connection.Transactionfactoritem.Create().
		SetTransactionFactorID(0).
		SetMin(input.Min).
		SetMax(input.Max).
		SetRank(input.Rank).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	result = res
	return result, err
}

func (r AmloService) GetAllTrasacItemTmp() (interface{}, error) {
	tftmp, err := r.connection.Transactionfactoritem.Query().Where(transactionfactoritem.TransactionFactorID(0)).Order(ent.Asc(transactionfactoritem.FieldRank)).All(context.Background())
	if err != nil {
		return "", err
	}
	return tftmp, nil
}

func (reportService AmloService) CreateTransactionFactor(c context.Context, input model.TransactionFactor) (interface{}, error) {
	var result interface{}
	result = "Error"
	count, err := reportService.connection.Transactionfactor.Query().
		Where(
			transactionfactor.TransactionFactorName(input.TransactionFactorName),
			transactionfactor.TransactionType(input.TransactionType),
			transactionfactor.PaymentChannel(input.PaymentChannel),
			transactionfactor.PaymentType(input.PaymentType),
		).Count(context.Background())
	if err != nil {
		return nil, err
	}

	if count > 0 {
		result = "Dup"
		return result, err
	}

	now := time.Now()
	numDay := input.Day
	dayMinus := now.AddDate(0, 0, -numDay)
	Date := dayMinus.Format("2006-01-02")

	_, err = reportService.connection.Transactionfactor.Create().
		SetTransactionFactorName(input.TransactionFactorName).
		SetTransactionType(input.TransactionType).
		SetPaymentChannel(input.PaymentChannel).
		SetPaymentType(input.PaymentType).
		SetNumDay(numDay).
		SetDate(Date).
		SetUpdateDate(now).
		SetStatusApprove("Waiting").
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	tf, err := reportService.connection.Transactionfactor.Query().Select(transactionfactor.FieldID).Order(ent.Desc(transactionfactor.FieldID)).First(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = reportService.connection.Transactionfactoritem.
		Update().
		Where(
			transactionfactoritem.TransactionFactorID(0),
		).
		SetTransactionFactorID(tf.ID).
		Save(context.Background())

	return result, nil
}

func (r AmloService) GetAllTransactionfactorList() (interface{}, error) {
	result, err := r.connection.Transactionfactor.Query().All(context.Background())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r AmloService) GetDataTransactionfactor(Uid int) (interface{}, error) {
	data, err := r.connection.Transactionfactor.Query().Where(transactionfactor.ID(Uid)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r AmloService) GetAllTrasacItem(Uid int) (interface{}, error) {
	tftmp, err := r.connection.Transactionfactoritem.Query().Where(transactionfactoritem.TransactionFactorID(Uid)).Order(ent.Asc(transactionfactoritem.FieldRank)).All(context.Background())
	if err != nil {
		return "", err
	}
	return tftmp, nil
}

func (reportService AmloService) CreateTransactionFactorItem(c context.Context, input model.TransactionFactorItem, Uid int) (interface{}, error) {
	var result interface{}
	res, err := reportService.connection.Transactionfactoritem.Create().
		SetTransactionFactorID(Uid).
		SetMin(input.Min).
		SetMax(input.Max).
		SetRank(input.Rank).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	result = res
	return result, err
}

func (r AmloService) SaveEditTransactionFactor(c context.Context, Uid int, input model.TransactionFactor, updateDate string) (interface{}, error) {

	now := time.Now()
	numDay := input.Day
	dayMinus := now.AddDate(0, 0, -numDay)
	Date := dayMinus.Format("2006-01-02")

	result, err := r.connection.Transactionfactor.UpdateOneID(Uid).
		SetTransactionFactorName(input.TransactionFactorName).
		SetTransactionType(input.TransactionType).
		SetPaymentChannel(input.PaymentChannel).
		SetPaymentType(input.PaymentType).
		SetNumDay(numDay).
		SetDate(Date).
		SetUpdateDate(now).
		SetStatusApprove("Waiting").
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r AmloService) GetTransactionItemPer(Uid int) (interface{}, error) {
	result, err := r.connection.Transactionfactoritem.Query().Where(transactionfactoritem.ID(Uid)).Select(transactionfactoritem.FieldMax, transactionfactoritem.FieldMin, transactionfactoritem.FieldRank).First(context.Background())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r AmloService) SaveEditTransactionFactorItemPer(c context.Context, Uid int, input model.TransactionFactorItem) (interface{}, error) {
	res, err := r.connection.Transactionfactoritem.UpdateOneID(Uid).
		SetMin(input.Min).
		SetMax(input.Max).
		SetRank(input.Rank).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r AmloService) SaveDelTransactionFactorItemPer(c context.Context, ID int) (interface{}, error) {
	var result interface{}
	result = "Error"
	err := r.connection.Transactionfactoritem.
		DeleteOneID(ID).
		Exec(context.Background())
	if err != nil {
		result = "Error"
		return nil, err
	}
	result = "Success"
	return result, nil
}

func (r AmloService) SaveDelTransactionFactor(c context.Context, ID int) (interface{}, error) {
	var result interface{}
	result = "Error"

	_, err1 := r.connection.Transactionfactoritem.
		Delete().
		Where(transactionfactoritem.TransactionFactorID(ID)).
		Exec(context.Background())
	if err1 != nil {
		return nil, err1
	}

	err2 := r.connection.Transactionfactor.
		DeleteOneID(ID).
		Exec(context.Background())
	if err2 != nil {
		result = "Error"
		return nil, err2
	}
	result = "Success"
	return result, nil
}

func (r AmloService) SaveDelTransactionFactorItemTmpPer(c context.Context, ID int) (interface{}, error) {
	var result interface{}
	result = "Error"
	err := r.connection.Transactionfactoritemtmp.
		DeleteOneID(ID).
		Exec(context.Background())
	if err != nil {
		result = "Error"
		return nil, err
	}
	result = "Success"
	return result, nil
}

func (r AmloService) SaveApproveTransaction(c context.Context, ID int) (interface{}, error) {

	res, err := r.connection.Transactionfactor.UpdateOneID(ID).
		SetStatusApprove("Approve").
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return res, err
}

func (r AmloService) GetTransactionListWaitingApprove() (interface{}, error) {

	result, err := r.connection.Transactionfactor.Query().
		Where(transactionfactor.StatusApproveEQ("Waiting")).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return result, err
}
