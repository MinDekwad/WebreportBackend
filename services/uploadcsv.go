package services

import (
	"context"
	"go-api-report2/model"
	"sync"

	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func (reportService ReportService) PostDataUpload(data []interface{}, len int, slug, filename string, user string) error {
	fileImport, err := reportService.CreateFileImport(filename, slug)
	if err != nil {
		return err
	}

	page, perPage := reportService.getPage(len)

	var wg sync.WaitGroup
	log.Println("[InsertWalletDaily] After assign data to channel")
	for i := 0; i < page; i++ {
		low := i * perPage
		high := 0

		high = low + perPage
		if i == (page - 1) {
			high = len
		}

		if slug == "userprofileamlocsv" {
			wg.Add(1)
			go func() {
				effect, err := reportService.InsertUserProfileAmlo(&wg, data[low:high], fileImport.ID)
				if err != nil {
					reportService.logger.Errorln("[Service] insert user profile error", err)
				} else {
					reportService.logger.Infoln("[Service] insert success efffect", effect, "rows")
				}
			}()
		}

		switch slug {
		case "walletcsv":
			go func() {
				channel <- model.Reports{
					WalletDaily:  data[low:high],
					Total:        len,
					FileImportID: fileImport.ID,
				}
			}()
		case "consumercsv":
			go func() {
				channel <- model.Reports{
					Consumer:     data[low:high],
					Total:        len,
					FileImportID: fileImport.ID,
				}
			}()

		case "merchantcsv":
			go func() {
				channel <- model.Reports{
					Merchant:     data[low:high],
					Total:        len,
					FileImportID: fileImport.ID,
				}
			}()

		case "agentkyccsv":
			go func() {
				channel <- model.Reports{
					AgentKycCSV:  data[low:high],
					Total:        len,
					FileImportID: fileImport.ID,
				}
			}()

		case "loanbindingcsv":
			go func() {
				channel <- model.Reports{
					LoanbindingCSV: data[low:high],
					Total:          len,
					FileImportID:   fileImport.ID,
				}
			}()

		case "kycpendingcsv":
			go func() {
				channel <- model.Reports{
					Kycpendingcsv: data[low:high],
					Total:         len,
					FileImportID:  fileImport.ID,
				}
			}()

		case "lbpendingcsv":
			go func() {
				channel <- model.Reports{
					Lbpendingcsv: data[low:high],
					Total:        len,
					FileImportID: fileImport.ID,
				}
			}()

			// case "c":
		// 	go func() {
		// 		channel <- model.Reports{
		// 			UserProfileAmloCSV: data[low:high],
		// 			Total:              len,
		// 			FileImportID:       fileImport.ID,
		// 		}
		// 	}()

		case "watchlistcsv":
			go func() {
				channel <- model.Reports{
					WatchlistImportCSV: data[low:high],
					Total:              len,
					FileImportID:       fileImport.ID,
					User:               user,
				}
			}()

		}

	}
	if slug == "userprofileamlocsv" {
		wg.Wait()

		_, err = reportService.connection.Fileimport.
			UpdateOneID(fileImport.ID).
			SetStatus("Success").
			Save(context.Background())

	}

	log.Println("[InsertCSV] End assign data to channel")

	return nil
}

func (reportService ReportService) getPage(count int) (int, int) {
	perPage := 2000
	page := count / perPage
	mod := count % perPage
	if mod > 0 {
		page = page + 1
	}
	return page, perPage
}
