// +build all utilities helper

package utilities

import (
	"testing"

	goth "github.com/LuigiAndrea/test-helper"
)

type testData struct {
	expectedValues []goth.Data
	newElements    []goth.Data
	testValues     []goth.Data
}

func TestPreAppend(t *testing.T) {

	tests := []testData{
		testData{testValues: []goth.Data{1, 2, -6, 111},
			expectedValues: []goth.Data{12, 3, 1, 2, -6, 111},
			newElements:    []goth.Data{3, 12}},
		testData{testValues: []goth.Data{"Ciao", "Hello", "Car"},
			expectedValues: []goth.Data{"NewCiao", "Ciao", "Hello", "Car"},
			newElements:    []goth.Data{"NewCiao"}},
		testData{testValues: []goth.Data{"Ciao", "Hello", "Car"},
			expectedValues: []goth.Data{123, "Ciao", "Hello", "Car"},
			newElements:    []goth.Data{123}},
	}

	for _, test := range tests {

		data := PreAppend(test.testValues, test.newElements...)
		if err := goth.CheckArraySameValues(goth.DataArrays{Expected: test.expectedValues, Actual: data}); err != nil {
			t.Errorf(err.Error())
		}
	}
}
