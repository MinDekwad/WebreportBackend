package services

import (
	"context"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/reportwallet"
	"log"
	"time"
)

func (reportService ReportService) PostUpdateReportKYC(date string) (interface{}, error) {

	result, err := reportService.connection.Agentkyc.Query().
		Where(
			agentkyc.KYCStatusEQ("APPROVED"),
			agentkyc.KYCDate(date),
		).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	i := 0
	for _, v := range result {
		i++
		if v == nil || v.Consumerwalletid == nil {
			continue
		}

		userWallet, err := reportService.GetUserWalletByID(*v.Consumerwalletid)
		if err != nil {
			return nil, err
		}

		wd, err := reportService.connection.ReportWallet.Query().
			Where(reportwallet.Walletid(*v.Consumerwalletid), reportwallet.WalletTypeNameEQ("CUSTOMER")).
			Order(ent.Desc(reportwallet.FieldID)).
			First(context.Background())
		if err != nil {
			reportService.logger.Errorln("[Routine] Get wallet daily error", err)
			return nil, err
		}
		if userWallet == nil {
			_, err = reportService.connection.Userwallet.
				Create().
				SetWalletid(wd.Walletid).
				SetNillableWalletTypeName(wd.WalletTypeName).
				SetNillableWalletPhoneno(wd.WalletPhoneno).
				SetNillableWalletName(wd.WalletName).
				SetNillableRegisterDate(wd.DateTime).
				SetNillableCitizenId(wd.CitizenId).
				SetNillableStatus(wd.Status).
				SetNillableATMCard(wd.ATMCard).
				SetNillableAccountNo(wd.AccountNo).
				SetNillableAddressDetail(wd.AddressDetail).
				SetNillableStreet(wd.Street).
				SetNillableDistrict(wd.District).
				SetNillableSubDistrict(wd.SubDistrict).
				SetNillableProvince(wd.Province).
				SetNillablePostalCode(wd.PostalCode).
				SetUpdateDate(time.Now()).
				SetIsKYC("Y").
				Save(context.Background())
			if err != nil {
				log.Println("Insert ReportWallet has error ", err)
				return nil, err
			}
			continue
		}

		userWallet.Update().
			SetNillableWalletName(wd.WalletName).
			SetNillableRegisterDate(wd.RegisterDateTime).
			SetNillableCitizenId(wd.CitizenId).
			SetNillableStatus(wd.Status).
			SetNillableATMCard(wd.ATMCard).
			SetNillableAccountNo(wd.AccountNo).
			SetNillableAddressDetail(wd.AddressDetail).
			SetNillableStreet(wd.Street).
			SetNillableDistrict(wd.District).
			SetNillableSubDistrict(wd.SubDistrict).
			SetNillableProvince(wd.Province).
			SetNillablePostalCode(wd.PostalCode).
			SetUpdateDate(time.Now()).
			SetIsKYC("Y").
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		continue
	}

	return "success", nil
}
