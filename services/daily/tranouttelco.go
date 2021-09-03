package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/consumer"
	"go-api-report2/model"
	"time"
)

func (reportService DailyService) GetAisTopup(date string) (model.AisTopup, error) {
	var result model.AisTopup
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisTopup
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("AIS"),
			consumer.PaymentType("TOP_UP_BILLPAY_AIS"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumAisTopupTranOutTelco = 0
		result.CountAisTopupTranOutTelco = 0
		return result, nil
	}

	result.SumAisTopupTranOutTelco = restotal[0].Sum
	result.CountAisTopupTranOutTelco = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetAisPackage(date string) (model.AisPackage, error) {
	var result model.AisPackage
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisPackage
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentChannel("AIS"),
			consumer.PaymentType("PAY_MENT_PACKAGE_AIS"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumAisPackageTranOutTelco = 0
		result.CountAisPackageTranOutTelco = 0
		return result, nil
	}

	result.SumAisPackageTranOutTelco = restotal[0].Sum
	result.CountAisPackageTranOutTelco = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetAisBillpay(date string) (model.AisBillpay, error) {
	var result model.AisBillpay
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisBillpay
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentChannel("AIS"),
			consumer.PaymentType("PAY_MENT_BILLPAYMENT_AIS"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumAisBillpayTranOutTelco = 0
		result.CountAisBillpayTranOutTelco = 0
		return result, nil
	}

	result.SumAisBillpayTranOutTelco = restotal[0].Sum
	result.CountAisBillpayTranOutTelco = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetAisFiber(date string) (model.AisFiber, error) {
	var result model.AisFiber
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.AisFiber
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentChannel("AIS"),
			consumer.PaymentType("PAY_MENT_FIBER_AIS"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumAisFiberTranOutTelco = 0
		result.CountAisFiberTranOutTelco = 0
		return result, nil
	}

	result.SumAisFiberTranOutTelco = restotal[0].Sum
	result.CountAisFiberTranOutTelco = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTrueTopup(date string) (model.TrueTopup, error) {
	var result model.TrueTopup
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TrueTopup
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("TRUEMOVE"),
			consumer.PaymentType("TOP_UP_BILLPAY_TRUEMOVE"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumTrueTopupTranOutTelco = 0
		result.CountTrueTopupTranOutTelco = 0
		return result, nil
	}

	result.SumTrueTopupTranOutTelco = restotal[0].Sum
	result.CountTrueTopupTranOutTelco = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetTruePackage(date string) (model.TruePackage, error) {
	var result model.TruePackage
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.TruePackage
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentChannel("TRUEMOVE"),
			consumer.PaymentType("PAY_MENT_PACKAGE_TRUEMOVE"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumTruePackageTranOutTelco = 0
		result.CountTruePackageTranOutTelco = 0
		return result, nil
	}

	result.SumTruePackageTranOutTelco = restotal[0].Sum
	result.CountTruePackageTranOutTelco = restotal[0].Count
	return result, nil
}

func (reportService DailyService) GetDtacTopup(date string) (model.DtacTopup, error) {
	var result model.DtacTopup
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal []model.DtacTopup
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TOP_UP"),
			consumer.PaymentChannel("DTAC"),
			consumer.PaymentType("TOP_UP_BILLPAY_DTAC"),
			//consumer.TransactionStatus("APPROVE"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		result.SumDtacTopupTranOutTelco = 0
		result.CountDtacTopupTranOutTelco = 0
		return result, nil
	}

	result.SumDtacTopupTranOutTelco = restotal[0].Sum
	result.CountDtacTopupTranOutTelco = restotal[0].Count
	return result, nil
}
