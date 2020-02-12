// +build all column4 coffee bean

package coffeebean

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
)

type testData struct {
	beans         []bean
	lastBeanColor colorBean
}

func TestCoffeeBean(t *testing.T) {
	tests := []testData{
		testData{beans: []bean{bean{color: white}, bean{color: white}, bean{color: white}}, lastBeanColor: white},
		testData{beans: []bean{bean{color: white}, bean{color: white}, bean{color: black}}, lastBeanColor: black},
		testData{beans: []bean{bean{color: white}, bean{color: black}, bean{color: black}}, lastBeanColor: white},
		testData{beans: []bean{bean{color: white}, bean{color: black}, bean{color: black}, bean{color: black}}, lastBeanColor: white},
		testData{beans: []bean{bean{color: black}, bean{color: white}, bean{color: black}, bean{color: black}, bean{color: black}}, lastBeanColor: white},
		testData{beans: []bean{bean{color: black}, bean{color: black}}, lastBeanColor: black},
	}

	for _, test := range tests {
		bean, _ := coffeeCanBeans(test.beans)
		if err := a.AssertDeepEqual(test.lastBeanColor.String(), bean.color.String()); err != nil {
			t.Error(err)
		}
	}
}

func TestCoffeeBeanWrongColor(t *testing.T) {
	tests := []testData{
		testData{beans: []bean{bean{color: 3}, bean{color: white}, bean{color: white}}},
		testData{beans: []bean{bean{color: black}}},
		testData{beans: []bean{}},
	}

	for _, test := range tests {
		if _, err := coffeeCanBeans(test.beans); err == nil {
			t.Error("Expected an exception!")
		}
	}
}
