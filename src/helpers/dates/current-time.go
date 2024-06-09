package dates

import "time"

func CurrentTime(formatString string) string {
	return time.Now().Format(formatString)
}
