package model

var (
	BadRequest = Response{
		Code:    40000,
		Message: "Input is invalid",
	}
)

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Count   int    `json:"count,omitempty"`
	// Data    interface{} `json:"data,omitempty"`
	Data interface{} `json:"data"`
	Url  interface{} `json:"url"`

	ResNewWallet                                           interface{} `json:"resNewWallet,omitempty"`
	ResNoOfWallet                                          interface{} `json:"resNoOfWallet,omitempty"`
	ResWalletBalance                                       interface{} `json:"reswalletbalance,omitempty"`
	ResDiffWalletBalance                                   interface{} `json:"resdiffwalletbalance,omitempty"`
	ResTotalBalanceInMicroPay                              interface{} `json:"restotalbalanceinmicropay,omitempty"`
	ResStatementEndingBalance                              interface{} `json:"resstatementendingbalance,omitempty"`
	ResStatementIncomingBalance                            interface{} `json:"resstatementincomingbalance,omitempty"`
	ResWalletToWallet                                      interface{} `json:"reswallettowallet,omitempty"`
	ResSettleMerchantOnline                                interface{} `json:"ressettlemerchantonline,omitempty"`
	ResCollectoralIn                                       interface{} `json:"rescollectoralin,omitempty"`
	ResCollectoralOut                                      interface{} `json:"rescollectoralout,omitempty"`
	ResCollectoralBalance                                  interface{} `json:"rescollectoralbalance,omitempty"`
	ResAdjustToWallet                                      interface{} `json:"resadjusttowallet,omitempty"`
	ResPromptPayInOtherBank                                interface{} `json:"respromptpayinotherbank,omitempty"`
	ResPromptPayInTCRB                                     interface{} `json:"respromptpayintcrb,omitempty"`
	ResTopupLoanDisbursement                               interface{} `json:"restopuploandisbursement,omitempty"`
	ResTopupDirectDebit                                    interface{} `json:"restopupdirectdebit,omitempty"`
	ResTopupPayRoll                                        interface{} `json:"restopuppayroll,omitempty"`
	ResOnlineLoanTopup                                     interface{} `json:"resonlineloantopup,omitempty"`
	ResAdjustFromWallet                                    interface{} `json:"resadjustfromwallet,omitempty"`
	ResPromptPayOutOtherBank                               interface{} `json:"respromptpayoutotherbank,omitempty"`
	ResPromptPayOutTCRB                                    interface{} `json:"respromptpayouttcrb,omitempty"`
	ResTcrbBillPayment                                     interface{} `json:"restcrbbillpayment,omitempty"`
	ResTransferToBankAccountTxn                            interface{} `json:"restransfertobankaccounttxn,omitempty"`
	ResTransferToBankAccountFee                            interface{} `json:"restransfertobankaccountfee,omitempty"`
	ResOnlineBillPayment                                   interface{} `json:"resonlinebillpayment,omitempty"`
	ResPromptpayInTagThirty                                interface{} `json:"respromptpayintagthirty,omitempty"`
	ResPromptpayOutTagThirty                               interface{} `json:"respromptpayouttagthirty,omitempty"`
	ResPromptpayInOtherBankTwentyThreeCutOff               interface{} `json:"respromptpayinotherbanktwentythreecutoff,omitempty"`
	ResPromptpayOutOtherBankTwentyThreeCutOff              interface{} `json:"respromptpayoutotherbanktwentythreecutoff,omitempty"`
	ResPromptpayOutTCRBTwentyThreeCutOff                   interface{} `json:"respromptpayouttcrbtwentythreecutoff,omitempty"`
	ResTCRBBillPaymentTwentyOneCutOff                      interface{} `json:"restcrbbillpaymenttwentyonecutoff,omitempty"`
	ResTCRBBillPaymentSettleCollectoralTwentyOneCutOff     interface{} `json:"restcrbbillpaymentsettlecollectoraltwentyonecutoff,omitempty"`
	ResOnlineBillPaymentSettleCollectoralTwentyThreeCutOff interface{} `json:"resonlinebillpaymentsettlecollectoraltwentythreecutoff,omitempty"`
	ResPromptPayOutSettleCollectoralTwentyThreeCutOff      interface{} `json:"respromptpayoutsettlecollectoraltwentythreecutoff,omitempty"`
}

type ResponseIntrospect struct {
	Active   bool   `json:"active"`
	Scope    string `json:"scope,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	Exp      int64  `json:"exp,omitempty"`
	Iat      int64  `json:"iat,omitempty"`
	Sub      string `json:"sub,omitempty"`
	// FirstName string `json:"first_name,omitempty"`
	// LastName  string `json:"last_name,omitempty"`
	// Email     string `json:"email,omitempty"`
	Ext map[string]interface{} `json:"ext,omitempty"`
}

type ResponsePermission struct {
	Message string          `json:"message,omitempty"`
	Data    OAuthPermission `json:"data,omitempty"`
}
