// +build all utilities helper

package utilities

import (
	"testing"
)

type testData struct {
	expectedValue []Data
	newElements   []Data
	testValues    []Data
}

func TestPreAppendInt(t *testing.T) {

	tests := []testData{
		testData{testValues: []Data{1, 2, -6, 111},
			expectedValue: []Data{12, 3, 1, 2, -6, 111},
			newElements:   []Data{3, 12}},
		testData{testValues: []Data{"Ciao", "Hello", "Car"},
			expectedValue: []Data{"NewCiao", "Ciao", "Hello", "Car"},
			newElements:   []Data{"NewCiao"}},
		testData{testValues: []Data{"Ciao", "Hello", "Car"},
			expectedValue: []Data{123, "Ciao", "Hello", "Car"},
			newElements:   []Data{123}},
	}

	for _, test := range tests {
		data := PreAppend(test.testValues, test.newElements...)
		CheckArraySameValues(t, DataArrays{Expected: test.expectedValue, Actual: data})
	}
}
