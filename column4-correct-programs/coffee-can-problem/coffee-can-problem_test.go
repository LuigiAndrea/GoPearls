// +build all column4 coffee bean

package coffeebean

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
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
		testData{beans: []bean{bean{color: black}}, lastBeanColor: black},
	}

	for i, test := range tests {
		bean, _ := coffeeCanBeans(test.beans)
		if err := a.AssertDeepEqual(test.lastBeanColor.String(), bean.color.String()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestCoffeeBeanWrongColor(t *testing.T) {
	tests := []testData{
		testData{beans: []bean{bean{color: 3}, bean{color: white}, bean{color: white}}},
		testData{beans: []bean{}},
	}

	for i, test := range tests {
		_, e := coffeeCanBeans(test.beans)
		if err := a.AssertException(errors.New(""), e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
