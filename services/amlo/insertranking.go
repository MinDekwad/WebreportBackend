package amlo

import (
	"context"
	"errors"
	"go-api-report2/ent"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/configarea"
	"go-api-report2/ent/configoccupation"
	"go-api-report2/ent/ranking"
	"go-api-report2/ent/userwallet"
)

func (r AmloService) InsertRanking() (interface{}, error) {

	lastDay, err := r.connection.Agentkyc.Query().
		Order(ent.Desc(agentkyc.FieldID)).
		Select(agentkyc.FieldKYCDate).
		First(context.Background())
	if err != nil {
		r.logger.Errorln("[Handler] Get AgentKYC lastDay error", err)
		return nil, err
	}

	r.logger.Debugln("[Handler] KYC last date is : ", lastDay)

	ak, err := r.connection.Agentkyc.Query(). // จำนวน agentkyc ที่ approve ทั้งหมดจากวันล่าสุด
							Where(agentkyc.KYCStatusEQ("APPROVED"), agentkyc.KYCDate(*lastDay.KYCDate)).
							All(context.Background())
	if err != nil {
		r.logger.Errorln("[Handler] Get AgentKYC lastDay status APPROVED error", err)
		return nil, err
	}

	r.logger.Debugln("[Handler] total lastDay of kyc approve : ", len(ak))

	i := 0
	for _, agentk := range ak {
		if agentk == nil || agentk.Consumerwalletid == nil {
			r.logger.Debugln("[Handler] AgentKYC is null or CustomerWalletID is null on lastDay")
			continue
		}

		// ranking, err := r.GetRankingByID(*agentk.Consumerwalletid)
		// if err != nil {
		// 	r.logger.Errorln("[Handler] Get data ranking error ", err)
		// 	return nil, err
		// }

		rw, err := r.connection.Userwallet.Query().
			Where(userwallet.Walletid(*agentk.Consumerwalletid), userwallet.WalletTypeNameEQ("CUSTOMER")).
			//Order(ent.Desc(userwallet.FieldUpdateDate)).
			Order(ent.Asc(userwallet.FieldWalletid)).
			First(context.Background())
		if err != nil {
			return nil, err
		}

		cfArea, err := r.connection.Configarea.Query().
			Select(configarea.FieldDistrictNameEN).
			Where(configarea.ZipCode(*rw.PostalCode)).
			First(context.Background())
		if err != nil {
			return nil, err
		}

		var occuName string
		occuName = "Other"
		if rw.OccupationId != 0 {
			cfOccu, err := r.connection.Configoccupation.Query().
				Select(configoccupation.FieldOccupationName).
				Where(configoccupation.ID(*&rw.OccupationId)).
				First(context.Background())
			if err != nil {
				return nil, err
			}
			occuName = cfOccu.OccupationName
		}

		ranking, err := r.GetRankingByID(*agentk.Consumerwalletid)
		if err != nil {
			r.logger.Errorln("[Handler] Get data ranking error ", err)
			return nil, err
		}

		if ranking == nil {
			_, err = r.connection.Ranking.
				Create().
				SetWalletID(rw.Walletid).
				SetName(*rw.WalletName).
				SetTaxID(*rw.CitizenId).
				SetProvinceNameTH(*rw.Province).
				SetDistrictNameTH(*rw.District).
				SetDistrictNameEN(cfArea.DistrictNameEN).
				SetOccupationName(occuName).
				SetLastRank(0).
				SetCurrentRank(1).
				SetStatusRanking("New").
				SetLastDateCalRank("-").
				SetNextDateCalRank("-").
				SetStateCal(1).
				SetZipCode(*rw.PostalCode).
				SetTransactionFactorRank(1).
				SetRegisDate(*rw.RegisterDate).
				SetSubDistrict(*rw.SubDistrict).
				SetPhoneno(*rw.WalletPhoneno).
				SetAddressDetail(*rw.AddressDetail).
				SetStreet(*rw.Street).
				Save(context.Background())
			if err != nil {
				r.logger.Println("Insert ReportWallet lastDay has error ", err)
				return err, err
			}
			i++
			continue
		}
		ranking.Update().
			SetProvinceNameTH(*rw.Province).
			SetDistrictNameTH(*rw.District).
			SetDistrictNameEN(cfArea.DistrictNameEN).
			SetOccupationName(occuName).
			SetZipCode(*rw.PostalCode).
			SetSubDistrict(*rw.SubDistrict).
			SetPhoneno(*rw.WalletPhoneno).
			SetAddressDetail(*rw.AddressDetail).
			SetStreet(*rw.Street).
			Save(context.Background())
		if err != nil {
			return err, err
		}
		i++
	}
	r.logger.Debugln("[Handler] Create or update ranking lastDay success", i)

	return nil, nil
}

func (r AmloService) GetRankingByID(walletid string) (*ent.Ranking, error) {
	rk, err := r.connection.Ranking.Query().
		Where(ranking.WalletID(walletid)).
		First(context.Background())
	var errRankNotfound *ent.NotFoundError
	if err != nil && errors.As(err, &errRankNotfound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rk, nil
}
