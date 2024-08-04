package utils

func IsStringDateEmpty(value string) bool {
	return value == "" ||
		value == "0000-00-00T00:00:00" ||
		value == "0000-00-00" ||
		value == "0001-01-01 00:00:00 +0000 UTC"
}
