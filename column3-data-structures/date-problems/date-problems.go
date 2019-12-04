package datepearls

import "time"

//NumberDaysBetweenTwoDatesGoWay Compute the number of days between two dates
func NumberDaysBetweenTwoDatesGoWay(date1, date2 time.Time) int {

	if date1.After(date2) {
		date1, date2 = date2, date1
	}

	return int(date2.Sub(date1).Hours() / 24)
}

//GetDayOfWeek Given a date, returns its day of the week
func GetDayOfWeek(Date time.Time) time.Weekday {
	return Date.Weekday()
}
