package services

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/configoccupation"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/loanbinding"
	"go-api-report2/ent/merchanttransaction"
	"go-api-report2/ent/pendingkyc"
	"go-api-report2/ent/pendingloanbinding"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/reportwallet"
	"go-api-report2/ent/userwallet"
	"go-api-report2/ent/watchlist"
	"go-api-report2/model"
	"strings"
	"sync"
	"time"
)

const (
	DefaultWatchlistType = 1
	DefaultRankWatchlist = 3
)

func (reportService ReportService) InsertWalletDaily(slice []interface{}, fileImportID int) (int, error) {
	walletdaily := make([]*ent.ReportWalletCreate, len(slice))

	for i, d := range slice {
		s, ok := d.(model.WalletDaily)
		if !ok {
			return 0, errors.New("Convert interface to wallet daily error")
		}
		if len(s.CitizenID) > 13 {
			s.CitizenID = "INVALID"
		}

		walletdaily[i] = reportService.connection.ReportWallet.Create().
			SetWalletid(s.WalletID).
			SetWalletTypeName(s.WalletTypeName).
			SetWalletPhoneno(s.WalletPhoneNumber).
			SetWalletName(s.WalletName).
			SetCitizenId(s.CitizenID).
			SetStatus(s.Status).
			SetDateTime(s.Datetime).
			SetBalance(s.Balance).
			SetEmail(s.Email).
			SetIsForgetPin(s.IsForgetPin).
			SetATMCard(s.ATMCard).
			SetAccountNo(s.AccountNo).
			SetAddressDetail(s.AddressDetail).
			SetStreet(s.Street).
			SetDistrict(s.District).
			SetSubDistrict(s.SubDistrict).
			SetProvince(s.Province).
			SetPostalCode(s.PostalCode).
			SetRegisterDateTime(s.RegisterDateTime).
			SetFileimportID(fileImportID)
	}

	if _, err := reportService.connection.ReportWallet.CreateBulk(walletdaily...).Save(context.Background()); err != nil {
		return 0, err
	}

	return reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.FileimportID(fileImportID),
		).Count(context.Background())

}

func (reportService ReportService) InsertConsumer(slice []interface{}, fileImportID int) (int, error) {
	consume := make([]*ent.ConsumerCreate, len(slice))

	for i, d := range slice {

		s, ok := d.(model.Consumer)
		if !ok {
			return 0, errors.New("Convert interface to consumer transaction error")
		}
		consume[i] = reportService.connection.Consumer.Create().
			SetTransactionID(s.TransactionID).
			SetDateTime(s.DateTime).
			SetTransactionStatus(s.TransactionStatus).
			SetTransactionType(s.TransactionType).
			SetPaymentChannel(s.PaymentChannel).
			SetPaymentType(s.PaymentType).
			SetTypeCode(s.TypeCode).
			SetApprovalCode(s.ApprovalCode).
			SetBillerID(s.BillerID).
			SetRef1(s.Ref1).
			SetRef2(s.Ref2).
			SetRef3(s.Ref3).
			SetAmount(s.Amount).
			SetFee(s.Fee).
			SetTotal(s.Total).
			SetFromReference(s.FromReference).
			SetFromPhoneNo(s.FromPhoneNo).
			SetFromName(s.FromName).
			SetToAccount(s.ToAccount).
			SetToAccountPhoneNo(s.ToAccountPhoneNo).
			SetToAccountName(s.ToAccountName).
			SetBankCode(s.BankCode).
			SetTerminalId(s.TerminalId).
			SetTerminalType(s.TerminalType).
			SetToAccount105(s.ToAccount105).
			SetPartnerRef(s.PartnerRef).
			SetResponseCode(s.ResponseCode).
			SetResponseDescription(s.ResponseDescription).
			SetFromReference(s.FromReference).
			SetFileimportID(fileImportID)
	}
	if _, err := reportService.connection.Consumer.CreateBulk(consume...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.Consumer.Query().
		Where(
			consumer.FileimportID(fileImportID),
		).Count(context.Background())
}

func (reportService ReportService) InsertKycPending(slice []interface{}, fileImportID int) (int, error) {
	kycpen := make([]*ent.PendingkycCreate, len(slice))

	for i, d := range slice {
		s, ok := d.(model.Kycpending)
		if !ok {
			return 0, errors.New("Convert interface to kyc pending error")
		}

		kycpen[i] = reportService.connection.Pendingkyc.Create().
			SetWalletID(s.WalletID).
			SetName(s.CustomerName).
			SetFileimportID(fileImportID)
	}
	if _, err := reportService.connection.Pendingkyc.CreateBulk(kycpen...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.Pendingkyc.Query().
		Where(
			pendingkyc.FileimportID(fileImportID),
		).Count(context.Background())
}

func (reportService ReportService) InsertLbPending(slice []interface{}, fileImportID int) (int, error) {
	lbpen := make([]*ent.PendingloanbindingCreate, len(slice))

	for i, d := range slice {
		s, ok := d.(model.Lbpending)
		if !ok {
			return 0, errors.New("Convert interface to rv pending error")
		}
		lbpen[i] = reportService.connection.Pendingloanbinding.Create().
			SetWalletID(s.WalletID).
			SetNameLB(s.CustomerName).
			SetCAWalletID(s.CAWalletID).
			SetCAPort(s.CAPort).
			SetMainBranch(s.MainBranch).
			SetFileimportID(fileImportID)
	}
	if _, err := reportService.connection.Pendingloanbinding.CreateBulk(lbpen...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.Pendingloanbinding.Query().
		Where(
			pendingloanbinding.FileimportID(fileImportID),
		).Count(context.Background())
}

func (reportService ReportService) InsertLoanbinding(slice []interface{}, fileImportID int) (int, error) {
	loanbindings := make([]*ent.LoanbindingCreate, len(slice))

	for i, d := range slice {
		s, ok := d.(model.LoanbindingCSV)
		if !ok {
			return 0, errors.New("Convert interface to consumer transaction error")
		}
		loanbindings[i] = reportService.connection.Loanbinding.Create().
			SetWalletId(s.WalletID).
			SetAccountReference(s.AccountReference).
			SetLoanAccountNo(s.LoanAccountNo).
			SetStatus(s.Status).
			SetDateTime(s.DateTime).
			SetRequestDateTime(s.RequestDateTime).
			SetFileimportID(fileImportID)
	}
	if _, err := reportService.connection.Loanbinding.CreateBulk(loanbindings...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.Loanbinding.Query().
		Where(
			loanbinding.FileimportID(fileImportID),
		).Count(context.Background())
}

func (reportService ReportService) InsertAgentKYC(slice []interface{}, fileImportID int) (int, error) {
	agentkycs := make([]*ent.AgentkycCreate, len(slice))

	for i, d := range slice {
		s, ok := d.(model.AgentKycCSV)
		if !ok {
			return 0, errors.New("Convert interface to AgentKYC error")
		}
		agentkycs[i] = reportService.connection.Agentkyc.Create().
			SetKYCDate(s.KYCDate).
			SetKYCTime(s.KYCTime).
			SetAgentID(s.AgentID).
			SetAgentemail(s.AgentEmail).
			SetAgentNameLastname(s.AgentNameLastname).
			SetKYCStatus(s.KYCStatus).
			SetConsumerwalletid(s.ConsumerWalletID).
			SetKYCRespond(s.KYCRespond).
			SetDOPARespond(s.DOPARespond).
			SetFileimportID(fileImportID)
	}
	if _, err := reportService.connection.Agentkyc.CreateBulk(agentkycs...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.FileimportID(fileImportID),
		).Count(context.Background())
}

func (reportService ReportService) InsertMerchant(slice []interface{}, fileImportID int) (int, error) {
	merchan := make([]*ent.MerchantTransactionCreate, len(slice))

	for i, d := range slice {
		s, ok := d.(model.Merchant)
		if !ok {
			return 0, errors.New("Convert interface to merchant transaction error")
		}
		merchan[i] = reportService.connection.MerchantTransaction.Create().
			SetTransactionID(s.TransactionID).
			SetDateTime(s.DateTime).
			SetMerchantID(s.MerchantID).
			SetTerminalID(s.TerminalID).
			SetMerchantFullName(s.MerchantFullName).
			SetFromAccount(s.FromAccount).
			SetSettlementAccount(s.SettlementAccount).
			SetAmount(s.Amount).
			SetMDRFEETHB(s.MDR_FEETHB).
			SetTransactionFEETHB(s.TransactionFEETHB).
			SetTotalFeeincVAT(s.TotalFeeincVAT).
			SetVATTHB(s.VATTHB).
			SetTotalFeeExcVAT(s.TotalFeeExcVAT).
			SetWHTTHB(s.WHTTHB).
			SetNetPayTHB(s.NetPayTHB).
			SetTransactionType(s.TransactionType).
			SetPaymentType(s.PaymentType).
			SetPaymentChannel(s.PaymentChannel).
			SetBankCode(s.BankCode).
			SetStatus(s.Status).
			SetFileimportID(fileImportID)
	}
	if _, err := reportService.connection.MerchantTransaction.CreateBulk(merchan...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.MerchantTransaction.Query().
		Where(
			merchanttransaction.FileimportID(fileImportID),
		).Count(context.Background())
}

func (reportService ReportService) InsertUserProfileAmlo(wg *sync.WaitGroup, slice []interface{}, fileImportID int) (int, error) {

	confOccu, err := reportService.connection.Configoccupation.Query().
		Select(configoccupation.FieldOccupationName, configoccupation.FieldID).
		All(context.Background())
	if err != nil {
		return 0, err
	}
	ocMap := make(map[int]string, 0)

	for _, confO := range confOccu {
		ocMap[confO.ID] = confO.OccupationName // เก็บดาต้า
	}

	defer wg.Done()
	i := 0
	for _, d := range slice {
		s, ok := d.(model.UserProfileAmlo)
		if !ok {
			continue
		}
		_, err := reportService.connection.Userwallet.Update().
			Where(
				userwallet.CitizenId(s.CitizenId),
			).
			SetOccupationId(s.OccupationId).
			Save(context.Background())
		if err != nil {
			return 0, err
		}

		occuName, ok := ocMap[s.OccupationId]
		if ok {
			_, err := reportService.connection.Ranking.Update().
				Where(ranking.TaxID(s.CitizenId)).
				SetOccupationName(occuName).
				Save(context.Background())
			if err != nil {
				return 0, err
			}
		}

		i++
	}

	return i, nil
}

func (reportService ReportService) InsertWatchlist(slice []interface{}, fileImportID int, user string) (int, error) {

	watchlists := make([]*ent.WatchlistCreate, len(slice))

	for i, d := range slice {

		s, ok := d.(model.WathclistCSVImport)
		if !ok {
			return 0, errors.New("Convert interface to watchlist error")
		}

		taxid := strings.Replace(s.TaxID, "-", "", -1)
		wl, err := reportService.connection.Watchlist.Query().
			Where(watchlist.TaxID(taxid), watchlist.TaxIDNEQ("")).
			First(context.Background())

		var notFoundErr *ent.NotFoundError
		if err != nil && !errors.As(err, &notFoundErr) {
			return 0, err
		}

		if wl != nil {
			reportService.logger.Errorln("[Routine] Data dupicate", wl)
			return 00, err
		}

		watchlists[i] = reportService.connection.Watchlist.Create().
			SetName(s.Name).
			SetTaxID(taxid).
			SetRankWatchlist(DefaultRankWatchlist).
			SetFileimportID(fileImportID).
			SetImportDate(time.Now()).
			SetUserUpload(user).
			SetRelatedID(DefaultWatchlistType)
	}

	if _, err := reportService.connection.Watchlist.CreateBulk(watchlists...).Save(context.Background()); err != nil {
		return 0, err
	}
	return reportService.connection.Watchlist.Query().
		Where(
			watchlist.FileimportID(fileImportID),
		).Count(context.Background())
}
