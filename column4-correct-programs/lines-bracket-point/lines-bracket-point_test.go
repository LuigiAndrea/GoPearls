// +build all column4 2lines point linesBracketPoint

package linesbracketpoint

import (
	"fmt"
	"reflect"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	points []point
	p      point
	result [2]line
}

var points = []point{
	point{x: 0.09},
	point{x: 0.08, y: 1},
	point{x: -0.05, y: 2},
	point{x: 0.05, y: 2.5},
	point{x: 0.3, y: 2.95},
	point{x: 0.005, y: 4},
}

func TestLinesBracketPoint(t *testing.T) {

	tests := []testData{
		testData{points: points,
			p:      point{x: 0.03, y: 2},
			result: [2]line{line(point{x: -0.05, y: 2}), line(point{x: 0.05, y: 2.5})}},
		testData{points: points,
			p:      point{x: 0.9, y: 3.5},
			result: [2]line{line(point{x: 0.3, y: 2.95}), line(point{x: 0.005, y: 4})}},
		testData{points: points,
			p:      point{x: 0.9, y: 0.5},
			result: [2]line{line(point{x: 0.09, y: 0}), line(point{x: 0.08, y: 1})}},
		testData{points: points,
			p:      point{x: 0.9, y: 3.2},
			result: [2]line{line(point{x: 0.05, y: 2.5}), line(point{x: 0.3, y: 2.95})}},
		testData{points: points,
			p:      point{x: 0.6, y: 3.2},
			result: [2]line{line(point{x: 0.3, y: 2.95}), line(point{x: 0.005, y: 4})}},
		testData{points: []point{point{x: 0.86, y: 2}, point{x: 0.86, y: 5}},
			p:      point{x: 0.3, y: 2.5},
			result: [2]line{line(point{x: 0.86, y: 2}), line(point{x: 0.86, y: 5})}},
	}

	for i, test := range tests {
		found, lines := getLinesBracketPoint(test.points, test.p)
		if found == false {
			t.Error(m.ErrorMessage(true, found))
		} else if !reflect.DeepEqual(lines[0], test.result[0]) || !reflect.DeepEqual(lines[1], test.result[1]) {
			t.Error(m.ErrorMessageTestCount(i+1,
				m.ErrorMessage(fmt.Sprintf("(%s,%s)", test.result[0], test.result[1]),
					fmt.Sprintf("(%s,%s)", lines[0], lines[1]))))
		}
	}
}

func TestLineToString(t *testing.T) {
	expected := "2.000x + 4.000"
	actual := line(point{x: 2, y: 4}).String()
	if err := a.AssertDeepEqual(expected, actual); err != nil {
		t.Error(err)
	}
}

func TestPointOutOfRange(t *testing.T) {
	tests := []testData{
		testData{points: points,
			p:      point{x: 1.5, y: 2},
			result: [2]line{}},
		testData{points: points,
			p:      point{x: 0.9, y: 5},
			result: [2]line{}},
		testData{points: []point{point{x: 0.86}},
			p:      point{x: 0.9, y: 2},
			result: [2]line{}},
	}

	for i, test := range tests {
		found, lines := getLinesBracketPoint(test.points, test.p)
		if found == true {
			t.Error(m.ErrorMessage(false, found))
		} else if !reflect.DeepEqual(lines[0], test.result[0]) || !reflect.DeepEqual(lines[1], test.result[1]) {
			t.Error(m.ErrorMessageTestCount(i+1,
				m.ErrorMessage(fmt.Sprintf("(%s,%s)", test.result[0], test.result[1]),
					fmt.Sprintf("(%s,%s)", lines[0], lines[1]))))
		}
	}
}
