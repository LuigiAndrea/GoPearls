// +build all column2 atleast twice

package twice

import (
	"testing"
)

func TestAtLeastTwice(t *testing.T) {
	n := 20
	list := make([]int, n)
	populateSequentialList(list, n, map[int]struct{}{2: {}, 7: {}, 13: {}})

	if ok, v := searchAtLeastTwiceValue(list); ok {
		if v != -9 {
			t.Errorf("\nExpected value: '%d' - Actual value '%d' \nList: %#v", -10, v, list)
		} else {
			t.Logf("Repeated element '%d' \nList %#v", v, list)
		}
	} else {
		t.Errorf("Repeated element not detected \nList: %#v", list)
	}

	populateSequentialList(list, n, nil)

	if ok, _ := searchAtLeastTwiceValue(list); ok {
		t.Errorf("No element appears at least twice in the list \nList: %#v", list)
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
	j := 0
	odd := n % 2
	midPoint := n / 2
	for i := -midPoint; i < midPoint+odd; i++ {
		idx := i + midPoint
		if indexes != nil && idx != 0 && indexes.contains(idx) {
			list[i+midPoint] = i + j - 1
			j -= 1
			continue
		}

		list[i+midPoint] = i + j
	}
}
