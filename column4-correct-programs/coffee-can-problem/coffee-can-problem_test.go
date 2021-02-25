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

var funcToRun = []func([]bean) (bean, error){coffeCanRemainingBean, coffeeCanBeans}

func TestCoffeeBean(t *testing.T) {
	tests := []testData{
		{beans: []bean{{color: white}, {color: white}, {color: white}}, lastBeanColor: white},
		{beans: []bean{{color: white}, {color: white}, {color: black}}, lastBeanColor: black},
		{beans: []bean{{color: white}, {color: black}, {color: black}}, lastBeanColor: white},
		{beans: []bean{{color: white}, {color: black}, {color: black}, {color: black}}, lastBeanColor: white},
		{beans: []bean{{color: black}, {color: white}, {color: black}, {color: black}, {color: black}}, lastBeanColor: white},
		{beans: []bean{{color: black}, {color: black}}, lastBeanColor: black},
		{beans: []bean{{color: black}}, lastBeanColor: black},
		{beans: []bean{{color: white}}, lastBeanColor: white},
	}

	for _, f := range funcToRun {
		for i, test := range tests {
			t.Run(m.GetFuncName(f), func(t *testing.T) {
				bean, _ := f(test.beans)
				if err := a.AssertDeepEqual(test.lastBeanColor.String(), bean.color.String()); err != nil {
					t.Error(m.ErrorMessageTestCount(i+1, err))
				}
			})
		}
	}
}

func TestCoffeeBeanWrongColor(t *testing.T) {
	tests := []testData{
		{beans: []bean{{color: 3}, {color: white}, {color: white}}},
		{beans: []bean{{color: -3}, {color: black}, {color: white}}},
		{beans: []bean{}},
		{beans: nil},
	}

	for _, f := range funcToRun {
		for i, test := range tests {
			t.Run(m.GetFuncName(f), func(t *testing.T) {
				_, e := f(test.beans)
				if err := a.AssertException(errors.New(""), e); err != nil {
					t.Error(m.ErrorMessageTestCount(i+1, err))
				}
			})
		}
	}
}
