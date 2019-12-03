package datepearls

import "time"

func numberDaysBetweenTwoDatesGoWay(date1, date2 time.Time) int {

	if date1.After(date2) {
		date1, date2 = date2, date1
	}

	return int(date2.Sub(date1).Hours() / 24)
}
