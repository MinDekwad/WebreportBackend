package daily

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/agenttype"
	"go-api-report2/ent/bulk"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/predicate"
	"go-api-report2/model"
	"math"
	"time"
)

func (reportService DailyService) GetKycCompleteFeeExpenses(date string) (model.KycCompleteFeeExpenses, error) {
	var result model.KycCompleteFeeExpenses
	if date == "" {
		return result, errors.New("Date is required")
	}
	_, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	var restotal []model.KycCompleteFeeExpenses
	err = reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.KYCDate(date),
			agentkyc.KYCStatus("APPROVED"),
		).
		GroupBy(agentkyc.FieldKYCStatus).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal)

	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		return result, nil
	}
	result.CountKycCompleteFeeExpenses = restotal[0].Count
	result.SumKycCompleteFeeExpenses = float64(restotal[0].Count) * 20
	return result, nil
}

func (reportService DailyService) GetPromptPayOutExpenses(date string) (model.PromptPayOutExpenses, error) {
	var result model.PromptPayOutExpenses
	var result1 model.PromptPayOutExpenses
	var result2 model.PromptPayOutExpenses
	if date == "" {
		return result1, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result1, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	var restotal1 []model.PromptPayOutExpenses
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY29_CASH_OUT"),
			consumer.BankCodeNEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal1)
	if err != nil {
		return result1, err
	}

	var CountRow1 int
	if len(restotal1) == 0 {
		CountRow1 = 0
	} else {
		CountRow1 = restotal1[0].Count
	}

	var restotal2 []model.PromptPayOutExpenses
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY29_CASH_OUT"),
			consumer.BankCodeEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Count()).
		Scan(context.Background(), &restotal2)
	if err != nil {
		return result2, err
	}

	var CountRow2 int
	if len(restotal2) == 0 {
		CountRow2 = 0
	} else {
		CountRow2 = restotal2[0].Count
	}

	// var SumCountRow int
	// SumCountRow = CountRow1 + CountRow2
	// result.SumCountRow = SumCountRow

	// var Sum float64
	// Sum = float64(SumCountRow) * 0.9
	// result.Sum = Sum

	var CountPromptPayOutExpenses int
	CountPromptPayOutExpenses = CountRow1 + CountRow2
	result.CountPromptPayOutExpenses = CountPromptPayOutExpenses

	var SumPromptPayOutExpenses float64
	SumPromptPayOutExpenses = float64(CountPromptPayOutExpenses) * 0.9
	result.SumPromptPayOutExpenses = SumPromptPayOutExpenses

	return result, nil
}

func (reportService DailyService) GetBulkTransferFeeExpense(date string) (model.BulkTransferFeeExpense, error) {
	var result model.BulkTransferFeeExpense
	if date == "" {
		return result, errors.New("Date is required")
	}

	stDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	startDate := stDate.Add(8 * time.Hour).Add(00 * time.Minute).Add(00 * time.Second)
	startDateMinus := startDate.AddDate(0, 0, -1)

	endDate := stDate.Add(7 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	restotal, err := reportService.connection.Bulk.Query().Select(bulk.FieldTransfertobankaccount).
		Where(
			bulk.DateTimeGTE(startDateMinus),
			bulk.DateTimeLTE(endDate),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		//result.SumCountTransfertobankaccount = 0
		result.CountTransfertobankaccount = 0
		return result, nil
	}
	//num := 15.00
	result.CountTransfertobankaccount = restotal[0].Transfertobankaccount
	//result.SumCountTransfertobankaccount = restotal[0].Transfertobankaccount * num

	return result, nil
}

func (reportService DailyService) GetPromptpayOutTagThirtyExpense(date string) (model.PromptpayOutTagThirtyExpense, error) {
	var result model.PromptpayOutTagThirtyExpense
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	var restotal []model.PromptpayOutTagThirtyExpense
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("PAYMENT"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("CASH_OUT_PROMPTPAY_TAG30_OFF_US"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal)
	if err != nil {
		return result, err
	}

	if len(restotal) == 0 {
		//result.ResTotal = 0
		// result.CountPromptpayOutTagThirtyExpense = 0
		// result.SumPromptpayOutTagThirtyExpense = 0
		return result, err
	}
	result.CountPromptpayOutTagThirtyExpense = restotal[0].Count
	result.SumPromptpayOutTagThirtyExpense = float64(restotal[0].Count) * 0.25

	return result, nil
}

func (reportService DailyService) GetPromptpayOutICFeeExpense(date string) (model.PromptpayOutICFeeExpense, error) {
	var result model.PromptpayOutICFeeExpense
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}
	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	var restotal1 []model.PromptpayOutICFeeExpense
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY29_CASH_OUT"),
			consumer.BankCodeNEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal1)
	if err != nil {
		return result, err
	}

	var SumData1 float64
	var CountData1 int
	if len(restotal1) == 0 {
		SumData1 = 0
		CountData1 = 0
	} else {
		SumData1 = restotal1[0].Sum
		CountData1 = restotal1[0].Count
	}

	var restotal2 []model.PromptpayOutICFeeExpense
	err = reportService.connection.Consumer.Query().
		Where(
			consumer.DateTimeGTE(startDate),
			consumer.DateTimeLTE(endDate),
			consumer.TransactionType("TRANSFER"),
			consumer.PaymentChannel("PROMPTPAY"),
			consumer.PaymentType("PROMPTPAY29_CASH_OUT"),
			consumer.BankCodeEQ("071"),
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &restotal2)
	if err != nil {
		return result, err
	}

	var SumData2 float64
	var CountData2 int
	if len(restotal2) == 0 {
		SumData2 = 0
		CountData2 = 0
	} else {
		SumData2 = restotal2[0].Sum
		CountData2 = restotal2[0].Count
	}

	var SumData3 float64
	var SumData float64
	SumData3 = SumData1 + SumData2
	SumData = math.Ceil((SumData3 * 1.5) / 100)

	var CountData int
	CountData = CountData1 + CountData2
	if SumData >= 8 {
		result.SumPromptpayOutICFeeExpense = float64(CountData * 8)
	} else if SumData >= 4 {
		result.SumPromptpayOutICFeeExpense = float64(CountData) * float64(SumData)
	} else if SumData <= 3 {
		result.SumPromptpayOutICFeeExpense = float64(CountData * 3)
	}
	result.CountPromptpayOutICFeeExpense = CountData1 + CountData2
	return result, err
}

func (reportService DailyService) GetPromptpayInICFeeExpense(date string) (model.PromptpayInICFeeExpense, error) {
	var result model.PromptpayInICFeeExpense
	if date == "" {
		return result, errors.New("Date is required")
	}

	startDate, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return result, err
	}

	endDate := startDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	offUs, err := reportService.SumPromptPayIn(startDate, endDate, true)
	if err != nil {
		return result, err
	}

	onUs, err := reportService.SumPromptPayIn(startDate, endDate, false)
	if err != nil {
		return result, err
	}

	var sumData3 float64
	var sumData float64
	sumData3 = offUs.Sum + onUs.Sum
	sumData = math.Ceil((sumData3 * 1.5) / 100)

	var countData int
	countData = offUs.Count + onUs.Count
	if sumData >= 10 {
		result.SumPromptpayInICFeeExpense = float64(countData * 10)
	} else if sumData >= 1.6 {
		result.SumPromptpayInICFeeExpense = float64(countData) * float64(sumData)
	} else if sumData <= 1.5 {
		result.SumPromptpayInICFeeExpense = float64(countData) * float64(1.5)
	}
	//result.Count = countData1 + countData2
	result.CountPromptpayInICFeeExpense = countData
	return result, err
}

func (reportService DailyService) SumPromptPayIn(start, end time.Time, offus bool) (model.PromptpayInICFeeExpense, error) {
	conditions := []predicate.Consumer{
		consumer.DateTimeGTE(start),
		consumer.DateTimeLTE(end),
		consumer.TransactionType("TOP_UP"),
		consumer.PaymentChannel("PROMPTPAY"),
		consumer.PaymentType("PROMPTPAY_IN"),
	}

	if offus {
		conditions = append(conditions, consumer.BankCodeNEQ("071"))
	} else {
		conditions = append(conditions, consumer.BankCodeEQ("071"))
	}
	var results []model.PromptpayInICFeeExpense
	err := reportService.connection.Consumer.Query().
		Where(
			conditions...,
		).
		GroupBy(consumer.FieldPaymentType).
		Aggregate(ent.Sum(consumer.FieldTotal), ent.Count()).
		Scan(context.Background(), &results)

	if err != nil {
		return model.PromptpayInICFeeExpense{}, err
	}
	if len(results) == 0 {
		return model.PromptpayInICFeeExpense{}, nil
	}
	return results[0], nil
}
func (reportService DailyService) GetTmdsKycCaseExpense(date string) (model.TmdsKycCaseExpense, error) {
	var result model.TmdsKycCaseExpense
	if date == "" {
		return result, errors.New("Date is required")
	}

	var agents []string
	err := reportService.connection.Agenttype.
		Query().
		Select(agenttype.FieldAgentid).
		Scan(context.Background(), &agents)
	if err != nil {
		return result, err
	}

	count, err := reportService.connection.Agentkyc.
		Query().
		Where(
			agentkyc.AgentIDIn(agents...),
			agentkyc.KYCDateEQ(date),
		).
		Count(context.Background())
	if err != nil {
		return result, err
	}
	// Sum := 0
	Sum := count * -20

	result.CountTmdsKycCaseExpense = count
	result.SumTmdsKycCaseExpense = Sum

	return result, nil
}
