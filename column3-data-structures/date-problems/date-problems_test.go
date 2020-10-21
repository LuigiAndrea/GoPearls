// +build all column3 date datepearls

package datepearls

import (
	"testing"
	"time"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestNumberDaysBetweenTwoDay(t *testing.T) {
	type testData struct {
		date1, date2 time.Time
		days         int
	}

	tests := []testData{
		testData{date1: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2000, 2, 1, 0, 0, 0, 0, time.UTC),
			days:  31},
		testData{date1: time.Date(2000, 3, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2000, 2, 1, 0, 0, 0, 0, time.UTC),
			days:  29},
		testData{date1: time.Date(1996, 2, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(1997, 2, 1, 0, 0, 0, 0, time.UTC),
			days:  366},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.days, NumberDaysBetweenTwoDates(test.date1, test.date2)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestGetDayOfWeek(t *testing.T) {
	type testData struct {
		date      time.Time
		dayOfWeek string
	}

	tests := []testData{
		testData{date: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			dayOfWeek: "Saturday"},
		testData{date: time.Date(2022, 7, 12, 0, 0, 0, 0, time.UTC),
			dayOfWeek: "Tuesday"},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.dayOfWeek, GetDayOfWeek(test.date).String()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestComputeCalendar(t *testing.T) {
	type testData struct {
		year, daysInMonth, in_day int
		month                     time.Month
		out_day                   string
	}

	tests := []testData{
		testData{year: 2019, month: time.Month(3), daysInMonth: 31, in_day: 4, out_day: "Tuesday"},
		testData{year: 2019, month: time.Month(6), daysInMonth: 30, in_day: 29, out_day: "Sunday"},
		testData{year: 2016, month: time.Month(2), daysInMonth: 29, in_day: 28, out_day: "Monday"},
	}

	for i, test := range tests {
		cal := ComputeCalendar(test.year, test.month)
		if err := a.AssertDeepEqual(test.daysInMonth, len(cal)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}

		if err := a.AssertDeepEqual(test.out_day, cal[test.in_day].String()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
