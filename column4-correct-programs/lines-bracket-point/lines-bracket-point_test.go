// +build all column4 2lines point linesBracketPoint

package linesbracketpoint

import (
	"fmt"
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
		if err := a.AssertDeepEqual(found, true); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
		if err := a.AssertDeepEqual(lines, test.result); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1,
				m.ErrorMessage(
					fmt.Sprintf("(%v,%v)", test.result[0], test.result[1]),
					fmt.Sprintf("(%v,%v)", lines[0], lines[1]))))
		}
	}
}

func TestLineToString(t *testing.T) {
	if err := a.AssertDeepEqual("2.000x + 4.000", fmt.Sprint(line(point{x: 2, y: 4}))); err != nil {
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

		if err := a.AssertDeepEqual(found, false); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
		if err := a.AssertDeepEqual(lines, test.result); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1,
				m.ErrorMessage(
					fmt.Sprintf("(%v,%v)", test.result[0], test.result[1]),
					fmt.Sprintf("(%v,%v)", lines[0], lines[1]))))
		}
	}
}
