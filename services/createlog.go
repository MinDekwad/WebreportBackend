package services

import (
	"context"
	"errors"
	"strconv"
	"time"
)

func (reportService ReportService) CreateLog(AdminID string, Resource string, Action string) error {
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

	return nil
}
