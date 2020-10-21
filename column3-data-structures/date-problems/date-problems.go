package datepearls

import (
	"time"
)

//NumberDaysBetweenTwoDates Compute the number of days between two dates
func NumberDaysBetweenTwoDates(date1, date2 time.Time) int {

	if date1.After(date2) {
		date1, date2 = date2, date1
	}

	return int(date2.Sub(date1).Hours() / 24)
}

//GetDayOfWeek Given a date, returns its day of the week
func GetDayOfWeek(Date time.Time) time.Weekday {
	return Date.Weekday()
}

//ComputeCalendar Produce a calendar for the month
func ComputeCalendar(Year int, Month time.Month) []time.Weekday {
	dataToCompute := time.Date(Year, Month, 1, 0, 0, 0, 0, time.UTC)
	days := NumberDaysBetweenTwoDates(dataToCompute, time.Date(Year, Month+1, 1, 0, 0, 0, 0, time.UTC))
	weekDayValue := GetDayOfWeek(dataToCompute)
	calendar := make([]time.Weekday, 0)

	for i := 0; i < days; i++ {
		calendar = append(calendar, time.Weekday((int(weekDayValue)+i)%7))
	}

	return calendar
}
