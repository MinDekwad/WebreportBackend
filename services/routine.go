package services

import (
	"context"
)

func (r ReportService) UpdateReportWallet() {
	r.logger.Infoln("Start update report wallet routine")
	for {
		select {
		case b := <-triggerChannel:
			if !b {
				continue
			}
			err := r.GetUpdateReportWallet()
			if err != nil {
				r.logger.Errorln("[Routine] Get update or create report wallet error", err)
				continue
			}
			r.logger.Infoln("[Routine] Get update or create report wallet success")
		}
	}
}

func (r ReportService) InsertCSV() (string, error) {

	r.logger.Info("Start goroutine")

	for {
		isWalletDaily := false
		select {
		case reports := <-channel:
			// wallet daily import csv
			r.logger.Info("After get item from channel")
			totalRow := reports.Total
			fileImportID := reports.FileImportID
			user := reports.User
			count := 0
			var err error

			if len(reports.WalletDaily) > 0 {
				isWalletDaily = true
				count, err = r.InsertWalletDaily(reports.WalletDaily, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] Insert wallet daily error", err)
					continue
				}
			}
			//
			// consumer transaction csv
			if len(reports.Consumer) > 0 {
				count, err = r.InsertConsumer(reports.Consumer, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] insert consumer error", err)
					continue
				}
			}

			// Metchant transaction csv
			if len(reports.Merchant) > 0 {
				count, err = r.InsertMerchant(reports.Merchant, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] insert merchant error", err)
					continue
				}
			}
			//

			// agent kyc csv
			if len(reports.AgentKycCSV) > 0 {
				count, err = r.InsertAgentKYC(reports.AgentKycCSV, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] insert agentkyc error", err)
					continue
				}
			}
			//

			// loanbinding csv
			if len(reports.LoanbindingCSV) > 0 {
				count, err = r.InsertLoanbinding(reports.LoanbindingCSV, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] insert loanbinding error", err)
					continue
				}
			}

			//kycpending csv
			if len(reports.Kycpendingcsv) > 0 {
				count, err = r.InsertKycPending(reports.Kycpendingcsv, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] insert kyc pending error", err)
					continue
				}
			}

			//lbpending csv
			if len(reports.Lbpendingcsv) > 0 {
				count, err = r.InsertLbPending(reports.Lbpendingcsv, fileImportID)
				if err != nil {
					r.logger.Errorln("[Routine] insert lb pending error", err)
					continue
				}
			}

			// watchlist csv
			if len(reports.WatchlistImportCSV) > 0 {
				count, err = r.InsertWatchlist(reports.WatchlistImportCSV, fileImportID, user)
				if err != nil {
					r.logger.Errorln("[Routine] insert watchlist error", err)
					continue
				}
			}

			if count == 00 {
				_, err = r.connection.Fileimport.
					UpdateOneID(fileImportID).
					SetStatus("Dup").
					Save(context.Background())
				if err != nil {
					r.logger.Errorln("[Routine] update status fileimport data duplicate", err)
				}
			}

			if count == totalRow {
				_, err = r.connection.Fileimport.
					UpdateOneID(fileImportID).
					SetStatus("Success").
					Save(context.Background())
				if err != nil {
					r.logger.Errorln("[Routine] update status fileimport error", err)
				}
				if isWalletDaily {
					triggerChannel <- true
				}
			}

		}
	}
}
