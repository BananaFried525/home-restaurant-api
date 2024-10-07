package utils

import (
	"fmt"
	"time"
)

func PaddingWithZero(number int) string {
	return fmt.Sprintf("%04d", number)
}

func CreateRunningNumber(number int) string {
	now := time.Now()
	return fmt.Sprintf("%s%s", now.Format("060102"), PaddingWithZero(number))
}

func GetStartOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func GetEndOfMonth(start *time.Time) time.Time {
	if start != nil {
		return start.AddDate(0, 1, 0)
	}

	return GetStartOfMonth().AddDate(0, 1, 0)
}

func GetStartEndOfMonth() (start time.Time, end time.Time) {
	start = GetStartOfMonth()
	end = GetEndOfMonth(&start)

	return
}
