// +build all column4 coffee bean

package coffeebean

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
)

type testData struct {
	beans         []bean
	lastBeanColor string
}

func TestCoffeeBean(t *testing.T) {
	tests := []testData{
		testData{beans: []bean{bean{color: "White"}, bean{color: "White"}, bean{color: "White"}}, lastBeanColor: "White"},
		testData{beans: []bean{bean{color: "White"}, bean{color: "White"}, bean{color: "Black"}}, lastBeanColor: "Black"},
		testData{beans: []bean{bean{color: "White"}, bean{color: "Black"}, bean{color: "Black"}}, lastBeanColor: "White"},
		testData{beans: []bean{bean{color: "White"}, bean{color: "Black"}, bean{color: "Black"}, bean{color: "Black"}}, lastBeanColor: "White"},
		testData{beans: []bean{bean{color: "Black"}, bean{color: "White"}, bean{color: "Black"}, bean{color: "Black"}, bean{color: "Black"}}, lastBeanColor: "White"},
	}

	for _, test := range tests {
		if err := a.AssertDeepEqual(coffeeCanBeans(test.beans).color, test.lastBeanColor); err != nil {
			t.Error(err)
		}
	}
}
