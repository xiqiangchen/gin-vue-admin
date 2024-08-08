package utils

import "time"

var UTC8 = time.FixedZone("UTC+8", 8*60*60)

func SetDateStart(date time.Time, utc int) time.Time {
	if utc == 0 {
		return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	} else {
		return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, UTC8)

	}
}

func SetDateEnd(date time.Time, utc int) time.Time {
	if utc == 0 {
		return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 59, time.UTC)
	} else {
		return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 59, UTC8)
	}
}
