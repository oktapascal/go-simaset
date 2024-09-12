package helper

import (
	"time"
)

func getTimezone() *time.Location {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	return location
}

func GetTime() time.Time {
	timezone := getTimezone()

	return time.Now().In(timezone)
}
