package bulk

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/bulk"
	"go-api-report2/model"
	"time"
)

// type BulkService struct {
// 	connection *ent.Client
// 	logger     *log.Logger
// }

func (reportService BulkService) CreateBulkTransaction(c context.Context, input model.CreateBulkTransaction) (interface{}, error) {
	date, err := time.Parse("2006-01-02 15:04:05", input.DateTime)
	if err != nil {
		return nil, err
	}
	resBulkTransaction, err := reportService.connection.Bulk.Create().
		SetBulkCreditSameday(input.BulkCreditSameday).
		SetBulkCreditSamedayFee(input.BulkCreditSamedayFee).
		SetTransfertobankaccount(float64(input.Transfertobankaccount)).
		SetDateTime(date).
		Save(context.Background())
	// resBulkTransaction, err := reportService.connection.Bulk.Create().
	// 	SetBulkCreditSameday(input.BulkCreditSameday).
	// 	SetBulkCreditSamedayFee(input.BulkCreditSamedayFee).
	// 	SetTransfertobankaccount(float64(input.Transfertobankaccount)).
	// 	SetDateTime(date).
	// 	Save(c)
	if err != nil {
		return nil, err
	}

	//input.Uid = resBulkTransaction.ID
	return resBulkTransaction, nil
}

func (reportService BulkService) GetAllBulkTransaction() (interface{}, error) {
	resBulkTransaction, err := reportService.connection.Bulk.
		Query().
		Order(ent.Desc(bulk.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return resBulkTransaction, nil
}

func (reportService BulkService) GetEditBulkTransaction(Uid int) (interface{}, error) {
	res, err := reportService.connection.Bulk.Query().Select(bulk.FieldBulkCreditSameday, bulk.FieldBulkCreditSamedayFee, bulk.FieldDateTime, bulk.FieldTransfertobankaccount).
		Where(
			bulk.ID(Uid),
		).
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (reportService BulkService) SaveEditBulkTransaction(c context.Context, Uid int, input model.EditBulkTransaction) (interface{}, error) {
	dt, err := time.Parse("2006-01-02 15:04:05", input.DateTime)
	if err != nil {
		return nil, err
	}

	res, err := reportService.connection.Bulk.
		UpdateOneID(Uid).
		SetBulkCreditSameday(input.BulkCreditSameday).
		SetBulkCreditSamedayFee(input.BulkCreditSamedayFee).
		SetTransfertobankaccount(float64(input.Transfertobankaccount)).
		SetDateTime(dt).
		Save(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
