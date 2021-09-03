package ekyc

import "time"

const (
	timeCustomLayout = "2006-01-02"
	layoutISO        = "2006-01-02 00:00:00"
	layoutISOA       = "2006-01-02 15:04:05"
)

func (reportService EkycService) GetStartAndEndOfMonth(date string) (start, end string, err error) {
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
