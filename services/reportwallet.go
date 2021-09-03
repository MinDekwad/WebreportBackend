package services

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/reportwallet"
	"go-api-report2/ent/userwallet"
	"time"
)

func (r ReportService) GetUpdateReportWallet() error {

	lastDay, err := r.connection.Agentkyc.Query().
		Order(ent.Desc(agentkyc.FieldID)).
		Select(agentkyc.FieldKYCDate).
		First(context.Background())
	if err != nil {
		r.logger.Errorln("[Routine] Get AgentKYC lastDay error", err)
		return err
	}

	r.logger.Debugln("[Routine] KYC last date is : ", lastDay)

	result, err := r.connection.Agentkyc.Query().
		Where(agentkyc.KYCStatusEQ("APPROVED"), agentkyc.KYCDate(*lastDay.KYCDate)).
		All(context.Background())
	if err != nil {
		r.logger.Errorln("[Routine] Get AgentKYC lastDay status APPROVED error", err)
		return err
	}

	r.logger.Debugln("[Routine] total lastDay of kyc approve : ", len(result))

	i := 0
	for _, v := range result {
		if v == nil || v.Consumerwalletid == nil {
			r.logger.Debugln("[Routine] AgentKYC is null or CustomerWalletID is null on lastDay")
			continue
		}

		userWallet, err := r.GetUserWalletByID(*v.Consumerwalletid)
		if err != nil {
			r.logger.Errorln("[Routine] Get report wallet lastDay error ", err)
			return err
		}

		wd, err := r.connection.ReportWallet.Query().
			Where(reportwallet.Walletid(*v.Consumerwalletid), reportwallet.WalletTypeNameEQ("CUSTOMER")).
			Order(ent.Desc(reportwallet.FieldDateTime)).
			First(context.Background())

		if err != nil {
			r.logger.Errorln("[Routine] Get wallet daily lastDay error", err)
			return err
		}
		if userWallet == nil {
			_, err = r.connection.Userwallet.
				Create().
				SetWalletid(wd.Walletid).
				SetNillableWalletTypeName(wd.WalletTypeName).
				SetNillableWalletPhoneno(wd.WalletPhoneno).
				SetNillableWalletName(wd.WalletName).
				//SetNillableRegisterDate(wd.DateTime).
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
				SetOccupationId(0).
				Save(context.Background())
			if err != nil {
				r.logger.Println("Insert ReportWallet lastDay has error ", err)
				return err
			}
			i++
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
			return err
		}
		i++
	}
	r.logger.Debugln("[Routine] Create or update report wallet lastDay success", i)

	/*
		// insert ranking for amlo calculate rank

		result2, err := r.connection.Agentkyc.Query().
			Where(agentkyc.KYCStatusEQ("APPROVED"), agentkyc.KYCDate(*lastDay.KYCDate)).
			All(context.Background())
		if err != nil {
			r.logger.Errorln("[Routine] Get AgentKYC lastDay status APPROVED error", err)
			return err
		}

		// fmt.Println()
		// fmt.Println("len agent kyc is : ", len(result2))
		// fmt.Println()
		ii := 0
		for _, agkyc := range result2 {
			if agkyc == nil || agkyc.Consumerwalletid == nil {
				r.logger.Debugln("[Routine] Ranking AgentKYC is null or CustomerWalletID is null on lastDay")
				continue
			}

			ranking, err := r.GetRankingByID(*agkyc.Consumerwalletid)
			if err != nil {
				r.logger.Errorln("[Routine] Get ranking error ", err)
				return err
			}

			wdRanking, err := r.connection.ReportWallet.Query().
				Where(reportwallet.Walletid(*agkyc.Consumerwalletid), reportwallet.WalletTypeNameEQ("CUSTOMER")).
				Order(ent.Desc(reportwallet.FieldDateTime)).
				First(context.Background())

			// if wdRanking.District != nil {
			// 	cfArea, err := r.connection.Configarea.Query().
			// 		Select(configarea.FieldDistrictNameEN).
			// 		Where(configarea.DistrictNameTH(*wdRanking.District)).
			// 		First(context.Background())
			// 	if err != nil {
			// 		return err
			// 	}
			// 	fmt.Println()
			// 	fmt.Println("cfArea : ", cfArea)
			// 	fmt.Println()
			// }

			cfArea, err := r.connection.Configarea.Query().
				Select(configarea.FieldDistrictNameEN).
				Where(configarea.DistrictNameTH(*wdRanking.District)).
				First(context.Background())

			if err != nil {
				r.logger.Errorln("[Routine] Ranking Get wallet daily lastDay error", err)
				return err
			}

			if ranking == nil {
				_, err = r.connection.Ranking.
					Create().
					SetWalletID(wdRanking.Walletid).
					SetName(*wdRanking.WalletName).
					SetTaxID(*wdRanking.CitizenId).
					SetProvinceNameTH(*wdRanking.Province).
					SetDistrictNameTH(*wdRanking.District).
					SetDistrictNameEN(cfArea.DistrictNameEN).
					SetOccupationName("other"). // ถ้ามีข้อมูลแล้วแก้ไขด้วย
					SetLastRank(0).
					SetCurrentRank(1).
					SetStatusRanking("New").
					SetLastDateCalRank("-").
					SetNextDateCalRank("-").
					SetStateCal(1).
					Save(context.Background())
				if err != nil {
					r.logger.Println("Insert ReportWallet lastDay has error ", err)
					return err
				}
				ii++
				continue
			}

			ranking.Update().
				SetProvinceNameTH(*wdRanking.Province).
				SetDistrictNameTH(*wdRanking.District).
				SetDistrictNameEN(cfArea.DistrictNameEN).
				SetOccupationName("-"). // ถ้ามีข้อมูลแล้วแก้ไขด้วย
				Save(context.Background())
			if err != nil {
				return err
			}
			ii++
		}
		r.logger.Debugln("[Routine] Create or update ranking lastDay success", ii)
	*/

	// new
	ld, err := time.Parse(timeCustomLayout, *lastDay.KYCDate)
	if err != nil {
		return err
	}

	rdDay := ld.AddDate(0, 0, -1)
	thirdDay := rdDay.Format("2006-01-02")

	ndDay := ld.AddDate(0, 0, -2)
	secondDay := ndDay.Format("2006-01-02")

	stDay := ld.AddDate(0, 0, -3)
	firstDay := stDay.Format("2006-01-02")

	r.logger.Debugln("[Routine] KYC third date is : ", thirdDay)
	resultThird, err := r.connection.Agentkyc.Query().
		Where(agentkyc.KYCStatusEQ("APPROVED"), agentkyc.KYCDate(thirdDay)).
		All(context.Background())
	if err != nil {
		r.logger.Errorln("[Routine] Get AgentKYC thirdDay status APPROVED error", err)
		return err
	}
	r.logger.Debugln("[Routine] total of kyc thirdDay approve : ", len(resultThird))
	j := 0
	for _, rt := range resultThird {
		if rt == nil || rt.Consumerwalletid == nil {
			r.logger.Debugln("[Routine] AgentKYC thirdDay is null or CustomerWalletID is null")
			continue
		}

		userWalletThirdDay, err := r.GetUserWalletByID(*rt.Consumerwalletid)
		if err != nil {
			r.logger.Errorln("[Routine] Get report wallet thirdDay error ", err)
			return err
		}

		wdThirdDay, err := r.connection.ReportWallet.Query().
			Where(reportwallet.Walletid(*rt.Consumerwalletid), reportwallet.WalletTypeNameEQ("CUSTOMER")).
			Order(ent.Desc(reportwallet.FieldDateTime)).
			First(context.Background())
		if err != nil {
			r.logger.Errorln("[Routine] Get wallet daily thirdDay error", err)
			return err
		}

		if userWalletThirdDay == nil {
			_, err = r.connection.Userwallet.
				Create().
				SetWalletid(wdThirdDay.Walletid).
				SetNillableWalletTypeName(wdThirdDay.WalletTypeName).
				SetNillableWalletPhoneno(wdThirdDay.WalletPhoneno).
				SetNillableWalletName(wdThirdDay.WalletName).
				//SetNillableRegisterDate(wdThirdDay.DateTime).
				SetNillableRegisterDate(wdThirdDay.RegisterDateTime).
				SetNillableCitizenId(wdThirdDay.CitizenId).
				SetNillableStatus(wdThirdDay.Status).
				SetNillableATMCard(wdThirdDay.ATMCard).
				SetNillableAccountNo(wdThirdDay.AccountNo).
				SetNillableAddressDetail(wdThirdDay.AddressDetail).
				SetNillableStreet(wdThirdDay.Street).
				SetNillableDistrict(wdThirdDay.District).
				SetNillableSubDistrict(wdThirdDay.SubDistrict).
				SetNillableProvince(wdThirdDay.Province).
				SetNillablePostalCode(wdThirdDay.PostalCode).
				SetUpdateDate(time.Now()).
				SetIsKYC("Y").
				Save(context.Background())
			if err != nil {
				r.logger.Println("Insert ReportWallet thirdDay has error ", err)
				return err
			}
			j++
			continue
		}
		userWalletThirdDay.Update().
			SetNillableWalletName(wdThirdDay.WalletName).
			SetNillableRegisterDate(wdThirdDay.RegisterDateTime).
			SetNillableCitizenId(wdThirdDay.CitizenId).
			SetNillableStatus(wdThirdDay.Status).
			SetNillableATMCard(wdThirdDay.ATMCard).
			SetNillableAccountNo(wdThirdDay.AccountNo).
			SetNillableAddressDetail(wdThirdDay.AddressDetail).
			SetNillableStreet(wdThirdDay.Street).
			SetNillableDistrict(wdThirdDay.District).
			SetNillableSubDistrict(wdThirdDay.SubDistrict).
			SetNillableProvince(wdThirdDay.Province).
			SetNillablePostalCode(wdThirdDay.PostalCode).
			SetUpdateDate(time.Now()).
			SetIsKYC("Y").
			Save(context.Background())
		if err != nil {
			return err
		}
		j++
	}
	r.logger.Debugln("[Routine] Create or update report wallet thirdDay success", j)

	///////////////////////////////////////////
	r.logger.Debugln("[Routine] KYC second date is : ", secondDay)
	resultSecond, err := r.connection.Agentkyc.Query().
		Where(agentkyc.KYCStatusEQ("APPROVED"), agentkyc.KYCDate(secondDay)).
		All(context.Background())
	if err != nil {
		r.logger.Errorln("[Routine] Get AgentKYC secondDay status APPROVED error", err)
		return err
	}
	r.logger.Debugln("[Routine] total of kyc secondDay approve : ", len(resultSecond))
	k := 0
	for _, rs := range resultSecond {
		if rs == nil || rs.Consumerwalletid == nil {
			r.logger.Debugln("[Routine] AgentKYC secondDay is null or CustomerWalletID is null")
			continue
		}

		userWalletSecondDay, err := r.GetUserWalletByID(*rs.Consumerwalletid)
		if err != nil {
			r.logger.Errorln("[Routine] Get report wallet secondDay error ", err)
			return err
		}

		wdSecondDay, err := r.connection.ReportWallet.Query().
			Where(reportwallet.Walletid(*rs.Consumerwalletid), reportwallet.WalletTypeNameEQ("CUSTOMER")).
			Order(ent.Desc(reportwallet.FieldDateTime)).
			First(context.Background())
		if err != nil {
			r.logger.Errorln("[Routine] Get wallet daily secondDay error", err)
			return err
		}

		if userWalletSecondDay == nil {
			_, err = r.connection.Userwallet.
				Create().
				SetWalletid(wdSecondDay.Walletid).
				SetNillableWalletTypeName(wdSecondDay.WalletTypeName).
				SetNillableWalletPhoneno(wdSecondDay.WalletPhoneno).
				SetNillableWalletName(wdSecondDay.WalletName).
				//SetNillableRegisterDate(wdSecondDay.DateTime).
				SetNillableRegisterDate(wdSecondDay.RegisterDateTime).
				SetNillableCitizenId(wdSecondDay.CitizenId).
				SetNillableStatus(wdSecondDay.Status).
				SetNillableATMCard(wdSecondDay.ATMCard).
				SetNillableAccountNo(wdSecondDay.AccountNo).
				SetNillableAddressDetail(wdSecondDay.AddressDetail).
				SetNillableStreet(wdSecondDay.Street).
				SetNillableDistrict(wdSecondDay.District).
				SetNillableSubDistrict(wdSecondDay.SubDistrict).
				SetNillableProvince(wdSecondDay.Province).
				SetNillablePostalCode(wdSecondDay.PostalCode).
				SetUpdateDate(time.Now()).
				SetIsKYC("Y").
				Save(context.Background())
			if err != nil {
				r.logger.Println("Insert ReportWallet secondDay has error ", err)
				return err
			}
			k++
			continue
		}
		userWalletSecondDay.Update().
			SetNillableWalletName(wdSecondDay.WalletName).
			SetNillableRegisterDate(wdSecondDay.RegisterDateTime).
			SetNillableCitizenId(wdSecondDay.CitizenId).
			SetNillableStatus(wdSecondDay.Status).
			SetNillableATMCard(wdSecondDay.ATMCard).
			SetNillableAccountNo(wdSecondDay.AccountNo).
			SetNillableAddressDetail(wdSecondDay.AddressDetail).
			SetNillableStreet(wdSecondDay.Street).
			SetNillableDistrict(wdSecondDay.District).
			SetNillableSubDistrict(wdSecondDay.SubDistrict).
			SetNillableProvince(wdSecondDay.Province).
			SetNillablePostalCode(wdSecondDay.PostalCode).
			SetUpdateDate(time.Now()).
			SetIsKYC("Y").
			Save(context.Background())
		if err != nil {
			return err
		}
		k++
	}
	r.logger.Debugln("[Routine] Create or update report wallet secondDay success", k)

	///////////////////////////////////////////
	r.logger.Debugln("[Routine] KYC first date is : ", firstDay)
	resultFirst, err := r.connection.Agentkyc.Query().
		Where(agentkyc.KYCStatusEQ("APPROVED"), agentkyc.KYCDate(firstDay)).
		All(context.Background())
	if err != nil {
		r.logger.Errorln("[Routine] Get AgentKYC firstDay status APPROVED error", err)
		return err
	}
	r.logger.Debugln("[Routine] total of kyc firstDay approve : ", len(resultFirst))
	l := 0
	for _, rf := range resultFirst {
		if rf == nil || rf.Consumerwalletid == nil {
			r.logger.Debugln("[Routine] AgentKYC firstDay is null or CustomerWalletID is null")
			continue
		}

		userWalletFirstDay, err := r.GetUserWalletByID(*rf.Consumerwalletid)
		if err != nil {
			r.logger.Errorln("[Routine] Get report wallet firstDay error ", err)
			return err
		}

		wdFirstDay, err := r.connection.ReportWallet.Query().
			Where(reportwallet.Walletid(*rf.Consumerwalletid), reportwallet.WalletTypeNameEQ("CUSTOMER")).
			Order(ent.Desc(reportwallet.FieldDateTime)).
			First(context.Background())
		if err != nil {
			r.logger.Errorln("[Routine] Get wallet daily firstDay error", err)
			return err
		}

		if userWalletFirstDay == nil {
			_, err = r.connection.Userwallet.
				Create().
				SetWalletid(wdFirstDay.Walletid).
				SetNillableWalletTypeName(wdFirstDay.WalletTypeName).
				SetNillableWalletPhoneno(wdFirstDay.WalletPhoneno).
				SetNillableWalletName(wdFirstDay.WalletName).
				//SetNillableRegisterDate(wdFirstDay.DateTime).
				SetNillableRegisterDate(wdFirstDay.RegisterDateTime).
				SetNillableCitizenId(wdFirstDay.CitizenId).
				SetNillableStatus(wdFirstDay.Status).
				SetNillableATMCard(wdFirstDay.ATMCard).
				SetNillableAccountNo(wdFirstDay.AccountNo).
				SetNillableAddressDetail(wdFirstDay.AddressDetail).
				SetNillableStreet(wdFirstDay.Street).
				SetNillableDistrict(wdFirstDay.District).
				SetNillableSubDistrict(wdFirstDay.SubDistrict).
				SetNillableProvince(wdFirstDay.Province).
				SetNillablePostalCode(wdFirstDay.PostalCode).
				SetUpdateDate(time.Now()).
				SetIsKYC("Y").
				Save(context.Background())
			if err != nil {
				r.logger.Println("Insert ReportWallet firstDay has error ", err)
				return err
			}
			l++
			continue
		}
		userWalletFirstDay.Update().
			SetNillableWalletName(wdFirstDay.WalletName).
			SetNillableRegisterDate(wdFirstDay.RegisterDateTime).
			SetNillableCitizenId(wdFirstDay.CitizenId).
			SetNillableStatus(wdFirstDay.Status).
			SetNillableATMCard(wdFirstDay.ATMCard).
			SetNillableAccountNo(wdFirstDay.AccountNo).
			SetNillableAddressDetail(wdFirstDay.AddressDetail).
			SetNillableStreet(wdFirstDay.Street).
			SetNillableDistrict(wdFirstDay.District).
			SetNillableSubDistrict(wdFirstDay.SubDistrict).
			SetNillableProvince(wdFirstDay.Province).
			SetNillablePostalCode(wdFirstDay.PostalCode).
			SetUpdateDate(time.Now()).
			SetIsKYC("Y").
			Save(context.Background())
		if err != nil {
			return err
		}
		l++
	}
	r.logger.Debugln("[Routine] Create or update report wallet firstDay success", l)

	return nil
}

func (r ReportService) GetUserWalletByID(walletid string) (*ent.Userwallet, error) {
	rw, err := r.connection.Userwallet.Query().
		Where(userwallet.Walletid(walletid)).
		First(context.Background())
	var errNotfound *ent.NotFoundError
	if err != nil && errors.As(err, &errNotfound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rw, nil
}

// func (r ReportService) GetRankingByID(walletid string) (*ent.Ranking, error) {
// 	rk, err := r.connection.Ranking.Query().
// 		Where(ranking.WalletID(walletid)).
// 		First(context.Background())
// 	var errRankNotfound *ent.NotFoundError
// 	if err != nil && errors.As(err, &errRankNotfound) {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rk, nil
// }
