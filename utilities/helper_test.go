// +build all utilities helper

package utilities

import (
	"testing"

	goth "github.com/LuigiAndrea/test-helper"
)

type testData struct {
	expectedValues []interface{}
	newElements    []interface{}
	testValues     []interface{}
}

func TestPreAppend(t *testing.T) {

	tests := []testData{
		testData{testValues: []interface{}{1, 2, -6, 111},
			expectedValues: []interface{}{12, 3, 1, 2, -6, 111},
			newElements:    []interface{}{3, 12}},
		testData{testValues: []interface{}{"Ciao", "Hello", "Car"},
			expectedValues: []interface{}{"NewCiao", "Ciao", "Hello", "Car"},
			newElements:    []interface{}{"NewCiao"}},
		testData{testValues: []interface{}{"Ciao", "Hello", "Car"},
			expectedValues: []interface{}{123, "Ciao", "Hello", "Car"},
			newElements:    []interface{}{123}},
	}

	for _, test := range tests {

		data := PreAppend(test.testValues, test.newElements...)
		if err := goth.AssertArraysEqual(goth.DataArrays{Expected: test.expectedValues, Actual: data}); err != nil {
			t.Errorf(err.Error())
		}
	}
}
