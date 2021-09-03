package services

import (
	"context"
	"errors"
	"strconv"
	"time"
)

func (reportService ReportService) CreateLogExportPointCSV(user, AdminID, Resource, Action, start_date, end_date string) error {
	if AdminID == "" {
		return errors.New("AdminID is required")
	}

	adminID, err := strconv.Atoi(AdminID)
	if err != nil {
		return errors.New("Cannot change type AdminID from string to integer")
	}

	if Resource == "" {
		return errors.New("Resource is required")
	}
	if Action == "" {
		return errors.New("Action is required")
	}

	_, err = reportService.connection.Writelog.Create().
		SetAdminID(adminID).
		SetResource(Resource).
		SetActions(Action).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		reportService.logger.Errorf("[Handler]Error cannot insert log", err)
	}

	if start_date == "" || end_date == "" {
		return errors.New("Date is required")
	}

	// fileName := "ADJUST_POINT_IN_" + start_date + "_" + end_date + ".csv"
	fileName := "ADJUST_POINT_IN_" + start_date + ".csv"

	_, err = reportService.connection.Logexport.Create().
		SetUserName(user).
		SetFileName(fileName).
		SetExportDate(time.Now()).
		Save(context.Background())
	if err != nil {
		reportService.logger.Errorf("[Handler]Error cannot insert log", err)
	}

	return nil
}
