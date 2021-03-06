// +build all column9 maximumvalue maxarray maxsentinel

package maximumvalue

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestMaximumValueSentinelArray(t *testing.T) {

	for i, test := range tests {
		x, _ := NewXWithSentinel(test.array)

		if err := a.AssertDeepEqual(test.maxValue, x.Max()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestMaximumValueSentinelArrayExceptions(t *testing.T) {
	for i, test := range testsEx {
		_, e := NewXWithSentinel(test.array)

		if err := a.AssertException(errors.New(""), e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
