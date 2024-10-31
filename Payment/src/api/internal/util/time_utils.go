package util

import (
	"time"
)

var location *time.Location

func InitTimezone(TZ string) error {
	loc, err := time.LoadLocation(TZ)
	if err != nil {
		return err
	}
	location = loc
	return nil
}

func Now() time.Time {
	return time.Now().In(location)
}

func ApplyLocalTZ(timestamp time.Time) time.Time {
	return timestamp.In(location)
}
