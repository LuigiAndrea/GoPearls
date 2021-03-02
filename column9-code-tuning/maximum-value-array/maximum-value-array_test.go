// +build all column9 maximumvalue maxarray maxdefault

package maximumvalue

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestMaximumValueArray(t *testing.T) {
	for i, test := range tests {
		max, _ := MaxInArray(test.array)
		if err := a.AssertDeepEqual(test.maxValue, max); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestMaximumValueArrayExceptions(t *testing.T) {
	for i, test := range testsEx {
		_, e := MaxInArray(test.array)

		if err := a.AssertException(errors.New(""), e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
