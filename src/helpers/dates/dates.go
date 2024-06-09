package dates

import "time"

func BeginningOfDay(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), date.Location())
}

func GetCurrentTime() time.Time {
	return time.Now()
}
