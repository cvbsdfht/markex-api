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

func IsStringDateEmpty(value string) bool {
	return value == "" ||
		value == "0000-00-00T00:00:00" ||
		value == "0000-00-00" ||
		value == "0001-01-01 00:00:00 +0000 UTC"
}

const (
	FORMAT_DATE = "2006-01-02"
	FORMAT_TIME = "15:04:05"
	FORMAT_DATETIME = "2006-01-02 15:04:05"
)
