package timeHelper

import (
	"fmt"
	"time"
)

var loc *time.Location

func SetTimeZone(timezone string) error {
	location, err := time.LoadLocation(timezone)

	if err != nil {
		return err
	}

	loc = location

	return nil
}

func GetTime(t time.Time) time.Time {
	return t.In(loc)
}

func Now() time.Time {
	fmt.Println(loc)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}
