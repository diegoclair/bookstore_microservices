package dateutils

import (
	"time"
)

const (
	apiDateLayout  = "2006-01-02T15:04:05Z"
	onlyDateLayout = "2006-01-02"
	apiDBLayout    = "2006-01-02 15:04:05"
)

// GetDateNowTime return a complete date in format time.Time
//  Ex: 2020-01-31 04:00:55.6855191 +0000 UTC
func GetDateNowTime() time.Time {
	return time.Now().UTC()
}

// GetCompleteDateNowString return a complete date in format string.
//  Ex: "2020-01-31T15:04:05Z"
func GetCompleteDateNowString() string {
	return GetDateNowTime().Format(apiDateLayout)
}

// GetCompleteDateNowDBLayout return a complete date in format to database.
//  Ex: "2020-01-31 15:04:05"
// Generally is used to save date on database
func GetCompleteDateNowDBLayout() string {
	return GetDateNowTime().Format(apiDBLayout)
}

// GetOnlyDateNowString return a date in format string.
//  Ex: "2020-01-31"
func GetOnlyDateNowString() string {
	return GetDateNowTime().Format(onlyDateLayout)
}

// GetFirstAndLastOfTheMonth return a first and last day of a date in format string.
//  Ex: 2020-02-01, 2020-02-29
// The parameter is a number:
//  1 to January
//  2 to February
//  ...
//  12 to December;
func GetFirstAndLastOfTheMonth(month time.Month) (firstDay, lastDat string) {
	y, _, _ := GetDateNowTime().Date()
	firstDay = time.Date(y, month, 1, 0, 0, 0, 0, time.UTC).Format(onlyDateLayout)
	lastDat = time.Date(y, month+1, 0, 0, 0, 0, 0, time.UTC).Format(onlyDateLayout)

	return firstDay, lastDat
}
