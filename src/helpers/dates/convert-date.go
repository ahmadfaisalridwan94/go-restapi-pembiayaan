package dates

import (
	"time"
)

func ConvertDate(dateStr, formatStr string) (string, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}

	// Format the date in the 'dmY' format
	return date.Format(formatStr), nil
}
