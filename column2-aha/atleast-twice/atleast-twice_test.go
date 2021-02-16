// +build all column2 atleast twice

package twice

import (
	"fmt"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	lengthList    int
	indexesToSkip skipIndexes
	expectedValue int
}

type skipIndexes map[int]struct{}

func (s skipIndexes) contains(index int) bool {
	if _, ok := s[index]; ok {
		return true
	}
	return false
}

func TestAtLeastTwice(t *testing.T) {

	tests := []testData{
		{lengthList: 20, indexesToSkip: indexesToSkip(2, 3, 7, 13), expectedValue: -9},
		{lengthList: 10, indexesToSkip: indexesToSkip(7), expectedValue: 1},
		{lengthList: 2, indexesToSkip: indexesToSkip(1), expectedValue: -1},
	}

	for i, test := range tests {

		list := make([]int, test.lengthList)
		populateSequentialList(list, test.lengthList, test.indexesToSkip)

		if ok, v := searchAtLeastTwiceValue(list); ok {
			if err := a.AssertDeepEqual(test.expectedValue, v); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, err))
			}
		} else {
			t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("Repeated element not detected \nList: %#v", list)))
		}
	}
}

func TestEdgeCases(t *testing.T) {

	tests := []testData{
		{lengthList: 5, indexesToSkip: nil, expectedValue: 5},
		{lengthList: 0, indexesToSkip: nil, expectedValue: 0},
	}

	for i, test := range tests {
		list := make([]int, test.lengthList)

		populateSequentialList(list, test.lengthList, nil)

		if ok, v := searchAtLeastTwiceValue(list); ok {
			t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("No element appears at least twice in the list \nList: %#v", list)))
		} else if err := a.AssertDeepEqual(test.expectedValue, v); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestAtLeastTwiceGeneral(t *testing.T) {

	type testData struct {
		list          []int
		expectedValue int
	}

	tests := []testData{
		{list: []int{0, 0, 1, 2, 3, 4, 5}, expectedValue: 0},
		{list: []int{-1, 0, 1, 1, 1, 2, 2, 2, 3, 4, 5, 5}, expectedValue: 1},
		{list: []int{2, 2, 2}, expectedValue: 2},
		{list: []int{10, 11, 12, 13, 13}, expectedValue: 13},
	}

	for i, test := range tests {

		if ok, v := searchAtLeastTwiceValue(test.list); ok {
			if err := a.AssertDeepEqual(test.expectedValue, v); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, err))
			}
		} else {
			t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("Repeated elements not detected \nList: %#v", test.list)))
		}
	}
}

//build a list from scratch skipping the indexes provided as parameter, nil indexes means no value is repeated
func populateSequentialList(list []int, n int, indexes skipIndexes) {
	repeated := 0
	odd := n % 2
	midPoint := n / 2
	for i := -midPoint; i < midPoint+odd; i++ {
		idx := i + midPoint
		if indexes != nil && idx != 0 && indexes.contains(idx) {
			list[i+midPoint] = i + repeated - 1
			repeated--
			continue
		}

		list[i+midPoint] = i + repeated
	}
}

//simplify test initialization
func indexesToSkip(index ...int) skipIndexes {
	result := make(skipIndexes, len(index))
	for _, v := range index {
		result[v] = struct{}{}
	}
	return result
}
