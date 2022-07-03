package action

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
)

func FormatDate(tanggal string) datatypes.Date {
	var layoutFormat string
	var date time.Time
	layoutFormat = "2006-01-02"
	date, _ = time.Parse(layoutFormat, tanggal)
	result := datatypes.Date(date)
	// date.Format("02-Jan-2006")
	return result
}

func FormatDateToString(result datatypes.Date) string {
	toTime := time.Time(result)
	toFormat := toTime.Format("2006-01-02")
	return toFormat
}

func FormatTimeBiasa(waktu datatypes.Time) string {
	return time.Duration(waktu).String()
}

func FormatTime(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
