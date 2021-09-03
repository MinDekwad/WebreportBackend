package services

import "time"

func (reportService ReportService) GetStartAndEndOfMonth(date string) (start, end string, err error) {
	d, err := time.Parse(timeCustomLayout, date)
	if err != nil {
		return "", "", err
	}
	year, month, _ := d.Date()
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, d.Location())
	endDate := time.Date(year, month+1, 1, 0, 0, 0, 0, d.Location())
	start = startDate.Format(timeCustomLayout)
	end = endDate.Format(timeCustomLayout)
	return start, end, nil
}
