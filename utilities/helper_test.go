// +build all utilities helper

package utilities

import (
	"testing"

	goth "github.com/LuigiAndrea/test-helper"
)

func TestPreAppend(t *testing.T) {
	type testData struct {
		expectedValues, newElements, testValues []interface{}
	}

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

func TestReverse(t *testing.T) {
	type testData struct {
		in, out []byte
	}

	tests := []testData{
		testData{in: []byte{125, 55, 90, 127}, out: []byte{127, 90, 55, 125}},
		testData{in: []byte{}, out: []byte{}},
		testData{in: nil, out: nil},
	}

	for _, test := range tests {
		Reverse(ByteSlice(test.in), 0, len(test.in)-1)
		if err := goth.AssertArraysEqual(goth.ByteArrays{Expected: test.out, Actual: test.in}); err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestByteSliceType(t *testing.T) {
	type swapLessIndex struct {
		i, j int
	}

	type testData struct {
		in, out         []byte
		lengthOut       int
		lessOut         bool
		swapLessIndexes swapLessIndex
	}

	tests := []testData{
		testData{in: []byte{125, 55, 90, 112}, out: []byte{112, 55, 90, 125},
			lengthOut: 4, lessOut: false, swapLessIndexes: swapLessIndex{i: 0, j: 3}},
		testData{in: []byte{125, 55, 90}, out: []byte{125, 90, 55},
			lengthOut: 3, lessOut: true, swapLessIndexes: swapLessIndex{i: 1, j: 2}},
	}

	for _, test := range tests {
		t.Logf("--> %v", test.in)
		byteSlice := ByteSlice(test.in)

		if byteSliceLength := byteSlice.Len(); byteSliceLength != test.lengthOut {
			t.Errorf("Expected '%d' - Actual '%d'", test.lengthOut, byteSliceLength)
		}

		if byteLess := byteSlice.Less(test.swapLessIndexes.i, test.swapLessIndexes.j); byteLess != test.lessOut {
			t.Errorf("Expected '%t' - Actual '%t'", test.lessOut, byteLess)
		}

		byteSlice.Swap(test.swapLessIndexes.i, test.swapLessIndexes.j)
		if err := goth.AssertArraysEqual(goth.ByteArrays{Expected: test.out, Actual: test.in}); err != nil {
			t.Errorf(err.Error())
		}
	}
}
