// +build all column8 turnpike costTravel

package turnpike

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestTurnpike(t *testing.T) {
	costs := newTurnpike([]float64{2.5, 1, 10, 0.5})

	type testData struct {
		from, to           int
		costBetweenStation float64
	}

	tests := []testData{
		testData{from: 0, to: 1, costBetweenStation: 2.5},
		testData{from: 1, to: 0, costBetweenStation: 2.5},
		testData{from: 4, to: 0, costBetweenStation: 14},
		testData{from: 0, to: 4, costBetweenStation: 14},
		testData{from: 0, to: 4, costBetweenStation: 14},
		testData{from: 1, to: 4, costBetweenStation: 11.5},
		testData{from: 2, to: 1, costBetweenStation: 1},
		testData{from: 1, to: 3, costBetweenStation: 11},
		testData{from: 1, to: 1, costBetweenStation: -1},
		testData{from: -1, to: 1, costBetweenStation: -1},
		testData{from: 1, to: -1, costBetweenStation: -1},
		testData{from: 1, to: 5, costBetweenStation: -1},
		testData{from: 5, to: 0, costBetweenStation: -1},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.costBetweenStation, costs.getCostBetweenStation(test.from, test.to)); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestTurnpikeLength(t *testing.T) {

	type testData struct {
		costs              turnpike
		costBetweenStation float64
	}

	tests := []testData{
		testData{costs: newTurnpike([]float64{2.5}), costBetweenStation: -1},
		testData{costs: newTurnpike([]float64{}), costBetweenStation: -1},
		testData{costs: newTurnpike(nil), costBetweenStation: -1},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.costBetweenStation, test.costs.getCostBetweenStation(0, 1)); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
