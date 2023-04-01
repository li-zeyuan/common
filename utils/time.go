package utils

import "time"

func NowUTC() time.Time {
	return time.Now().UTC()
}

func TodayStartUTC() time.Time {
	t := time.Now()

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UTC()
}

func TimeStamp2Time(tStamp int64) time.Time {
	return time.Unix(tStamp, 0)
}

func Time2TimeStamp(t time.Time) int64 {
	return t.Unix()
}

func ParseDay(tStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, tStr)
}
