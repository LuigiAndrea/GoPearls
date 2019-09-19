// +build all column2 atleast twice

package twice

import (
	"testing"
)

type testData struct {
	lengthList    int
	indexesToSkip map[int]struct{}
	expectedValue int
}

func TestAtLeastTwice(t *testing.T) {

	tests := []testData{
		testData{lengthList: 20, indexesToSkip: map[int]struct{}{2: {}, 3: {}, 7: {}, 13: {}}, expectedValue: -9},
		testData{lengthList: 10, indexesToSkip: map[int]struct{}{7: {}}, expectedValue: 1},
		testData{lengthList: 2, indexesToSkip: map[int]struct{}{1: {}}, expectedValue: -1},
	}

	for _, test := range tests {

		list := make([]int, test.lengthList)
		populateSequentialList(list, test.lengthList, test.indexesToSkip)

		if ok, v := searchAtLeastTwiceValue(list); ok {
			if v != test.expectedValue {
				t.Errorf("\nExpected value: '%d' - Actual value '%d' \nList: %#v", test.expectedValue, v, list)
			} else {
				t.Logf("Repeated element: '%d' \nList: %#v", v, list)
			}
		} else {
			t.Errorf("Repeated element not detected \nList: %#v", list)
		}
	}
}

func TestEdgeCases(t *testing.T) {

	tests := []testData{
		testData{lengthList: 5, indexesToSkip: nil, expectedValue: 5},
		testData{lengthList: 0, indexesToSkip: nil, expectedValue: 0},
	}

	for _, test := range tests {
		list := make([]int, test.lengthList)

		populateSequentialList(list, test.lengthList, nil)

		if ok, v := searchAtLeastTwiceValue(list); ok {
			t.Errorf("No element appears at least twice in the list \nList: %#v", list)
		} else if v != test.expectedValue {
			t.Errorf("\nExpected value: '%d' - Actual value '%d' \nList: %#v", test.expectedValue, test.lengthList, list)
		}
	}
}

type skipIndexes map[int]struct{}

func (s skipIndexes) contains(index int) bool {
	if _, ok := s[index]; ok {
		return true
	}
	return false
}

//nil indexes means no value is repeated
func populateSequentialList(list []int, n int, indexes skipIndexes) {
	repeated := 0
	odd := n % 2
	midPoint := n / 2
	for i := -midPoint; i < midPoint+odd; i++ {
		idx := i + midPoint
		if indexes != nil && idx != 0 && indexes.contains(idx) {
			list[i+midPoint] = i + repeated - 1
			repeated -= 1
			continue
		}

		list[i+midPoint] = i + repeated
	}
}
