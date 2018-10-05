package datetime

import (
	"os"
)

const (
	WI_DATETIME_LOCATION         string = "WI_DATETIME_LOCATION"
	WI_DATETIME_LOCATION_DEFAULT string = "UTC"

	WI_DATETIME_FORMAT         string = "WI_DATETIME_FORMAT"
	WI_DATETIME_FORMAT_DEFAULT string = "2006-01-02T15:04:05-07:00"

	WI_DATE_FORMAT         string = "WI_DATE_FORMAT"
	WI_DATE_FORMAT_DEFAULT string = "2006-01-02-07:00"

	WI_TIME_FORMAT         string = "WI_TIME_FORMAT"
	WI_TIME_FORMAT_DEFAULT string = "15:04:05-07:00"
)

func GetLocation() string {
	location, ok := os.LookupEnv(WI_DATETIME_LOCATION)
	if ok && location != "" {
		return location
	}
	return WI_DATETIME_LOCATION_DEFAULT
}

func GetDatetimeFormat() string {
	format, ok := os.LookupEnv(WI_DATETIME_FORMAT)
	if ok && format != "" {
		return format
	}
	return WI_DATETIME_FORMAT_DEFAULT
}

func GetDateFormat() string {
	format, ok := os.LookupEnv(WI_DATE_FORMAT)
	if ok && format != "" {
		return format
	}
	return WI_DATE_FORMAT_DEFAULT
}

func GetTimeFormat() string {
	format, ok := os.LookupEnv(WI_TIME_FORMAT)
	if ok && format != "" {
		return format
	}
	return WI_TIME_FORMAT_DEFAULT
}
