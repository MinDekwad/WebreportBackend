package daily

import (
	"go-api-report2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	timeCustomLayout = "2006-01-02"
	layoutISO        = "2006-01-02 00:00:00"
	layoutISOA       = "2006-01-02 15:04:05"
)

func (h Handler) GetReports(c echo.Context) error {
	slug := c.Param("slug")
	date := c.QueryParam("date")

	var err error
	var result interface{}

	switch slug {
	case "introduction":
		introMap := make(map[string]interface{}, 0)

		newWallet, err := h.dailyService.GetNewWallet(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report new wallet has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500014,
				Message: "New wallet has rrror",
			})
		}
		introMap["new_wallet"] = newWallet

		noOfWallet, err := h.dailyService.GetNoOfWallet(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report no of wallet has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500015,
				Message: "No of wallet has rrror",
			})
		}
		introMap["no_of_wallet"] = noOfWallet

		walletBalance, err := h.dailyService.GetWalletBalance(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report wallet balance has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500016,
				Message: "Wallet balance has rrror",
			})
		}
		introMap["wallet_balance"] = walletBalance

		diffWalletBalance, err := h.dailyService.GetDiffWalletBalance(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report diff wallet balance has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500017,
				Message: "Diff wallet balance has rrror",
			})
		}
		introMap["diff_wallet_balance"] = diffWalletBalance

		totalBalanceInMicroPay, err := h.dailyService.GetTotalBalanceInMicroPay(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report total balance in micro pay has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500018,
				Message: "Total balance in micro pay has rrror",
			})
		}
		introMap["total_balance_in_micro_pay"] = totalBalanceInMicroPay

		statementEndingBalance, err := h.dailyService.GetStatementEndingBalance(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report statement ending balance has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500019,
				Message: "Statement ending balance has rrror",
			})
		}
		introMap["statement_ending_balance"] = statementEndingBalance

		statementIncomingBalance, err := h.dailyService.GetStatementIncomingBalance(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report statement income balance has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500020,
				Message: "Statement income balance has rrror",
			})
		}
		introMap["statement_incoming_balance"] = statementIncomingBalance

		result = introMap

	case "collectoral":
		collectMap := make(map[string]interface{}, 0)

		collectIn, err := h.dailyService.GetCollectoralIn(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report collectoral in has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500021,
				Message: "Collectoral in has rrror",
			})
		}
		collectMap["collectoral_in"] = collectIn

		collectOut, err := h.dailyService.GetCollectoralOut(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report collectoral out has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500022,
				Message: "Collectoral out has rrror",
			})
		}
		collectMap["collectoral_out"] = collectOut

		collectBalance, err := h.dailyService.GetCollectoralBalance(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report collectoral balance has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500023,
				Message: "Collectoral balance has rrror",
			})
		}
		collectMap["collectoral_balance"] = collectBalance

		result = collectMap

	case "transactionIn":
		transactionInMap := make(map[string]interface{}, 0)

		walletToWallet, err := h.dailyService.GetWalletToWallet(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report wallet to wallet has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500024,
				Message: "Wallet to wallet has rrror",
			})
		}
		transactionInMap["wallet_to_wallet"] = walletToWallet

		settleMerchantOnline, err := h.dailyService.GetSettleMerchantOnline(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report settle merchant online has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500025,
				Message: "Settle merchant online has rrror",
			})
		}
		transactionInMap["settle_merchant_online"] = settleMerchantOnline

		adjustToWallet, err := h.dailyService.GetAdjustToWallet(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report adjust to wallet has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500026,
				Message: "Adjust to wallet has rrror",
			})
		}
		transactionInMap["adjust_to_wallet"] = adjustToWallet

		promptpayInOtherBank, err := h.dailyService.GetPromptPayInOtherBank(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in other bank has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500027,
				Message: "Promptpay in other bank has error",
			})
		}
		transactionInMap["promptpay_in_other_bank_tran_in"] = promptpayInOtherBank

		promptpayInTcrb, err := h.dailyService.GetPromptPayInTCRB(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in tcrb has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500028,
				Message: "Promptpay in tcrb has error",
			})
		}
		transactionInMap["promptpay_in_tcrb_tran_in"] = promptpayInTcrb

		promptpayInTagThirty, err := h.dailyService.GetPromptpayInTagThirty(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in tag thirty has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500029,
				Message: "Promptpay in tag thirty has error",
			})
		}
		transactionInMap["promptpay_in_tag_thirty_tran_in"] = promptpayInTagThirty

		topupLoanDisbursement, err := h.dailyService.GetTopupLoanDisbursement(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report topup loan disbursement has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500030,
				Message: "Topup loan disbursement has error",
			})
		}
		transactionInMap["topup_loan_disbursement_tran_in"] = topupLoanDisbursement

		topupDirectDebit, err := h.dailyService.GetTopupDirectDebit(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report topup direct debit has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500031,
				Message: "Topup direct debit has error",
			})
		}
		transactionInMap["topup_loan_direct_debit_tran_in"] = topupDirectDebit

		topupPayRoll, err := h.dailyService.GetTopupPayRoll(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report topup pay roll has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500032,
				Message: "Topup pay roll has error",
			})
		}
		transactionInMap["topup_pay_roll_tran_in"] = topupPayRoll

		onlineLoanTopup, err := h.dailyService.GetOnlineLoanTopup(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online loan topup has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500033,
				Message: "Online loan topup has error",
			})
		}
		transactionInMap["online_loan_topup_tran_in"] = onlineLoanTopup

		cashBack, err := h.dailyService.GetCashBack(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report cash back has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500034,
				Message: "Cash back has error",
			})
		}
		transactionInMap["cash_back_tran_in"] = cashBack

		result = transactionInMap

	case "transactionOut":
		transactionOutMap := make(map[string]interface{}, 0)

		adjustFromWallet, err := h.dailyService.GetAdjustFromWallet(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report adjust from wallet has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500035,
				Message: "Adjust from wallet has error",
			})
		}
		transactionOutMap["adjust_from_wallet"] = adjustFromWallet

		promptPayOutOtherBank, err := h.dailyService.GetPromptPayOutOtherBank(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out other bank has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500036,
				Message: "Promptpay out other bank has error",
			})
		}
		transactionOutMap["prompt_pay_out_other_bank_tran_out"] = promptPayOutOtherBank

		promptPayOutTcrb, err := h.dailyService.GetPromptPayOutTCRB(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out tcrb has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500037,
				Message: "Promptpay out tcrb has error",
			})
		}
		transactionOutMap["prompt_pay_out_tcrb_tran_out"] = promptPayOutTcrb

		promptPayOutTagThirty, err := h.dailyService.GetPromptpayOutTagThirty(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out tag thirty has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500038,
				Message: "Promptpay out tag thirty has error",
			})
		}
		transactionOutMap["prompt_pay_out_tag_thirty_tran_out"] = promptPayOutTagThirty

		tcrbBillPayment, err := h.dailyService.GetTcrbBillPayment(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report trcb bill payment has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500039,
				Message: "Tcrb bill payment has error",
			})
		}
		transactionOutMap["tcrb_bill_payment_tran_out"] = tcrbBillPayment

		transferToBankAccountTxn, err := h.dailyService.GetTransferToBankAccountTxn(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report tranfer to bank account txc has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500040,
				Message: "Tranfer to bank account txc has error",
			})
		}
		transactionOutMap["transfer_to_bank_account_txn_tran_out"] = transferToBankAccountTxn

		transferToBankAccountFee, err := h.dailyService.GetTransferToBankAccountFee(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report tranfer to bank account fee has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500041,
				Message: "Tranfer to bank account fee has error",
			})
		}
		transactionOutMap["transfer_to_bank_account_fee_tran_out"] = transferToBankAccountFee

		onlineBillPayment, err := h.dailyService.GetOnlineBillPayment(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online bill payment has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500042,
				Message: "Online bill payment has error",
			})
		}
		transactionOutMap["online_bill_payment_tran_out"] = onlineBillPayment

		rtpTCRBLoanTranOut, err := h.dailyService.GetRtpTCRBLoanTranOut(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp tcrb loan transaction out has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500043,
				Message: "Rtp tcrb loan transaction out has error",
			})
		}
		transactionOutMap["rtp_tcrb_loan_tran_out"] = rtpTCRBLoanTranOut

		rtpThaiHealthTranOut, err := h.dailyService.GetRtpThaiHealthTranOut(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp thai health transaction out has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500044,
				Message: "Rtp thai health transaction out has error",
			})
		}
		transactionOutMap["rtp_thai_health_tran_out"] = rtpThaiHealthTranOut

		rtpThaiPaiboonTranOut, err := h.dailyService.GetRtpThaiPaiboonTranOut(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp thai paiboon transaction out has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500045,
				Message: "Rtp thai paiboon transaction out has error",
			})
		}
		transactionOutMap["rtp_thai_paiboon_tran_out"] = rtpThaiPaiboonTranOut

		billPayMeaTranOut, err := h.dailyService.GetBillPayMeaTranOut(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bill pay mea transation out has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500109,
				Message: "Bill pay mea transation out has error",
			})
		}
		transactionOutMap["bill_pay_mea_tran_out"] = billPayMeaTranOut

		billPayPeaTranOut, err := h.dailyService.GetBillPayPeaTranOut(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bill pay pea transation out has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500125,
				Message: " Bill pay pea transation out has error ",
			})
		}
		transactionOutMap["bill_pay_pea_tran_out"] = billPayPeaTranOut

		result = transactionOutMap

	case "suspendAfterCutOffTimeMerchant":
		suspendAfterCutOffTimeMerchantMap := make(map[string]interface{}, 0)

		promptpayInTagThirtySuspendMerchant, err := h.dailyService.GetPromptpayInTagThirtySuspendMerchant(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in tag thirty suspend merchant has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500046,
				Message: "Promptpay in tag thirty suspend merchant has error",
			})
		}
		suspendAfterCutOffTimeMerchantMap["promptpay_in_tag_thirty_suspend_merchant"] = promptpayInTagThirtySuspendMerchant

		promptpayOutTagThirtySuspendMerchant, err := h.dailyService.GetPromptpayOutTagThirtySuspendMerchant(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out tag thirty suspend merchant has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500047,
				Message: "Promptpay out tag thirty suspend merchant has error",
			})
		}
		suspendAfterCutOffTimeMerchantMap["promptpay_out_tag_thirty_suspend_merchant"] = promptpayOutTagThirtySuspendMerchant

		rtpTcrbLoanSuspendMerchant, err := h.dailyService.GetRtpTcrbLoanSuspendMerchant(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp tcrb loan suspend merchant has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500048,
				Message: "Rtp tcrb loan suspend merchant has error",
			})
		}
		suspendAfterCutOffTimeMerchantMap["rtp_tcrb_loan_suspend_merchant"] = rtpTcrbLoanSuspendMerchant

		result = suspendAfterCutOffTimeMerchantMap

	// suspend after cut-off-time
	case "suspendAfterCutOffTime":
		suspendAfterCutOffTimeMap := make(map[string]interface{}, 0)

		promptpayInOtherBankTwentyThreeCutOff, err := h.dailyService.GetPromptpayInOtherBankTwentyThreeCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in other bank twenty three cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500049,
				Message: "Promptpay in other bank twenty three cut off has error",
			})
		}
		//suspendAfterCutOffTimeMap["online-bill-payment"] = promptpayInOtherBankTwentyThreeCutOff
		suspendAfterCutOffTimeMap["promptpay_in_other_bank_suspend"] = promptpayInOtherBankTwentyThreeCutOff

		promptpayOutOtherBankTwentyThreeCutOff, err := h.dailyService.GetPromptpayOutOtherBankTwentyThreeCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out other bank twenty three cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500050,
				Message: "Promptpay out other bank twenty three cut off has error",
			})
		}
		suspendAfterCutOffTimeMap["promptpay_out_other_bank_suspend"] = promptpayOutOtherBankTwentyThreeCutOff

		promptpayOutTCRBTwentyThreeCutOff, err := h.dailyService.GetPromptpayOutTCRBTwentyThreeCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out tcrb twenty three cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500051,
				Message: "Promptpay out tcrb twenty three cut off has error",
			})
		}
		suspendAfterCutOffTimeMap["promptpay_out_tcrb_suspend"] = promptpayOutTCRBTwentyThreeCutOff

		tcrbBillPaymentTwentyOneCutOff, err := h.dailyService.GetTCRBBillPaymentTwentyOneCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report tcrb bill payment twenty one cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500052,
				Message: "Tcrb bill payment twenty one cut off has error",
			})
		}
		suspendAfterCutOffTimeMap["tcrb_bill_payment_suspend"] = tcrbBillPaymentTwentyOneCutOff

		result = suspendAfterCutOffTimeMap

	// settlement-collectoral
	case "settlementCollectoral":
		settlementCollectoralMap := make(map[string]interface{}, 0)

		tcrbBillPaymentSettleCollectoralTwentyOneCutOff, err := h.dailyService.GetTCRBBillPaymentSettleCollectoralTwentyOneCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report tcrb bill payment settle collectoral twenty one cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500053,
				Message: "Tcrb bill payment settle collectoral twenty one cut off has error",
			})
		}
		settlementCollectoralMap["tcrb_bill_payment_settle_collectoral"] = tcrbBillPaymentSettleCollectoralTwentyOneCutOff

		onlineBillPaymentSettleCollectoralTwentyThreeCutOff, err := h.dailyService.GetOnlineBillPaymentSettleCollectoralTwentyThreeCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online bill payment settle collectoral twenty three cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500054,
				Message: "Online bill payment settle collectoral twenty three cut off has error",
			})
		}
		settlementCollectoralMap["online_bill_payment_settle_collectoral"] = onlineBillPaymentSettleCollectoralTwentyThreeCutOff

		promptPayOutSettleCollectoralTwentyThreeCutOff, err := h.dailyService.GetPromptPayOutSettleCollectoralTwentyThreeCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out settle collectoral twenty three cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500055,
				Message: "Promptpay out settle collectoral twenty three cut off has error",
			})
		}
		settlementCollectoralMap["prompt_pay_out_settle_collectoral"] = promptPayOutSettleCollectoralTwentyThreeCutOff

		promptPayOutTagThitySettleCollectoralTwentyThreeCutOff, err := h.dailyService.GetPromptPayOutTagThirtySettleCollectoralTwentyThreeCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out tag thirty settle collectoral twenty three cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500056,
				Message: "Promptpay out tag thirty settle collectoral twenty three cut off has error",
			})
		}
		settlementCollectoralMap["prompt_pay_out_tag_thirty_settle_collectoral"] = promptPayOutTagThitySettleCollectoralTwentyThreeCutOff

		rtpTcrbLoanSettleCollectoralTwentyOneCutOff, err := h.dailyService.GetRtpTcrbLoanSettleCollectoralTwentyOneCutOff(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp tcrb loan settle collectoral twenty one cut off has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500057,
				Message: "Rtp tcrb loan settle collectoral twenty one cut off has error",
			})
		}
		settlementCollectoralMap["rtp_tcrb_loan_settle_collectoral"] = rtpTcrbLoanSettleCollectoralTwentyOneCutOff

		billPayMeaSettleCollectoralTwentyThree, err := h.dailyService.GetBillPayMeaSettleCollectoralTwentyThree(date)
		if err != nil {
			h.logger.Error("[Handler] Get report bill pay mea settel collectoral twenty three has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500110,
				Message: "Bill pay mea settle collectoral twenty three has error",
			})
		}
		settlementCollectoralMap["bill_pay_mea_settle_collectoral"] = billPayMeaSettleCollectoralTwentyThree

		billPayPeaSettleCollectoralTwentyThree, err := h.dailyService.GetBillPayPeaSettleCollectoralTwentyThree(date)
		if err != nil {
			h.logger.Error("[Handler] Get report bill pay pea settel collectoral twenty three has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500126,
				Message: " Bill pay pea settle collectoral twenty three has error ",
			})
		}
		settlementCollectoralMap["bill_pay_pea_settle_collectoral"] = billPayPeaSettleCollectoralTwentyThree

		aisTopupSettleCollectoral, err := h.dailyService.GetAisTopupSettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report AIS Topup settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500190,
				Message: " AIS Topup settle collectoral has error ",
			})
		}
		settlementCollectoralMap["ais_topup_settle_collectoral"] = aisTopupSettleCollectoral

		aisPackageSettleCollectoral, err := h.dailyService.GetAisPackageSettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report AIS package settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500191,
				Message: " AIS package settle collectoral has error ",
			})
		}
		settlementCollectoralMap["ais_package_settle_collectoral"] = aisPackageSettleCollectoral

		aisBillpaySettleCollectoral, err := h.dailyService.GetAisBillpaySettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report AIS billpay settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500192,
				Message: " AIS billpay settle collectoral has error ",
			})
		}
		settlementCollectoralMap["ais_billpay_settle_collectoral"] = aisBillpaySettleCollectoral

		aisFiberSettleCollectoral, err := h.dailyService.GetAisFiberSettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report AIS fiber settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500193,
				Message: " AIS fiber collectoral has error ",
			})
		}
		settlementCollectoralMap["ais_fiber_settle_collectoral"] = aisFiberSettleCollectoral

		dtacTopupSettleCollectoral, err := h.dailyService.GetDtacTopupSettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report Dtac topup settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500194,
				Message: " Dtac topup collectoral has error ",
			})
		}
		settlementCollectoralMap["dtac_topup_settle_collectoral"] = dtacTopupSettleCollectoral

		trueTopupSettleCollectoral, err := h.dailyService.GetTrueTopupSettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report True topup settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500195,
				Message: " True topup collectoral has error ",
			})
		}
		settlementCollectoralMap["true_topup_settle_collectoral"] = trueTopupSettleCollectoral

		truePackageSettleCollectoral, err := h.dailyService.GetTruePackageSettleCollectoral(date)
		if err != nil {
			h.logger.Error("[Handler] Get report True package settel collectoral has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500196,
				Message: " True package collectoral has error ",
			})
		}
		settlementCollectoralMap["true_package_settle_collectoral"] = truePackageSettleCollectoral

		result = settlementCollectoralMap

	// settlement-transaction
	case "settlementTransaction":
		settlementTransactionMap := make(map[string]interface{}, 0)

		tcrbBillPaymentSettlementTransaction, err := h.dailyService.GetTcrbBillPaymentSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report tcrb bill payment settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500058,
				Message: "Tcrb bill payment settlement transaction has error",
			})
		}
		settlementTransactionMap["tcrb_bill_payment_settlement_transaction"] = tcrbBillPaymentSettlementTransaction

		onlineBillPaymentSettlementTransaction, err := h.dailyService.GetOnlineBillPaymentSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online bill payment settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500059,
				Message: "Online bill payment settlement transaction has error",
			})
		}
		settlementTransactionMap["online_bill_payment_settlement_transaction"] = onlineBillPaymentSettlementTransaction

		promptPayOutSettlementTransaction, err := h.dailyService.GetPromptPayOutSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500060,
				Message: "Promptpay out settlement transaction has error",
			})
		}
		settlementTransactionMap["prompt_pay_out_settlement_transaction"] = promptPayOutSettlementTransaction

		bulkCreditSameday, err := h.dailyService.GetBulkCreditSameday(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bulk credit sameday has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500061,
				Message: "Bulk credit sameday has error",
			})
		}
		settlementTransactionMap["bulk_credit_sameday_settlement_transaction"] = bulkCreditSameday

		bulkCreditSamedayFee, err := h.dailyService.GetBulkCreditSamedayFee(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bulk credit sameday fee has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500062,
				Message: "Bulk credit sameday fee has error",
			})
		}
		settlementTransactionMap["bulk_credit_sameday_fee_settlement_transaction"] = bulkCreditSamedayFee

		topupLoanDisbursementSettlementTransaction, err := h.dailyService.GetTopupLoanDisbursementSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report topup loan disbursement settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500063,
				Message: "Topup loan disbursement settlement transaction has error",
			})
		}
		settlementTransactionMap["topup_loan_disbursement_settlement_transaction"] = topupLoanDisbursementSettlementTransaction

		onlineLoanTopupSettlementTransaction, err := h.dailyService.GetOnlineLoanTopupSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online loan topup settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500064,
				Message: "Online loan topup settlement transaction has error",
			})
		}
		settlementTransactionMap["online_loan_topup_settlement_transaction"] = onlineLoanTopupSettlementTransaction

		promptPayInOtherBankSettlementTransaction, err := h.dailyService.GetPromptPayInOtherBankSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in other bank settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500065,
				Message: "Promptpay in other bank settlement transaction has error",
			})
		}
		settlementTransactionMap["prompt_pay_in_other_bank_settlement_transaction"] = promptPayInOtherBankSettlementTransaction

		promptPayInTagThirtySettlementTransaction, err := h.dailyService.GetPromptPayInTagThirtySettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in tag thirty settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500066,
				Message: "Promptpay in tag thirty settlement transaction has error",
			})
		}
		settlementTransactionMap["prompt_pay_in_tag_thirty_settlement_transaction"] = promptPayInTagThirtySettlementTransaction

		promptPayOutTagThirtySettlementTransaction, err := h.dailyService.GetPromptPayOutTagThirtySettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay tag thirty settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500067,
				Message: "Promptpay tag thirty settlement transaction has error",
			})
		}
		settlementTransactionMap["prompt_pay_out_tag_thirty_settlement_transaction"] = promptPayOutTagThirtySettlementTransaction

		rtpTcrbLoanTwentyOneSettlementTransaction, err := h.dailyService.GetRtpTcrbLoanTwentyOneSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp tcrb loan twenty one settlement transaction has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500068,
				Message: "Rtp tcrb loan twenty one settlement transaction has error",
			})
		}
		settlementTransactionMap["rtp_tcrb_loan_twenty_one_settlement_transaction"] = rtpTcrbLoanTwentyOneSettlementTransaction

		billPayMeaTwentyThreeSettlementTransaction, err := h.dailyService.GetBillPayMeaTwentyThreeSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bill pay mea twenty three settlement transacdtion has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500111,
				Message: "Bill pay mea twenty three settlement transaction has error",
			})
		}
		settlementTransactionMap["bill_pay_mea_twenty_three_settlement_transaction"] = billPayMeaTwentyThreeSettlementTransaction

		billPayPeaTwentyThreeSettlementTransaction, err := h.dailyService.GetBillPayPeaTwentyThreeSettlementTransaction(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bill pay pea twenty three settlement transacdtion has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500127,
				Message: " Bill pay pea twenty three settlement transaction has error ",
			})
		}
		settlementTransactionMap["bill_pay_pea_twenty_three_settlement_transaction"] = billPayPeaTwentyThreeSettlementTransaction

		result = settlementTransactionMap

	case "income":
		incomeMap := make(map[string]interface{}, 0)

		transferToBankAccountIncome, err := h.dailyService.GetTransferToBankAccountIncome(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report transfer to bank account income has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500069,
				Message: "Transfer to bank account income has error",
			})
		}
		incomeMap["transfer_to_bank_account_income"] = transferToBankAccountIncome

		newBindingRevolving, err := h.dailyService.GetNewBindingRevolvingIncome(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report new binding revolving income has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500070,
				Message: "New binding revolving income has error",
			})
		}
		incomeMap["new_binding_revolving"] = newBindingRevolving

		onlineLoanTopupIncome, err := h.dailyService.GetOnlineLoanTopupIncome(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online loan topup income has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500071,
				Message: "Online loan topup income has error",
			})
		}
		incomeMap["online_loan_topup_income"] = onlineLoanTopupIncome

		onlineBillPaymentIncome, err := h.dailyService.GetOnlineBillPaymentIncome(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report online bill payment income has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500072,
				Message: "Online bill payment income has error",
			})
		}
		incomeMap["online_bill_payment_income"] = onlineBillPaymentIncome

		rtpTcrbLoanIncome, err := h.dailyService.GetRtpTcrbLoanIncome(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp tcrb loan income has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500073,
				Message: "Rtp tcrb loan income has error",
			})
		}
		incomeMap["rtp_tcrb_loan_income"] = rtpTcrbLoanIncome

		rtpThaiPaiboonIncome, err := h.dailyService.GetRtpThaiPaiboonIncome(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report rtp thai paiboon income has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500074,
				Message: "Rtp thai paiboon income has error",
			})
		}
		incomeMap["rtp_thai_paiboon_income"] = rtpThaiPaiboonIncome

		result = incomeMap

	// expenses
	case "expenses":
		expensesMap := make(map[string]interface{}, 0)

		kycCompleteFeeExpenses, err := h.dailyService.GetKycCompleteFeeExpenses(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report kyc complete fee expenses has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500075,
				Message: "Kyc complete fee expenses has error",
			})
		}
		expensesMap["kyc_complete_fee_expenses"] = kycCompleteFeeExpenses

		promptPayOutExpenses, err := h.dailyService.GetPromptPayOutExpenses(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out expenses has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500076,
				Message: "Promptpay out expenses has error",
			})
		}
		expensesMap["prompt_pay_out_expenses"] = promptPayOutExpenses

		bulkTransferFeeExpense, err := h.dailyService.GetBulkTransferFeeExpense(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report bulk transfer fee expense has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500077,
				Message: "Bulk transfer fee expense has error",
			})
		}
		expensesMap["bulk_transfer_fee_expense"] = bulkTransferFeeExpense

		tmdsKycCaseExpense, err := h.dailyService.GetTmdsKycCaseExpense(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report tmds kyc case expense has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500078,
				Message: "Tmds kyc case expense has error",
			})
		}
		expensesMap["tmds_kyc_case_expense"] = tmdsKycCaseExpense

		promptpayOutIcfeeExpense, err := h.dailyService.GetPromptpayOutICFeeExpense(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out ic fee expense has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500079,
				Message: "Promptpay out ic fee expense has error",
			})
		}
		expensesMap["promptpay_out_icfee_expense"] = promptpayOutIcfeeExpense

		promptpayInIcfeeExpense, err := h.dailyService.GetPromptpayInICFeeExpense(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay in ic fee expense has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500080,
				Message: "Promptpay in ic fee expense has error",
			})
		}
		expensesMap["promptpay_in_icfee_expense"] = promptpayInIcfeeExpense

		promptpayOutTagThirtyExpense, err := h.dailyService.GetPromptpayOutTagThirtyExpense(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report promptpay out tag thirty expense has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500081,
				Message: "Promptpay out tag thirty expense has error",
			})
		}
		expensesMap["promptpay_out_tag_thirty_expense"] = promptpayOutTagThirtyExpense

		result = expensesMap

	case "tranOutTelco":
		tranOutTelcoMap := make(map[string]interface{}, 0)

		aisTopup, err := h.dailyService.GetAisTopup(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report ais topup has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500183,
				Message: "Get report ais topup has error",
			})
		}
		tranOutTelcoMap["ais_topup"] = aisTopup

		aisPackage, err := h.dailyService.GetAisPackage(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report ais package has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500184,
				Message: "Get report ais package has error",
			})
		}
		tranOutTelcoMap["ais_package"] = aisPackage

		aisBillpay, err := h.dailyService.GetAisBillpay(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report ais billpay has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500185,
				Message: "Get report ais billpay has error",
			})
		}
		tranOutTelcoMap["ais_billpay"] = aisBillpay

		aisFiber, err := h.dailyService.GetAisFiber(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report ais fiber has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500186,
				Message: "Get report ais fiber has error",
			})
		}
		tranOutTelcoMap["ais_fiber"] = aisFiber

		trueTopup, err := h.dailyService.GetTrueTopup(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report true topup has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500187,
				Message: "Get report true topup has error",
			})
		}
		tranOutTelcoMap["true_topup"] = trueTopup

		truePackage, err := h.dailyService.GetTruePackage(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report true package has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500188,
				Message: "Get report true package has error",
			})
		}
		tranOutTelcoMap["true_package"] = truePackage

		dtacTopup, err := h.dailyService.GetDtacTopup(date)
		if err != nil {
			h.logger.Errorf("[Handler] Get report dtac topup has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500189,
				Message: "Get report dtac topup  has error",
			})
		}
		tranOutTelcoMap["dtac_topup"] = dtacTopup

		result = tranOutTelcoMap

	default:
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Error", Data: "This report type is not supported"})
	}

	if err != nil {
		h.logger.Errorf("[Handler] Get report has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500102,
			Message: "Get Data report has error",
		})
	}
	return c.JSON(http.StatusOK, model.Response{Message: "Success", Data: result})
}
