package handler

import (
	"go-api-report2/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetImportCsv(c echo.Context) error {
	slug := c.Param("slug")
	filename := c.QueryParam("filename")

	var err error
	var result interface{}

	switch slug {

	case "walletFileName":
		dataWalletFileName, err := h.reportService.GetWalletFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get wallet file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500087,
				Message: "Wallet file name csv has error",
			})
		}
		result = dataWalletFileName

	case "consumerFileName":
		dataConsumerFileName, err := h.reportService.GetConsumerFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get consumer file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500088,
				Message: "Consumer file name csv has error",
			})
		}
		result = dataConsumerFileName

	case "merchantFileName":
		dataMerchantFileName, err := h.reportService.GetMerchantFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get merchant file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500089,
				Message: "Merchant file name csv has error",
			})
		}
		result = dataMerchantFileName

	case "agentKycFileName":
		dataAgentKycFileName, err := h.reportService.GetAgentKycFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get agentkyc file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500090,
				Message: "Agentkyc file name csv has error",
			})
		}
		result = dataAgentKycFileName

	case "loanbindingFileName":
		dataLoanbindingFileName, err := h.reportService.GetLoanbindingFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get loanbinding file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500091,
				Message: "Loanbinding file name csv has error",
			})
		}
		result = dataLoanbindingFileName

	// new
	case "userProfileAmloName":
		dataUserProfileAmloFileName, err := h.reportService.GetUserProfileAmloFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get user pro file amlo file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500152,
				Message: "User profile amlo file name csv has error",
			})
		}
		result = dataUserProfileAmloFileName

	case "MerchantFileimportStatus":
		dataMerchantFileimportStatus, err := h.reportService.GetDataMerchantFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status merchant fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500092,
				Message: "Status merchant fileimport csv has error",
			})
		}
		result = dataMerchantFileimportStatus

	case "MerchantCSV":
		fileimportid := c.QueryParam("fileimportid")
		if fileimportid == "" {
			return c.JSON(http.StatusBadRequest, model.Response{Message: "File import id is requide"})
		}
		Fileimportid, err1 := strconv.Atoi(fileimportid)
		if err1 != nil {
			log.Println("GetReports", err1)
		}
		dataMerchantCSV, err := h.reportService.GetDataMerchantCSV(Fileimportid)
		if err != nil {
			h.logger.Errorf("[Handler] Get data merchant csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500093,
				Message: "Get Data merchant csv has error",
			})
		}
		result = dataMerchantCSV

	case "AgentKycFileimportStatus":
		dataAgentKycFileimportStatus, err := h.reportService.GetDataAgentKycFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status agent kyc fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500094,
				Message: "Status agent kyc fileimport csv has error",
			})
		}
		result = dataAgentKycFileimportStatus

	case "AgentKycCSV":
		fileimportid := c.QueryParam("fileimportid")
		if fileimportid == "" {
			return c.JSON(http.StatusBadRequest, model.Response{Message: "File import id is requide"})
		}
		Fileimportid, err1 := strconv.Atoi(fileimportid)
		if err1 != nil {
			log.Println("GetReports", err1)
		}
		dataAgentKycCSV, err := h.reportService.GetDataAgentKycCSV(Fileimportid)
		if err != nil {
			h.logger.Errorf("[Handler] Get data agent csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500095,
				Message: "Get Data agent csv has error",
			})
		}
		result = dataAgentKycCSV

	case "LoanbindingFileimportStatus":
		dataLoanbindingFileimportStatus, err := h.reportService.GetDataLoanbindingFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status loanbinding fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500096,
				Message: "Status loanbinding fileimport csv has error",
			})
		}
		result = dataLoanbindingFileimportStatus

	case "LoanbindingCSV":
		fileimportid := c.QueryParam("fileimportid")
		if fileimportid == "" {
			return c.JSON(http.StatusBadRequest, model.Response{Message: "File import id is requide"})
		}
		Fileimportid, err1 := strconv.Atoi(fileimportid)
		if err1 != nil {
			log.Println("GetReports", err1)
		}
		dataLoanbindingCSV, err := h.reportService.GetDataLoanbindingCSV(Fileimportid)
		if err != nil {
			h.logger.Errorf("[Handler] Get data loanbinding csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500097,
				Message: "Get Data loanbinding csv has error",
			})
		}
		result = dataLoanbindingCSV

	case "WalletdailyFileimportStatus":
		dataWalletdailyFileimportStatus, err := h.reportService.GetDataWalletdailyFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status wallet daily fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500098,
				Message: "Status wallet daily fileimport csv has error",
			})
		}
		result = dataWalletdailyFileimportStatus

	case "WalletdailyCSV":
		fileimportid := c.QueryParam("fileimportid")
		if fileimportid == "" {
			return c.JSON(http.StatusBadRequest, model.Response{Message: "File import id is requide"})
		}
		Fileimportid, err1 := strconv.Atoi(fileimportid)
		if err1 != nil {
			log.Println("GetReports", err1)
		}
		dataWalletdailyCSV, err := h.reportService.GetDataWalletdailyCSV(Fileimportid)
		if err != nil {
			h.logger.Errorf("[Handler] Get data wallet daily csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500099,
				Message: "Get Data wallet daily csv has error",
			})
		}
		result = dataWalletdailyCSV

	case "ConsumerFileimportStatus":
		dataConsumerFileimportStatus, err := h.reportService.GetDataConsumerFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status consumer transaction fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500100,
				Message: "Status consumer transaction fileimport csv has error",
			})
		}
		result = dataConsumerFileimportStatus

	case "ConsumerCSV":
		fileimportid := c.QueryParam("fileimportid")
		if fileimportid == "" {
			return c.JSON(http.StatusBadRequest, model.Response{Message: "File import id is requide"})
		}
		Fileimportid, err1 := strconv.Atoi(fileimportid)
		if err1 != nil {
			log.Println("GetReports", err1)
		}

		dataConsumerCSV, err := h.reportService.GetDataConsumerCSV(Fileimportid)
		if err != nil {
			h.logger.Errorf("[Handler] Get data consumer transaction csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500101,
				Message: "Get Data consumer transaction csv has error",
			})
		}
		result = dataConsumerCSV

	// new
	case "userProfileAmloFileimportStatus":
		dataUserProfileAmloFileimportStatus, err := h.reportService.GetDataUserProfileAmloFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status user profile amlo fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500153,
				Message: "Status user profile amlo fileimport csv has error",
			})
		}
		result = dataUserProfileAmloFileimportStatus

	// new
	case "watchlistName":
		dataWatchlistFileName, err := h.reportService.GetUserProfileAmloFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get watchlist file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500176,
				Message: "User watchlist file name csv has error",
			})
		}
		result = dataWatchlistFileName

	case "watchlistFileimportStatus":
		dataWatchlistFileimportStatus, err := h.reportService.GetDataWatchlistFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status watchlist fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500177,
				Message: "Status watchlist fileimport csv has error",
			})
		}
		result = dataWatchlistFileimportStatus

	case "kycPendingFileName":
		dataKycPendingFileName, err := h.reportService.GetKycPendingFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get kyc pending file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500088,
				Message: "Kyc pending file name csv has error",
			})
		}
		result = dataKycPendingFileName

	case "KycPendingFileimportStatus":
		dataKycPendingFileimportStatus, err := h.reportService.GetDataKycPendingFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status kyc pending fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500100,
				Message: "Status kyc pending fileimport csv has error",
			})
		}
		result = dataKycPendingFileimportStatus

	case "lbPendingFileName":
		dataLbPendingFileName, err := h.reportService.GetLbPendingFileName(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get loanbinding pending file name csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500088,
				Message: "Status loanbinding pending file name csv has error",
			})
		}
		result = dataLbPendingFileName

	case "LbPendingFileimportStatus":
		dataLbPendingFileimportStatus, err := h.reportService.GetDataLbPendingFileimportStatus(filename)
		if err != nil {
			h.logger.Errorf("[Handler] Get status rv pending fileimport csv has error %s", err)
			return c.JSON(http.StatusInternalServerError, model.Response{
				Code:    500100,
				Message: "Status rv pending fileimport csv has error",
			})
		}
		result = dataLbPendingFileimportStatus

	/*
		case "UserProfileAmloCSV":
			fileimportid := c.QueryParam("fileimportid")
			if fileimportid == "" {
				return c.JSON(http.StatusBadRequest, model.Response{Message: "File import id is requide"})
			}
			Fileimportid, err1 := strconv.Atoi(fileimportid)
			if err1 != nil {
				log.Println("GetReports", err1)
			}

			dataConsumerCSV, err := h.reportService.GetDataConsumerCSV(Fileimportid)
			if err != nil {
				h.logger.Errorf("[Handler] Get data consumer transaction csv has error %s", err)
				return c.JSON(http.StatusInternalServerError, model.Response{
					Code:    500101,
					Message: "Get Data consumer transaction csv has error",
				})
			}
			result = dataConsumerCSV
	*/

	default:
		return c.JSON(http.StatusBadRequest, model.Response{Message: "Error", Data: "This import csv has error"})
	}

	if err != nil {
		h.logger.Errorf("[Handler] Get import csv has error %s", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500102,
			Message: "Get Data import csv has error",
		})
	}
	return c.JSON(http.StatusOK, model.Response{Message: "Success", Data: result})
}
