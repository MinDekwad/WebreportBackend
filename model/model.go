package model

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

type Category string

const DefaultCategory = CategoryNano

const (
	CategoryNano  Category = "nano"
	CategoryEvent Category = "event"
	CategoryOther Category = "other"
)

func (c Category) String() string {
	return string(c)
}

type NewWallet struct {
	ResTotal       interface{} `json:"restotal,omitempty"`
	Count          int         `json:"count,omitempty"`
	KYCStatus      string      `json:"kycstatus,omitempty"`
	NewWallet      interface{} `json:"new_wallet,omitempty"`
	CountNewWallet interface{} `json:"count_new_wallet,omitempty"`
}
type NoOfWallet struct {
	Report interface{} `json:"report,omitempty"`
	//CountNoOfWallet int         `json:"countnoofwallet,omitempty"`

	ResTotal        interface{} `json:"restotal,omitempty"`
	Count           int         `json:"count,omitempty"`
	WalletTypeName  string      `json:"wallettypename,omitempty"`
	NoOfWallet      interface{} `json:"no_of_wallet,omitempty"`
	CountNoOfWallet interface{} `json:"count_no_of_wallet,omitempty"`
}

type WalletBalance struct {
	Report interface{} `json:"report,omitempty"`
	//CountWalletBalance int         `json:"countwalletbalance,omitempty"`

	//ResTotal       interface{} `json:"restotal,omitempty"`
	Sum            float64     `json:"sum,omitempty"`
	WalletTypeName string      `json:"wallettypename,omitempty"`
	WalletBalance  interface{} `json:"wallet_balance,omitempty"`
	//ResTotal       float64     `json:"restotal,omitempty"`
	CountWalletBalance interface{} `json:"count_wallet_balance,omitempty"`
}

type DiffWalletBalance struct {
	Report1 interface{} `json:"report1,omitempty"`
	Report2 interface{} `json:"report2,omitempty"`
	//CountWalletBalance int         `json:"countwalletbalance,omitempty"`
	Total interface{} `json:"total,omitempty"`
	//DiffWalletBalance interface{} `json:"diff_wallet_balance,omitempty"`
	DiffWalletBalance float64 `json:"diff_wallet_balance,omitempty"`

	Sum            float64     `json:"sum,omitempty"`
	WalletTypeName string      `json:"wallettypename,omitempty"`
	ResTotal       interface{} `json:"restotal,omitempty"`
	ResTotal1      interface{} `json:"restotal1,omitempty"`
	ResTotal2      interface{} `json:"restotal2,omitempty"`

	Result               interface{} `json:"result,omitempty"`
	SumDiffWalletBalance interface{} `json:"sum_diff_wallet_balance,omitempty"`

	SumPrevDay interface{} `json:"sum_prev_day"`
	SumToDay   interface{} `json:"sum_to_day"`
}

type TotalBalanceInMicroPay struct {
	Total interface{} `json:"total,omitempty"`

	Sum                    float64     `json:"sum,omitempty"`
	WalletTypeName         string      `json:"wallettypename,omitempty"`
	ResTotal               interface{} `json:"restotal,omitempty"`
	TotalBalanceInMicroPay interface{} `json:"total_balance_in_micro_pay"`
}

type StatementEndingBalanc struct {
	TotalStatementEndingBalance interface{} `json:"totalstatementendingbalance,omitempty"`
	TotalSum                    float64     `json:"totalsum,omitempty"`
	BankUID                     int         `json:"bank_uid,omitempty"`
	ResTotal                    interface{} `json:"restotal,omitempty"`
	Sum                         float64     `json:"sum,omitempty"`
	Total                       interface{} `json:"total,omitempty"`
	StatementEndingBalanc       interface{} `json:"statement_ending_balance,omitempty"`

	Uid int `json:"uid"`

	// StatementBalance float64   `json:"statement_balance,omitempty"`
	// StatementDate    time.Time `json:"statement_date,omitempty"`
	// Uid              int       `json:"uid,omitempty"`
}

type CreateStatementEndingBalanc struct {
	StatementBalance float64 `json:"statement_balance"`
	StatementDate    string  `json:"statement_date"`
	Uid              int     `json:"uid"`
	BankUID          int     `json:"bank_uid"`
	//StatementDateCreate string    `json:"statement_date_create"`
}

type GetEditStatementEndingBalanc struct {
	StatementBalance float64   `json:"statement_balance"`
	StatementDate    time.Time `json:"statement_date"`
	Uid              int       `json:"uid"`
	BankUID          int       `json:"bank_uid"`
	//StatementDateCreate string    `json:"statement_date_create"`
	//res interface{} `json:"res"`
}

type EditStatementEndingBalanc struct {
	StatementBalance float64 `json:"statement_balance"`
	StatementDate    string  `json:"statement_date"`
	Uid              int     `json:"uid"`
	BankUID          int     `json:"bank_uid"`
}

type CreateBulkTransaction struct {
	Uid                   int     `json:"uid"`
	DateTime              string  `json:"date_time"`
	BulkCreditSameday     float64 `json:"bulk_credit_sameday"`
	BulkCreditSamedayFee  float64 `json:"bulk_credit_sameday_fee"`
	Transfertobankaccount int     `json:"transfertobankaccount"`
}

type EditBulkTransaction struct {
	Uid                   int     `json:"uid" query:"id"`
	DateTime              string  `json:"dateTime"`
	BulkCreditSameday     float64 `json:"bulkCreditSameday"`
	BulkCreditSamedayFee  float64 `json:"bulkCreditSamedayFee"`
	Transfertobankaccount int     `json:"transfertobankaccount"`
}

type StatementIncomingBalance struct {
	TotalStatementIncomingBalance interface{} `json:"totalstatementincomingbalance,omitempty"`
	ResTotal                      interface{} `json:"restotal,omitempty"`
	Sum                           float64     `json:"sum,omitempty"`
	BankUID                       int         `json:"bank_uid,omitempty"`
	StatementIncomingBalance      interface{} `json:"statement_incoming_balance"`

	Uid int `json:"uid"`
}

type WalletToWallet struct {
	ResTotal       interface{} `json:"restotal,omitempty"`
	CountRow       interface{} `json:"countrow,omitempty"`
	Total          float64     `json:"total,omitempty"`
	PaymentChannel string      `json:"paymentchannel,omitempty"`
	Sum            float64     `json:"sum,omitempty"`
	Count          int         `json:"count,omitempty"`

	SumWalletToWallet   float64 `json:"sum_wallet_to_wallet,omitempty"`
	CountWalletToWallet int     `json:"count_wallet_to_wallet,omitempty"`
}

type SettleMerchantOnline struct {
	ResTotal       interface{} `json:"restotal,omitempty"`
	CountRow       interface{} `json:"countrow,omitempty"`
	Total          float64     `json:"total,omitempty"`
	PaymentChannel string      `json:"paymentchannel,omitempty"`
	Sum            float64     `json:"sum,omitempty"`
	Count          int         `json:"count,omitempty"`

	SumSettleMerchantOnline   float64 `json:"sum_settle_merchant_online,omitempty"`
	CountSettleMerchantOnline int     `json:"count_settle_merchant_online,omitempty"`
}

type CollectoralIn struct {
	ResTotal       interface{} `json:"restotal,omitempty"`
	CountRow       interface{} `json:"countrow,omitempty"`
	PaymentChannel string      `json:"paymentchannel,omitempty"`
	PaymentType    string      `json:"paymenttype,omitempty"`
	Sum            float64     `json:"sum"`
	Count          int         `json:"count"`

	SumCollectoralIn   float64 `json:"sum_collectoral_in"`
	CountCollectoralIn int     `json:"count_collectoral_in"`
}

type CollectoralOut struct {
	ResTotal    interface{} `json:"restotal"`
	CountRow    interface{} `json:"countrow"`
	Count       int         `json:"count"`
	PaymentType string      `json:"paymenttype"`
	Sum         float64     `json:"sum"`

	SumCollectoralOut   float64 `json:"sum_collectoral_out"`
	CountCollectoralOut int     `json:"count_collectoral_out"`
}

type CollectoralBalance struct {
	//CountRow int         `json:"countrow,omitempty"`
	Balance  interface{} `json:"balance,omitempty"`
	ResTotal interface{} `json:"restotal,omitempty"`
}

type AdjustToWallet struct {
	ResTotal       interface{} `json:"restotal,omitempty"`
	CountRow       interface{} `json:"countrow,omitempty"`
	PaymentChannel string      `json:"paymentchannel,omitempty"`
	Sum            float64     `json:"sum,omitempty"`
	Count          int         `json:"count,omitempty"`
	PaymentType    string      `json:"paymenttype,omitempty"`

	SumAdjustToWallet   float64 `json:"sum_adjust_to_wallet"`
	CountAdjustToWallet int     `json:"count_adjust_to_wallet"`
}

type PromptPayInOtherBank struct {
	ResTotal    interface{} `json:"res_total,omitempty"`
	CountRow    interface{} `json:"count_row,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumPromptPayInOtherBankTranIn   float64 `json:"sum_prompt_pay_in_other_bank_tran_in"`
	CountPromptPayInOtherBankTranIn int     `json:"count_prompt_pay_in_other_bank_tran_in"`
}

type PromptPayInTCRB struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	CountRow interface{} `json:"count_row,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	// Total          float64 `json:"total,omitempty"`
	PaymentChannel string `json:"paymentchannel,omitempty"`
	Count          int    `json:"count,omitempty"`

	SumPromptPayInTCRBTranIn   float64 `json:"sum_prompt_pay_in_tcrb_tran_in"`
	CountPromptPayInTCRBTranIn int     `json:"count_prompt_pay_in_tcrb_tran_in"`
}

type TopupLoanDisbursement struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	CountRow interface{} `json:"count_row,omitempty"`

	PaymentType                      string  `json:"paymenttype,omitempty"`
	Sum                              float64 `json:"sum,omitempty"`
	Count                            int     `json:"count,omitempty"`
	SumTopupLoanDisbursementTranIn   float64 `json:"sum_topup_loan_disbursement_tran_in"`
	CountTopupLoanDisbursementTranIn int     `json:"count_topup_loan_disbursement_tran_in"`
}

type TopupDirectDebit struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	CountRow interface{} `json:"count_row,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountTopupDirectDebitTranIn int     `json:"count_topup_direct_debit_tran_in"`
	SumTopupDirectDebitTranIn   float64 `json:"sum_topup_direct_debit_tran_in"`
}

type TopupPayRoll struct {
	ResTotal    interface{} `json:"res_total,omitempty"`
	CountRow    interface{} `json:"count_row,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumTopupPayRollTranIn   float64 `json:"sum_topup_pay_roll_tran_in"`
	CountTopupPayRollTranIn int     `json:"count_topup_pay_roll_tran_in"`
}

type OnlineLoanTopup struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	CountRow interface{} `json:"count_row,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountOnlineLoanTopupTranIn int     `json:"count_online_loan_topup_tran_in"`
	SumOnlineLoanTopupTranIn   float64 `json:"sum_online_loan_topup_tran_in"`
}

type CashBack struct {
	ResTotal interface{} `json:"restotal"`
	CountRow interface{} `json:"count_row"`

	PaymentType string  `json:"paymenttype"`
	Sum         float64 `json:"sum"`
	Count       int     `json:"count"`

	CountCashBackTranIn int     `json:"count_cash_back_tran_in"`
	SumCashBackTranIn   float64 `json:"sum_cash_back_tran_in"`
}

type AdjustFromWallet struct {
	ResTotal       interface{} `json:"restotal,omitempty"`
	CountRow       interface{} `json:"countrow,omitempty"`
	PaymentChannel string      `json:"paymentchannel,omitempty"`
	Sum            float64     `json:"sum,omitempty"`
	Count          int         `json:"count,omitempty"`

	SumAdjustFromWalletTranOut   float64 `json:"sum_adjust_from_wallet_tran_out"`
	CountAdjustFromWalletTranOut int     `json:"count_adjust_from_wallet_tran_out"`
}

type PromptPayOutOtherBank struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumPromptPayOutOtherBankTranOut   float64 `json:"sum_prompt_pay_out_other_bank_tran_out"`
	CountPromptPayOutOtherBankTranOut int     `json:"count_prompt_pay_out_other_bank_tran_out"`
}

type PromptPayOutTCRB struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumPromptPayOutTCRBTranOut   float64 `json:"sum_prompt_pay_out_tcrb_tran_out"`
	CountPromptPayOutTCRBTranOut int     `json:"count_prompt_pay_out_tcrb_tran_out"`
}

type TcrbBillPayment struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumTcrbBillPaymentTranOut   float64 `json:"sum_tcrb_bill_payment_tran_out"`
	CountTcrbBillPaymentTranOut int     `json:"count_tcrb_bill_payment_tran_out"`
}

type TransferToBankAccountTxn struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumTransferToBankAccountTxnTranOut   float64 `json:"sum_transfer_to_bank_account_txn_tran_out"`
	CountTransferToBankAccountTxnTranOut int     `json:"count_transfer_to_bank_account_txn_tran_out"`
}

type TransferToBankAccountFee struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumTransferToBankAccountFeeTranOut   float64 `json:"sum_transfer_to_bank_account_fee_tran_out"`
	CountTransferToBankAccountFeeTranOut int     `json:"count_transfer_to_bank_account_fee_tran_out"`
}

type OnlineBillPayment struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumOnlineBillPaymentTranOut   float64 `json:"sum_online_bill_payment_tran_out"`
	CountOnlineBillPaymentTranOut int     `json:"count_online_bill_payment_tran_out"`
}

type TCRBLoanTranOut struct {
	ResTotal    interface{} `json:"restotal"`
	CountRow    interface{} `json:"countrow"`
	PaymentType string      `json:"paymenttype"`
	Sum         float64     `json:"sum"`
	Count       int         `json:"count"`

	SumTCRBLoanTranOut   float64 `json:"sum_tcrb_loan_tran_out"`
	CountTCRBLoanTranOut int     `json:"count_tcrb_loan_tran_out"`
}

type ThaiHealthTranOut struct {
	ResTotal    interface{} `json:"restotal"`
	CountRow    interface{} `json:"countrow"`
	PaymentType string      `json:"paymenttype"`
	Sum         float64     `json:"sum"`
	Count       int         `json:"count"`

	SumThaiHealthTranOut   float64 `json:"sum_thai_health_tran_out"`
	CountThaiHealthTranOut int     `json:"count_thai_health_tran_out"`
}

type ThaiPaiboonTranOut struct {
	ResTotal    interface{} `json:"restotal"`
	CountRow    interface{} `json:"countrow"`
	PaymentType string      `json:"paymenttype"`
	Sum         float64     `json:"sum"`
	Count       int         `json:"count"`

	SumThaiPaiboonTranOut   float64 `json:"sum_thai_paiboon_tran_out"`
	CountThaiPaiboonTranOut int     `json:"count_thai_paiboon_tran_out"`
}

type PromptpayInTagThirty struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumPromptpayInTagThirtyTranIn   float64 `json:"sum_promptpay_in_tag_thirty_tran_in"`
	CountPromptpayInTagThirtyTranIn int     `json:"count_promptpay_in_tag_thirty_tran_in"`
}

type PromptpayOutTagThirty struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumPromptpayOutTagThirtyTranOut   float64 `json:"sum_promptpay_out_tag_thirty_tran_out"`
	CountPromptpayOutTagThirtyTranOut int     `json:"count_promptpay_out_tag_thirty_tran_out"`
}

type PromptpayInOtherBankTwentyThreeCutOff struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumPromptpayInOtherBankSuspend   float64 `json:"sum_promptpay_in_other_bank_suspend"`
	CountPromptpayInOtherBankSuspend int     `json:"count_promptpay_in_other_bank_suspend"`
}

type PromptpayOutOtherBankTwentyThreeCutOff struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	CountPromptpayOutOtherBankSuspend int     `json:"count_promptpay_out_other_bank_suspend"`
	SumPromptpayOutOtherBankSuspend   float64 `json:"sum_promptpay_out_other_bank_suspend"`
}

type PromptpayOutTCRBTwentyThreeCutOff struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	CountPromptpayOutTCRBSuspend int     `json:"count_promptpay_out_tcrb_suspend"`
	SumPromptpayOutTCRBSuspend   float64 `json:"sum_promptpay_out_tcrb_suspend"`
}

type TCRBBillPaymentTwentyOneCutOff struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	CountTCRBBillPaymentSuspend int     `json:"count_tcrb_bill_payment_suspend"`
	SumTCRBBillPaymentSuspend   float64 `json:"sum_tcrb_bill_payment_suspend"`
}

type TCRBBillPaymentSettleCollectoralTwentyOneCutOff struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
}

type OnlineBillPaymentSettleCollectoralTwentyThreeCutOff struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
}

type PromptPayOutSettleCollectoralTwentyThreeCutOff struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
}

type PromptPayOutTagThirySettleCollectoralTwentyThreeCutOff struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type TcrbBillPaymentSettlementTransaction struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumTcrbBillPaymentSettlementTransaction   float64 `json:"sum_tcrb_bill_payment_settlement_transaction"`
	CountTcrbBillPaymentSettlementTransaction int     `json:"count_tcrb_bill_payment_settlement_transaction"`
}

type OnlineBillPaymentSettlementTransaction struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumOnlineBillPaymentSettlementTransaction   float64 `json:"sum_online_bill_payment_settlement_transaction"`
	CountOnlineBillPaymentSettlementTransaction int     `json:"count_online_bill_payment_settlement_transaction"`
}

type PromptPayOutSettlementTransaction struct {
	//ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow    interface{} `json:"countrow,omitempty"`

	ResTotal float64 `json:"restotal,omitempty"`
	CountRow int     `json:"countrow,omitempty"`
	// ResTotal2 float64 `json:"restotal2,omitempty"`
	// CountRow2 int     `json:"countrow2,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Count       int     `json:"count,omitempty"`
	Sum         float64 `json:"sum,omitempty"`

	Result float64 `json:"result,omitempty"`

	SumPromptPayOutSettlementTransaction   float64 `json:"sum_prompt_pay_out_settlement_transaction"`
	CountPromptPayOutSettlementTransaction int     `json:"count_prompt_pay_out_settlement_transaction"`

	//SumPromptPayOutTCRBTranOut float64 `json:"sum_prompt_pay_out_tcrb_tran_out"`
}

type BulkCreditSameday struct {
	//ResTotal float64 `json:"restotal,omitempty"`
	ResTotal      interface{} `json:"restotal,omitempty"`
	ResTotal1     interface{} `json:"restotal1,omitempty"`
	CountRow      int         `json:"countrow,omitempty"`
	Result        float64     `json:"result,omitempty"`
	Creditsameday float64     `json:"creditsameday,omitempty"`
	Tranferbank   int         `json:"tranferbank,omitempty"`

	// BulkCreditSameday     interface{} `json:"bulk_credit_sameday"`
	// Transfertobankaccount interface{} `json:"transfertobankaccount"`

	BulkCreditSameday     interface{} `json:"bulk_credit_sameday"`
	Transfertobankaccount interface{} `json:"transfer_to_bank_account"`
}

type BulkCreditSamedayFee struct {
	BulkCreditSamedayFee     interface{} `json:"bulk_credit_sameday_fee"`
	TransfertobankaccountFee interface{} `json:"transfer_to_bank_account_fee"`
}

type GetTopupLoanDisbursementSettlementTransaction struct {
	// ResTotal interface{} `json:"restotal,omitempty"`
	// CountRow interface{} `json:"count_row,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	ResTotal float64 `json:"restotal,omitempty"`
	CountRow int     `json:"countrow,omitempty"`

	CountTopupLoanDisbursementSettlementTransaction int     `json:"count_topup_loan_disbursement_settlement_transaction"`
	SumTopupLoanDisbursementSettlementTransaction   float64 `json:"sum_topup_loan_disbursement_settlement_transaction"`
}

type OnlineLoanTopupSettlementTransaction struct {
	// ResTotal interface{} `json:"restotal,omitempty"`
	// CountRow interface{} `json:"count_row,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountOnlineLoanTopupSettlementTransaction int     `json:"count_online_loan_topup_settlement_transaction"`
	SumOnlineLoanTopupSettlementTransaction   float64 `json:"sum_online_loan_topup_settlement_transaction"`
}

type PromptPayInOtherBankSettlementTransaction struct {
	// ResTotal    interface{} `json:"res_total,omitempty"`
	// CountRow    interface{} `json:"count_row,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountPromptPayInOtherBankSettlementTransaction int     `json:"count_prompt_pay_in_other_bank_settlement_transaction"`
	SumPromptPayInOtherBankSettlementTransaction   float64 `json:"sum_prompt_pay_in_other_bank_settlement_transaction"`
}

type PromptPayInTagThirtySettlementTransaction struct {
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountPromptPayInTagThirtySettlementTransaction int     `json:"count_prompt_pay_in_tag_thirty_settlement_transaction"`
	SumPromptPayInTagThirtySettlementTransaction   float64 `json:"sum_prompt_pay_in_tag_thirty_settlement_transaction"`
}

type PromptPayOutTagThirtySettlementTransaction struct {
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountPromptPayOutTagThirtySettlementTransaction int     `json:"count_prompt_pay_out_tag_thirty_settlement_transaction"`
	SumPromptPayOutTagThirtySettlementTransaction   float64 `json:"sum_prompt_pay_out_tag_thirty_settlement_transaction"`
}

type RtpTcrbLoanTwentyOneSettlementTransaction struct {
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountRtpTcrbLoanTwentyOneSettlementTransaction int     `json:"count_rtp_tcrb_loan_twenty_one_settlement_transaction"`
	SumRtpTcrbLoanTwentyOneSettlementTransaction   float64 `json:"sum_rtp_tcrb_loan_twenty_one_settlement_transaction"`
}

type TransferToBankAccountIncome struct {
	SumTransferToBankAccountIncome   interface{} `json:"sum_transfer_to_bank_account_income"`
	CountTransferToBankAccountIncome interface{} `json:"count_transfer_to_bank_account_income"`
}

type OnlineLoanTopupIncome struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"count_row,omitempty"`

	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountOnlineLoanTopupIncome int     `json:"count_online_loan_topup_income"`
	SumOnlineLoanTopupIncome   float64 `json:"sum_online_loan_topup_income"`
}

type OnlineBillPaymentIncome struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Count       int     `json:"count,omitempty"`

	CountOnlineBillPaymentIncome int     `json:"count_online_bill_payment_income"`
	SumOnlineBillPaymentIncome   float64 `json:"sum_online_bill_payment_income"`
}

type NewBindingRevolvingIncome struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	Sum      float64     `json:"sum,omitempty"`
	Count    int         `json:"count,omitempty"`
	Status   string      `json:"status,omitempty"`

	CountNewBindingRevolvingIncome int     `json:"count_new_binding_revolving_income"`
	SumNewBindingRevolvingIncome   float64 `json:"sum_new_binding_revolving_income"`
}

type RtpTcrbLoanIncome struct {
	//ResTotal interface{} `json:"restotal,omitempty"`
	//Sum      float64     `json:"sum,omitempty"`
	Count       int    `json:"count,omitempty"`
	PaymentType string `json:"paymenttype,omitempty"`

	CountRtpTcrbLoanIncome int     `json:"count_rtp_tcrb_loan_income"`
	SumRtpTcrbLoanIncome   float64 `json:"sum_rtp_tcrb_loan_income"`
}

type RtpThaiPaiboonIncome struct {
	//ResTotal interface{} `json:"restotal,omitempty"`
	//Sum      float64     `json:"sum,omitempty"`
	Count     int    `json:"count,omitempty"`
	ToAccount string `json:"toaccount,omitempty"`

	CountRtpThaiPaiboonIncome int     `json:"count_rtp_thai_paiboon_income"`
	SumRtpThaiPaiboonIncome   float64 `json:"sum_rtp_thai_paiboon_income"`
}

type KycCompleteFeeExpenses struct {
	// Report         interface{} `json:"report,omitempty"`
	// CountNewWallet int         `json:"countnewwallet,omitempty"`

	ResTotal  interface{} `json:"restotal,omitempty"`
	Count     int         `json:"count"`
	KYCStatus string      `json:"kycstatus,omitempty"`
	Sum       float64     `json:"sum"`

	CountKycCompleteFeeExpenses int     `json:"count_kyc_complete_fee_expenses"`
	SumKycCompleteFeeExpenses   float64 `json:"sum_kyc_complete_fee_expenses"`
}

type PromptPayOutExpenses struct {
	//ResTotal    interface{} `json:"restotal,omitempty"`
	//CountRow    interface{} `json:"countrow,omitempty"`
	// PaymentType string  `json:"paymenttype,omitempty"`
	// Count       int     `json:"count,omitempty"`
	// Sum         float64 `json:"sum,omitempty"`

	ResTotal1 interface{} `json:"restotal1,omitempty"`
	ResTotal2 interface{} `json:"restotal2,omitempty"`

	PaymentType string `json:"paymenttype,omitempty"`
	Count       int    `json:"count,omitempty"`
	//Count1      int    `json:"count1,omitempty"`
	//Count2      int    `json:"count2,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	SumCountRow int     `json:"sumcountrow,omitempty"`

	CountPromptPayOutExpenses int     `json:"count_prompt_pay_out_expenses"`
	SumPromptPayOutExpenses   float64 `json:"sum_prompt_pay_out_expenses"`
}

type BulkTransferFeeExpense struct {
	//ResTotal float64 `json:"restotal,omitempty"`
	ResTotal      interface{} `json:"restotal,omitempty"`
	ResTotal1     interface{} `json:"restotal1,omitempty"`
	CountRow      int         `json:"countrow,omitempty"`
	Result        float64     `json:"result,omitempty"`
	Creditsameday float64     `json:"creditsameday,omitempty"`
	Tranferbank   int         `json:"tranferbank,omitempty"`

	CountTransfertobankaccount    interface{} `json:"count_transfertobankaccount"`
	SumCountTransfertobankaccount float64     `json:"sum_count_transfertobankaccount"`
}

type PromptpayOutTagThirtyExpense struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	CountPromptpayOutTagThirtyExpense int     `json:"count_promptpay_out_tag_thirty_expense"`
	SumPromptpayOutTagThirtyExpense   float64 `json:"sum_promptpay_out_tag_thirty_expense"`
}

type PromptpayOutICFeeExpense struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	ResTotal1   interface{} `json:"restotal1,omitempty"`
	ResTotal2   interface{} `json:"restotal2,omitempty"`

	SumPromptpayOutICFeeExpense   float64 `json:"sum_promptpay_out_ic_fee_expense"`
	CountPromptpayOutICFeeExpense int     `json:"count_promptpay_out_ic_fee_expense"`
}

type PromptpayInICFeeExpense struct {
	ResTotal    interface{} `json:"res_total,omitempty"`
	CountRow    interface{} `json:"count_row,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`
	ResTotal1   interface{} `json:"restotal1,omitempty"`
	ResTotal2   interface{} `json:"restotal2,omitempty"`

	SumPromptpayInICFeeExpense   float64 `json:"sum_promptpay_in_ic_fee_expense"`
	CountPromptpayInICFeeExpense int     `json:"count_promptpay_in_ic_fee_expense"`
}

type TmdsKycCaseExpense struct {
	//ResTotal1 interface{} `json:"restotal1,omitempty"`
	//ResTotal2 interface{} `json:"restotal2,omitempty"`
	Count interface{} `json:"count,omitempty"`
	Sum   interface{} `json:"sum,omitempty"`

	CountTmdsKycCaseExpense int `json:"count_tmds_kyc_case_expense"`
	SumTmdsKycCaseExpense   int `json:"sum_tmds_kyc_case_expense"`
}

type AllKycMonthly struct {
	KYCDate           interface{} `json:"kyc_date,omitempty"`
	KYCTime           interface{} `json:"kyc_time,omitempty"`
	AgentID           interface{} `json:"agent_id,omitempty"`
	Agentemail        interface{} `json:"agentemail,omitempty"`
	AgentNameLastname interface{} `json:"agent_name_lastname,omitempty"`
	KYCStatus         interface{} `json:"kyc_status,omitempty"`
	Consumerwalletid  interface{} `json:"consumerwalletid,omitempty"`
	KYCRespond        interface{} `json:"kyc_respond,omitempty"`
	DOPARespond       interface{} `json:"dopa_respond,omitempty"`
	AgentType         interface{} `json:"agent_type,omitempty"`
}

/* ///////////////////////////////////////////////////////////////////////////////////// */
type ReportWallet struct {
	Report   interface{} `json:"report,omitempty"`
	Count    int         `json:"count,omitempty"`
	Walletid string      `json:"walletid,omitempty"`
}

type PymentTypeSummary struct {
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
}

type PromptpayInTagThirtySuspendMerchant struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Sum         float64     `json:"sum,omitempty"`
	Count       int         `json:"count,omitempty"`

	SumPromptpayInTagThirtySuspendMerchant   float64 `json:"sum_promptpay_in_tag_thirty_suspend_merchant"`
	CountPromptpayInTagThirtySuspendMerchant int     `json:"count_promptpay_in_tag_thirty_suspend_merchant"`
}

type PromptpayOutTagThirtySuspendMerchant struct {
	Sum                                       float64 `json:"sum,omitempty"`
	Count                                     int     `json:"count,omitempty"`
	PaymentType                               string  `json:"paymenttype,omitempty"`
	SumPromptpayOutTagThirtySuspendMerchant   float64 `json:"sum_promptpay_out_tag_thirty_suspend_merchant"`
	CountPromptpayOutTagThirtySuspendMerchant int     `json:"count_promptpay_out_tag_thirty_suspend_merchant"`
}

type RtpTcrbLoanSuspendMerchant struct {
	Sum                             float64 `json:"sum,omitempty"`
	Count                           int     `json:"count,omitempty"`
	PaymentType                     string  `json:"paymenttype,omitempty"`
	SumRtpTcrbLoanSuspendMerchant   float64 `json:"sum_rtp_tcrb_loan_suspend_merchant"`
	CountRtpTcrbLoanSuspendMerchant int     `json:"count_rtp_tcrb_loan_suspend_merchant"`
}

type Reports struct {
	WalletDaily        []interface{}
	Consumer           []interface{}
	Merchant           []interface{}
	AgentKycCSV        []interface{}
	LoanbindingCSV     []interface{}
	Total              int
	FileImportID       int
	WalletID           string
	UserProfileAmloCSV []interface{}
	WatchlistImportCSV []interface{}
	Kycpendingcsv      []interface{}
	Lbpendingcsv       []interface{}
	Rvpendingcsv       []interface{}
	User               string
}

type WalletDaily struct {
	WalletID          string      `csv:"Wallet ID"`
	WalletTypeName    string      `csv:"Wallet Type Name"`
	WalletPhoneNumber string      `csv:"Wallet Phone no."`
	Email             string      `csv:"Email"`
	WalletName        string      `csv:"Wallet name"`
	CitizenID         string      `csv:"Citizen Id"`
	Status            string      `csv:"Status"`
	Balance           float64     `csv:"Balance"`
	IsForgetPin       string      `csv:"isForgetPin"`
	ATMCard           string      `csv:"ATM Card"`
	AccountNo         string      `csv:"Account No"`
	AddressDetail     string      `csv:"Address Detail"`
	Street            string      `csv:"Street"`
	District          string      `csv:"District"`
	SubDistrict       string      `csv:"Sub District"`
	Province          string      `csv:"Province"`
	PostalCode        string      `csv:"Postal Code"`
	Datetime          time.Time   `csv:"Datetime"`
	RegisterDateTime  time.Time   `csv:"Register DateTime"`
	FileimportID      interface{} `json:"fileimport_id"`
}

type Consumer struct {
	TransactionID       string      `csv:"TransactionID"`
	PartnerRef          string      `csv:"Partner Ref"`
	DateTime            time.Time   `csv:"DateTime"`
	TransactionStatus   string      `csv:"TransactionStatus"`
	TransactionType     string      `csv:"TransactionType"`
	PaymentChannel      string      `csv:"PaymentChannel"`
	PaymentType         string      `csv:"PaymentType"`
	TypeCode            string      `csv:"TypeCode"`
	ApprovalCode        string      `csv:"ApprovalCode"`
	BillerID            string      `csv:"BillerID"`
	Ref1                string      `csv:"Ref1"`
	Ref2                string      `csv:"Ref2"`
	Ref3                string      `csv:"Ref3"`
	Amount              float64     `csv:"Amount"`
	Fee                 float64     `csv:"Fee"`
	Total               float64     `csv:"Total"`
	FromReference       string      `csv:"FromReference"`
	FromPhoneNo         string      `csv:"FromPhoneNo"`
	FromName            string      `csv:"FromName"`
	ToAccount           string      `csv:"ToAccount"`
	ToAccountPhoneNo    string      `csv:"ToAccountPhoneNo"`
	ToAccountName       string      `csv:"ToAccountName"`
	BankCode            string      `csv:"BankCode"`
	TerminalId          string      `csv:"TerminalId"`
	TerminalType        string      `csv:"TerminalType"`
	ResponseCode        string      `csv:"ResponseCode"`
	ResponseDescription string      `csv:"ResponseDescription"`
	ToAccount105        string      `csv:"ToAccount105"`
	FromReference105    string      `csv:"FromReference105"`
	FileimportID        interface{} `json:"fileimport_id"`
}

type Kycpending struct {
	WalletID     string `csv:"Wallet_ID"`
	CustomerName string `csv:"Customer Name"`
	//KYCDate      time.Time   `csv:"KYCDate"`
	FileimportID interface{} `json:"fileimport_id"`
}

type Lbpending struct {
	WalletID     string `csv:"Customer Wallet_ID"`
	CustomerName string `csv:"Customer Name"`
	CAWalletID   string `csv:"CA Wallet ID"`
	CAPort       string `csv:"CAPort"`
	MainBranch   string `csv:"Main Branch"`
	//RVDate       time.Time   `csv:"RVDate"`
	FileimportID interface{} `json:"fileimport_id"`
}

type Merchant struct {
	TransactionID    string    `csv:"TransactionID"`
	DateTime         time.Time `csv:"DateTime"`
	MerchantID       string    `csv:"Merchant ID"`
	TerminalID       string    `csv:"Terminal ID"`
	MerchantFullName string    `csv:"Merchant FullName"`
	FromAccount      string    `csv:"From Account"`
	//FromName          string      `csv:"From Name"`
	SettlementAccount string      `csv:"Settlement Account"`
	Amount            float64     `csv:"Amount"`
	MDR_FEETHB        float64     `csv:"MDR_FEE(THB)"`
	TransactionFEETHB float64     `csv:"TransactionFEE(THB)"`
	TotalFeeincVAT    float64     `csv:"Total Fee (inc.VAT)"`
	VATTHB            float64     `csv:"VAT(THB)"`
	TotalFeeExcVAT    float64     `csv:"Total Fee (Exc.VAT)"`
	WHTTHB            float64     `csv:"WHT(THB)"`
	NetPayTHB         float64     `csv:"NetPay(THB)"`
	TransactionType   string      `csv:"Transaction Type"`
	PaymentType       string      `csv:"Payment Type"`
	PaymentChannel    string      `csv:"Payment Channel"`
	BankCode          string      `csv:"BankCode"`
	Status            string      `csv:"Status"`
	FileimportID      interface{} `json:"fileimport_id"`
}

type AgentKycCSV struct {
	KYCDate           string      `csv:"KYC Date"`
	KYCTime           string      `csv:"KYC Time"`
	AgentID           string      `csv:"Agent ID"`
	AgentEmail        string      `csv:"Agent email"`
	AgentNameLastname string      `csv:"Agent Name Lastname"`
	KYCStatus         string      `csv:"KYC Status"`
	ConsumerWalletID  string      `csv:"Consumer wallet id"`
	KYCRespond        string      `csv:"KYC Respond"`
	DOPARespond       string      `csv:"DOPA Respond"`
	FileimportID      interface{} `json:"fileimport_id"`
}

type LoanbindingCSV struct {
	WalletID         string      `csv:"Wallet Id"`
	AccountReference string      `csv:"Account Reference"`
	LoanAccountNo    string      `csv:"Loan Account No"`
	Status           string      `csv:"Status"`
	DateTime         time.Time   `csv:"DateTime"`
	RequestDateTime  time.Time   `csv:"Request DateTime"`
	FileimportID     interface{} `json:"fileimport_id"`
}

type MerchantFileimport struct {
	MerchantFileimportID     interface{} `json:"merchant_fileimport_id"`
	MerchantFileimportStatus interface{} `json:"merchant_fileimport_status"`
}

type AgentKycFileimport struct {
	AgentKycFileimportID     interface{} `json:"agent_kyc_fileimport_id"`
	AgentKycFileimportStatus interface{} `json:"agent_kyc_fileimport_status"`
}

type LoanbindingFileimport struct {
	LoanbindingFileimportID     interface{} `json:"loanbinding_fileimport_id"`
	LoanbindingFileimportStatus interface{} `json:"loanbinding_fileimport_status"`
}

type WalletdailyFileimport struct {
	WalletdailyFileImportID     interface{} `json:"walletdaily_file_import_id"`
	WalletdailyFileImportStatus interface{} `json:"walletdaily_file_import_status"`
}

type ConsumerFileimport struct {
	ConsumerFileimportID     interface{} `json:"consumer_fileimport_id"`
	ConsumerFileimportStatus interface{} `json:"consumer_fileimport_status"`
}

type KycPendingFileimport struct {
	KycPendingFileimportID     interface{} `json:"kyc_pending_fileimport_id"`
	KycPendingFileimportStatus interface{} `json:"kyc_pending_fileimport_status"`
}

type LbPendingFileimport struct {
	LbPendingFileimportID     interface{} `json:"lb_pending_fileimport_id"`
	LbPendingFileimportStatus interface{} `json:"lb_pending_fileimport_status"`
}

type OAuthPermission struct {
	Service  string `json:"service,omitempty"`
	Resource string `json:"resource,omitempty"`
	Action   string `json:"action,omitempty"`
	AdminID  string `json:"admin_id,omitempty"`
	Allow    bool   `json:"allow,omitempty"`
}

type CustomContext struct {
	echo.Context
	ClientID           string
	TokenResponse      *ResponseIntrospect
	PermissionResponse *OAuthPermission
	AdminID            string
}

func RespError(err error) Response {
	return Response{Message: err.Error()}
}

var UnmarshalTime = func(data []byte, t *time.Time) error {
	tt, err := time.Parse("2006/01/02 15:04:05", string(data))
	if err == nil {
		*t = tt
		return nil
	}

	tt, err = time.Parse("01/02/2006 15:04:05", string(data))
	if err == nil {
		*t = tt
		return nil
	}

	tt, err = time.Parse("2006-01-02 15:04:05", string(data))
	if err == nil {
		*t = tt
		return nil
	}

	tt, err = time.Parse("2006-01-02T15:04:05", string(data))
	if err == nil {
		*t = tt
		return nil
	}

	return errors.New("convert time error")
}

type BillPayMeaTranOut struct {
	ResTotal interface{} `json:"restotal"`
	//CountRow    interface{} `json:"countrow"`
	PaymentType string  `json:"paymenttype"`
	Sum         float64 `json:"sum"`
	Count       int     `json:"count"`

	SumBillPayMeaTranOut   float64 `json:"sum_bill_pay_mea_tran_out"`
	CountBillPayMeaTranOut int     `json:"count_bill_pay_mea_tran_out"`
}

type BillPayMeaSettleCollectoralTwentyThree struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type BillPayMeaTwentyThreeSettlementTransaction struct {
	PaymentType string  `json:"paymenttype"`
	Sum         float64 `json:"sum"`
	Count       int     `json:"count"`

	CountBillPayMeaTwentyThreeSettlementTransaction int     `json:"count_bill_pay_mea_twenty_three_settlement_transaction"`
	SumBillPayMeaTwentyThreeSettlementTransaction   float64 `json:"sum_bill_pay_mea_twenty_three_settlement_transaction"`
}

type CreateConfigPoint struct {
	ID                int    `json:"id"`
	TransactionName   string `json:"transactionname"`
	TransactionType   string `json:"transactiontype"`
	PaymentChannel    string `json:"paymentchannel"`
	PaymentType       string `json:"paymenttype"`
	DummyWallet       string `json:"dummywallet"`
	Amount            int    `json:"amount"`
	Point             int    `json:"point"`
	Expire            int    `json:"expire"`
	StatusTransaction string `json:"statustransaction"`
}

type EditConfigPoint struct {
	//ID              int    `json:"id" query:"id"`
	Uid               int    `json:"uid" query:"id"`
	TransactionName   string `json:"transactionname"`
	TransactionType   string `json:"transactiontype"`
	PaymentChannel    string `json:"paymentchannel"`
	PaymentType       string `json:"paymenttype"`
	DummyWallet       string `json:"dummywallet"`
	Amount            int    `json:"amount"`
	Point             int    `json:"point"`
	Expire            int    `json:"expire"`
	StatusTransaction string `json:"statustransaction"`
}

type LimitPoint struct {
	Uid        int `json:"uid" query:"id"`
	LimitPoint int `json:"limitpoint"`
}

type TransactionPoint struct {
	ID                    int       `json:"id"`
	Date                  time.Time `json:"date"`
	WalletID              string    `json:"wallet_id"`
	TransactionName       string    `json:"transaction_name"`
	Point                 int       `json:"point"`
	DateExpire            time.Time `json:"date_expire"`
	Reference             string    `json:"reference"`
	Count                 int       `json:"count"`
	CountTransactionPoint int       `json:"count_transaction_point"`
	Dates                 string    `json:"dates"` // new
}

type ConsumerPoint struct {
	TransactionName string      `json:"transactionname,omitempty"`
	TransactionType string      `json:"transactiontype,omitempty"`
	PaymentChannel  string      `json:"paymentchannel,omitempty"`
	PaymentType     string      `json:"paymenttype,omitempty"`
	Total           float64     `json:"total,omitempty"`
	ToAccount       string      `json:"toaccount,omitempty"`
	DateTime        time.Time   `json:"datetime,omitempty"`
	Sum             float64     `json:"sum,omitempty"`
	ResTotal        interface{} `json:"restotal,omitempty"`
	Point           float64     `json:"point,omitempty"`
	FromReference   string      `json:"fromreference,omitempty"`

	// TransactionType string      `json:"transactiontype,omitempty"`
	// PaymentChannel  string      `json:"paymentchannel,omitempty"`
	// PaymentType     string      `json:"paymenttype,omitempty"`
	// Total           float64     `json:"total,omitempty"`
	// ToAccount       string      `json:"toaccount,omitempty"`
	// DateTime        time.Time   `json:"datetime,omitempty"`
	// Sum             float64     `json:"sum,omitempty"`
	// ResTotal        interface{} `json:"restotal,omitempty"`
	// Point           float64     `json:"point,omitempty"`
	// FromReference   string      `json:"fromreference,omitempty"`
}

type ConsumerTranFactor struct {
	TransactionName string  `json:"transactionname,omitempty"`
	TransactionType string  `json:"transactiontype,omitempty"`
	PaymentChannel  string  `json:"paymentchannel,omitempty"`
	PaymentType     string  `json:"paymenttype,omitempty"`
	Total           float64 `json:"total,omitempty"`
	//ToAccount       string      `json:"toaccount,omitempty"`
	DateTime time.Time   `json:"datetime,omitempty"`
	Sum      float64     `json:"sum,omitempty"`
	ResTotal interface{} `json:"restotal,omitempty"`
	//Point           float64     `json:"point,omitempty"`
	FromReference string `json:"fromreference,omitempty"`
	Count         int    `json:"count,omitempty"`
}

// type PointCSV struct {
// 	//ID       int       `json:"id"`
// 	//WalletID string    `json:"wallet_id"`
// 	Note string `json:"Note"`
// 	// DateTime time.Time `json:"date_time"`
// 	// User     int       `json:"user"`
// }

type BillPayPeaTranOut struct {
	ResTotal interface{} `json:"restotal"`
	//CountRow    interface{} `json:"countrow"`
	PaymentType string  `json:"paymenttype"`
	Sum         float64 `json:"sum"`
	Count       int     `json:"count"`

	SumBillPayPeaTranOut   float64 `json:"sum_bill_pay_pea_tran_out"`
	CountBillPayPeaTranOut int     `json:"count_bill_pay_pea_tran_out"`
}

type BillPayPeaSettleCollectoralTwentyThree struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type BillPayPeaTwentyThreeSettlementTransaction struct {
	//Res interface{}
	PaymentType string  `json:"paymenttype"`
	Sum         float64 `json:"sum"`
	Count       int     `json:"count"`

	CountBillPayPeaTwentyThreeSettlementTransaction int     `json:"count_bill_pay_pea_twenty_three_settlement_transaction"`
	SumBillPayPeaTwentyThreeSettlementTransaction   float64 `json:"sum_bill_pay_pea_twenty_three_settlement_transaction"`
}

type PointExportCsv struct {
	ResTotal interface{} `json:"restotal,omitempty"`

	WalletID     string `json:"walletid,omitempty"`
	Adjustamount int    `json:"adjustamount,omitempty"`
	Sum          int    `json:"sum,omitempty"`
	Note         string `json:"note,omitempty"`

	SumAdjustamount int `json:"sum_adjustamount,omitempty"`
}

type CreateConfigOccupation struct {
	ID             int    `json:"id"`
	OccupationName string `json:"occupationname"`
	Rank           string `json:"rank"`
	//RankTmp          string `json:"ranktmp"`
	//StatusOccupation string `json:"statusoccupation"`
}

type ConfigArea struct {
	ID int `json:"id"`
	//OccupationName string `json:"occupationname"`
	Rank string `json:"rank"`
	//RankTmp          string `json:"ranktmp"`
	//StatusOccupation string `json:"statusoccupation"`
}

type ConfigDateCalculateRank struct {
	ID                      int    `json:"id"`
	Rank                    string `json:"rank"`
	NumDateCalculateRank    int    `json:"numdatecalculaterank"`
	NumDateCalculateRankTmp string `json:"numdatecalculateranktmp"`
}

type UserProfileAmloFileimport struct { // new
	UserProfileAmloFileimportID     interface{} `json:"user_profile_amlo_fileimport_id"`
	UserProfileAmloFileimportStatus interface{} `json:"user_profile_amlo_fileimport_status"`
}

type WatchlistFileimport struct { // new
	WatchlistFileimportID     interface{} `json:"watchlist_fileimport_id,omitempty"`
	WatchlistFileimportStatus interface{} `json:"watchlist_fileimport_status,omitempty"`
}

type UserProfileAmlo struct { // new
	UserId    string `csv:"userId"`
	Firstname string `csv:"firstname"`
	Lastname  string `csv:"lastname"`
	//FirstnameEN string `csv:"firstnameEN"`
	//FirstnameEN     string      `csv:"firstnameEN"`
	PhoneNo   string `csv:"phoneNo"`
	Email     string `csv:"email"`
	CitizenId string `csv:"citizenId"`
	//Passport        string      `csv:"passport"`
	//LaserCode            string      `csv:"laserCode"`
	//DisplayName                string      `csv:"displayName"`
	BirthDate string `csv:"birthDate"`
	Gender    string `csv:"gender"`
	//Address   string `csv:"address"`
	//Pincode                 float64     `csv:"pincode"`
	//Remake               float64     `csv:"remake"`
	//Status       string      `csv:"status"`
	//PreviousStatus         string      `csv:"previousStatus"`
	//CreateAt    time.Time `csv:"createAt"`
	//UpdateAt    time.Time `csv:"updateAt"`
	//CreateBy    string    `csv:"createBy"`
	//CreaterType string    `csv:"createrType"`
	//Latitude            string      `csv:"latitude"`
	//Longitude          string      `csv:"longitude"`
	//IsFirstTime        string      `csv:"isFirstTime"`
	//IsForgetPin        string      `csv:"isForgetPin"`
	//IsRegister 				string      `csv:"isRegister"`
	//IsLockOtp        string      `csv:"isLockOtp"`
	//LastLogin    string      `csv:"lastLogin"`
	//Password    string      `csv:"password"`
	//AcquirerId    string      `csv:"acquirerId"`
	//AddressId string `csv:"addressId"`
	//CitizenAddressId    string      `csv:"citizenAddressId"`
	//BusinessAddress string `csv:"businessAddress"`
	//MemberRefNo    string      `csv:"memberRefNo"`
	//TermVersion    string      `csv:"termVersion"`
	OccupationId int `csv:"occupationId"`
	//BadgeCount    string      `csv:"badgeCount"`
	//BranchBankId    string      `csv:"branchBankId"`
	//Issue_date    string      `csv:"issue_date"`
	//Expiry_date    string      `csv:"expiry_date"`
	//BirthPlace    string      `csv:"birthPlace"`
	//TokenRegister    string      `csv:"tokenRegister"`
	//ExpiryToken    string      `csv:"expiryToken"`
	//Message    string      `csv:"message"`
	//MessageOld    string      `csv:"messageOld"`
	FileimportID interface{} `json:"fileimport_id"`
}

type WathclistCSVImport struct { // new
	Name         string      `csv:"Name"`
	TaxID        string      `csv:"TaxID"`
	FileimportID interface{} `json:"fileimport_id"`
}

type TransactionFactor struct {
	ID                    int    `json:"id"`
	TransactionFactorName string `json:"transactionfactorname"`
	TransactionType       string `json:"transactiontype"`
	PaymentChannel        string `json:"paymentchannel"`
	PaymentType           string `json:"paymenttype"`
	Day                   int    `json:"day"`
	//Date                  string `json:"date"`
}

type TransactionFactorItem struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Rank int     `json:"rank"`
}

type TransactionFactorItemTmp struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Rank int     `json:"rank"`
}

type ManualCalculateOpts struct {
	Occupation        bool
	Area              bool
	Watchlist         bool
	TransactionFactor bool
}

type Area struct {
	ID             int    `json:"id"`
	Rank           string `json:"rank"`
	ProvinceNameTH string `json:"province_name_th"`
	DistrictNameTH string `json:"district_name_th"`
}

type Fileinsert struct {
	FileinsertStatus interface{} `json:"fileinsert_status"`

	// MerchantFileimportID     interface{} `json:"merchant_fileimport_id"`
	// MerchantFileimportStatus interface{} `json:"merchant_fileimport_status"`
}

type AisTopup struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumAisTopupTranOutTelco   float64 `json:"sum_ais_topup_tran_out_telco"`
	CountAisTopupTranOutTelco int     `json:"count_ais_topup_tran_out_telco"`
}

type AisPackage struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumAisPackageTranOutTelco   float64 `json:"sum_ais_package_tran_out_telco"`
	CountAisPackageTranOutTelco int     `json:"count_ais_package_tran_out_telco"`
}

type AisBillpay struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumAisBillpayTranOutTelco   float64 `json:"sum_ais_billpay_tran_out_telco"`
	CountAisBillpayTranOutTelco int     `json:"count_ais_billpay_tran_out_telco"`
}

type AisFiber struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumAisFiberTranOutTelco   float64 `json:"sum_ais_fiber_tran_out_telco"`
	CountAisFiberTranOutTelco int     `json:"count_ais_fiber_tran_out_telco"`
}

type TrueTopup struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumTrueTopupTranOutTelco   float64 `json:"sum_true_topup_tran_out_telco"`
	CountTrueTopupTranOutTelco int     `json:"count_true_topup_tran_out_telco"`
}

type TruePackage struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumTruePackageTranOutTelco   float64 `json:"sum_true_package_tran_out_telco"`
	CountTruePackageTranOutTelco int     `json:"count_true_package_tran_out_telco"`
}

type DtacTopup struct {
	ResTotal    interface{} `json:"restotal,omitempty"`
	CountRow    interface{} `json:"countrow,omitempty"`
	PaymentType string      `json:"paymenttype,omitempty"`
	Count       int         `json:"count,omitempty"`
	Sum         float64     `json:"sum,omitempty"`

	SumDtacTopupTranOutTelco   float64 `json:"sum_dtac_topup_tran_out_telco"`
	CountDtacTopupTranOutTelco int     `json:"count_dtac_topup_tran_out_telco"`
}

type AisTopupSettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type AisPackageSettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type AisBillpaySettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type AisFiberSettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type DtacTopupSettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type TrueTopupSettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type TruePackageSettleCollectoral struct {
	ResTotal interface{} `json:"restotal,omitempty"`
	//CountRow interface{} `json:"countrow,omitempty"`
	PaymentType string  `json:"paymenttype,omitempty"`
	Sum         float64 `json:"sum"`
}

type PointExportKYCCsv struct {
	ResTotal interface{} `json:"restotal,omitempty"`

	WalletID string `json:"walletid,omitempty"`
	Point    int    `json:"point,omitempty"`
	//Sum          int    `json:"sum,omitempty"`
	Note string `json:"note,omitempty"`
	//DateExport

	//SumAdjustamount int `json:"sum_adjustamount,omitempty"`
}
