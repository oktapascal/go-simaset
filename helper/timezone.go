package helper

import (
	"github.com/spf13/viper"
	"time"
)

func getTimezone() *time.Location {
	location, err := time.LoadLocation(viper.GetString("TZ"))
	if err != nil {
		panic(err)
	}

	return location
}

func GetTime() time.Time {
	timezone := getTimezone()

	return time.Now().In(timezone)
}
