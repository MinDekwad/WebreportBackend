package statement

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/statementendingbalanc"
	"go-api-report2/model"
	"time"
)

func (reportService StatementService) CreateStatementEndingBalanc(c context.Context, input model.CreateStatementEndingBalanc) (interface{}, error) {
	date, err := time.Parse("2006-01-02 15:04:05", input.StatementDate)
	if err != nil {
		return nil, err
	}
	res, err := reportService.connection.StatementEndingBalanc.Create().
		SetNillableBankID(&input.BankUID).
		SetNillableStatementBalance(&input.StatementBalance).
		SetNillableStatementDate(&date).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (reportService StatementService) GetAllStatementEndingBalance() (interface{}, error) {
	resStatementEndingBalance, err := reportService.connection.StatementEndingBalanc.
		Query().WithBank().
		Order(ent.Desc(statementendingbalanc.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return resStatementEndingBalance, nil
}

func (reportService StatementService) GetEditStatementEndingBalance(Uid int) (interface{}, error) {
	statement, err := reportService.connection.StatementEndingBalanc.
		Query().Where(statementendingbalanc.ID(Uid)).
		WithBank().
		First(context.Background())
	if err != nil {
		return err, nil
	}
	return statement, nil
}

func (reportService StatementService) SaveEditStatementEndingBalance(c context.Context, Uid int, input model.EditStatementEndingBalanc) (interface{}, error) {

	dt, err := time.Parse("2006-01-02 15:04:05", input.StatementDate)
	if err != nil {
		return nil, err
	}

	res, err := reportService.connection.StatementEndingBalanc.
		UpdateOneID(Uid).
		SetBankID(input.BankUID).
		SetStatementBalance(input.StatementBalance).
		SetStatementDate(dt).
		Save(c)

	if err != nil {
		return nil, err
	}
	return res, nil
}
