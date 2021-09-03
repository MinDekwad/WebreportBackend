package services

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/consumer"
	"go-api-report2/ent/fileimport"
	"go-api-report2/ent/loanbinding"
	"go-api-report2/ent/merchanttransaction"
	"go-api-report2/ent/reportwallet"
	"go-api-report2/model"
)

func (reportService ReportService) GetDataMerchantFileimportStatus(filename string) (model.MerchantFileimport, error) {
	var result model.MerchantFileimport

	merchantFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(merchantFileimport) == 0 {
		result.MerchantFileimportStatus = ""
		result.MerchantFileimportID = 0
		return result, nil
	}
	result.MerchantFileimportStatus = merchantFileimport[0].Status
	result.MerchantFileimportID = merchantFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataMerchantCSV(Fileimportid int) (interface{}, error) {
	merchant, err := reportService.connection.MerchantTransaction.Query().
		Where(
			merchanttransaction.FileimportID(Fileimportid),
		).
		Order(ent.Desc(merchanttransaction.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return merchant, nil

}

func (reportService ReportService) GetDataAgentKycFileimportStatus(filename string) (model.AgentKycFileimport, error) {
	var result model.AgentKycFileimport

	agentkycFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(agentkycFileimport) == 0 {
		result.AgentKycFileimportStatus = ""
		result.AgentKycFileimportID = 0
		return result, nil
	}
	result.AgentKycFileimportStatus = agentkycFileimport[0].Status
	result.AgentKycFileimportID = agentkycFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataAgentKycCSV(Fileimportid int) (interface{}, error) {
	agentkyc, err := reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.FileimportID(Fileimportid),
		).
		Order(ent.Desc(agentkyc.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return agentkyc, nil
}

func (reportService ReportService) GetDataLoanbindingFileimportStatus(filename string) (model.LoanbindingFileimport, error) {
	var result model.LoanbindingFileimport

	loanbindingFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(loanbindingFileimport) == 0 {
		result.LoanbindingFileimportStatus = ""
		result.LoanbindingFileimportID = 0
		return result, nil
	}
	result.LoanbindingFileimportStatus = loanbindingFileimport[0].Status
	result.LoanbindingFileimportID = loanbindingFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataLoanbindingCSV(Fileimportid int) (interface{}, error) {
	loanbinding, err := reportService.connection.Loanbinding.Query().
		Where(
			loanbinding.FileimportID(Fileimportid),
		).
		Order(ent.Desc(loanbinding.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return loanbinding, nil
}

func (reportService ReportService) GetDataWalletdailyFileimportStatus(filename string) (model.WalletdailyFileimport, error) {
	var result model.WalletdailyFileimport

	walletDailyFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(walletDailyFileimport) == 0 {
		result.WalletdailyFileImportStatus = ""
		result.WalletdailyFileImportID = 0
		return result, nil
	}
	result.WalletdailyFileImportStatus = walletDailyFileimport[0].Status
	result.WalletdailyFileImportID = walletDailyFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataWalletdailyCSV(Fileimportid int) (interface{}, error) {
	walletdailys, err := reportService.connection.ReportWallet.Query().
		Where(
			reportwallet.FileimportID(Fileimportid),
		).
		Order(ent.Desc(reportwallet.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return walletdailys, nil
}

func (reportService ReportService) GetDataConsumerFileimportStatus(filename string) (model.ConsumerFileimport, error) {
	var result model.ConsumerFileimport

	consumerFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(consumerFileimport) == 0 {
		result.ConsumerFileimportStatus = ""
		result.ConsumerFileimportID = 0
		return result, nil
	}
	result.ConsumerFileimportStatus = consumerFileimport[0].Status
	result.ConsumerFileimportID = consumerFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataConsumerCSV(Fileimportid int) (interface{}, error) {
	consumers, err := reportService.connection.Consumer.Query().
		Where(
			consumer.FileimportID(Fileimportid),
		).
		Order(ent.Desc(consumer.FieldID)).
		All(context.Background())
	if err != nil {
		return "", err
	}
	return consumers, nil
}

// new
func (reportService ReportService) GetDataUserProfileAmloFileimportStatus(filename string) (model.UserProfileAmloFileimport, error) {
	var result model.UserProfileAmloFileimport

	userProfileAmloFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(userProfileAmloFileimport) == 0 {
		result.UserProfileAmloFileimportStatus = ""
		result.UserProfileAmloFileimportID = 0
		return result, nil
	}
	result.UserProfileAmloFileimportStatus = userProfileAmloFileimport[0].Status
	result.UserProfileAmloFileimportID = userProfileAmloFileimport[0].ID
	return result, nil
}

// new
func (reportService ReportService) GetDataWatchlistFileimportStatus(filename string) (model.WatchlistFileimport, error) {
	var result model.WatchlistFileimport

	userWatchlistFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(userWatchlistFileimport) == 0 {
		result.WatchlistFileimportStatus = ""
		result.WatchlistFileimportID = 0
		return result, nil
	}
	result.WatchlistFileimportStatus = userWatchlistFileimport[0].Status
	result.WatchlistFileimportID = userWatchlistFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataKycPendingFileimportStatus(filename string) (model.KycPendingFileimport, error) {
	var result model.KycPendingFileimport

	kycpendingFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(kycpendingFileimport) == 0 {
		result.KycPendingFileimportStatus = ""
		result.KycPendingFileimportID = 0
		return result, nil
	}
	result.KycPendingFileimportStatus = kycpendingFileimport[0].Status
	result.KycPendingFileimportID = kycpendingFileimport[0].ID
	return result, nil
}

func (reportService ReportService) GetDataLbPendingFileimportStatus(filename string) (model.LbPendingFileimport, error) {
	var result model.LbPendingFileimport

	lbpendingFileimport, err := reportService.connection.Fileimport.Query().
		Select(fileimport.FieldStatus, fileimport.FieldID).
		Where(
			fileimport.Filename(filename),
		).
		All(context.Background())
	if err != nil {
		return result, err
	}
	if len(lbpendingFileimport) == 0 {
		result.LbPendingFileimportStatus = ""
		result.LbPendingFileimportID = 0
		return result, nil
	}
	result.LbPendingFileimportStatus = lbpendingFileimport[0].Status
	result.LbPendingFileimportID = lbpendingFileimport[0].ID
	return result, nil
}
