package date

import (
	"log"
	"time"
)

// FormatYYYYMMDD formats a date from the given layout & value to YYYY-MM-DD
func FormatYYYYMMDD(layout, value string) string {
	dt, err := time.Parse(layout, value)
	if err != nil {
		log.Panic(err, layout, value)
	}
	if dt.After(time.Now().AddDate(-18, 0, 0)) {
		value = dt.AddDate(-100, 0, 0).UTC().Format("2006-01-02")
	} else {
		value = dt.UTC().Format("2006-01-02")
	}
	return value
}
