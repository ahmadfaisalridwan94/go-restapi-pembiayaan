package dates

import (
	"time"
)

func ConvertUnixToDatetime(unixtt int64) string {
	unixTimestamp := int64(unixtt)
	timeObj := time.Unix(unixTimestamp, 0)
	dateTimeString := timeObj.Format("2006-01-02 15:04:05")

	return dateTimeString
}
