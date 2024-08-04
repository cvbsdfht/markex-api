package utils

import "time"

func StringToDateFormat(date string, format string) time.Time {
	if IsStringDateEmpty(date) {
		return time.Time{}
	}
	
	t, err := time.ParseInLocation(format, date, time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

const (
	FORMAT_DATE = "2006-01-02"
	FORMAT_TIME = "15:04:05"
	FORMAT_DATETIME = "2006-01-02 15:04:05"
)
