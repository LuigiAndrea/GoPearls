// +build all utilities helper

package utilities

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
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
		if err := a.AssertSlicesEqual(a.DataSlicesMatch{Expected: test.expectedValues, Actual: data}); err != nil {
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
		if err := a.AssertSlicesEqual(a.ByteSlicesMatch{Expected: test.out, Actual: test.in}); err != nil {
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
		if err := a.AssertSlicesEqual(a.ByteSlicesMatch{Expected: test.out, Actual: test.in}); err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestMax(t *testing.T) {
	type testData struct {
		in  []float64
		out float64
	}

	tests := []testData{
		testData{in: []float64{18.2, 12.4, 3.5}, out: 18.2},
		testData{in: []float64{}, out: 0},
		testData{in: nil, out: 0},
		testData{in: []float64{1.2, -12.4, 3.5}, out: 3.5},
		testData{in: []float64{3.75}, out: 3.75},
	}

	for _, test := range tests {
		if err := a.AssertDeepEqual(test.out, Max(test.in...)); err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestRound(t *testing.T) {
	type testData struct {
		in  float64
		in2 int
		out float64
	}

	tests := []testData{
		testData{in: 183.467, in2: 2, out: 183.47},
		testData{in: 183.467, in2: 5, out: 183.46700},
		testData{in: 146.7032, in2: 0, out: 147},
		testData{in: -12.455787, in2: 2, out: -12.46},
		testData{in: 1.79, in2: 308, out: 1.79},
	}

	for _, test := range tests {
		num, _ := Round(test.in, test.in2)
		if err := a.AssertDeepEqual(test.out, num); err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestRoundEdgeCases(t *testing.T) {

	if _, err := Round(1.80, 308); err == nil {
		t.Errorf("Excepted an exception!")
	}

	if _, err := Round(180.43, 3108); err == nil {
		t.Errorf("Excepted an exception!")
	}

	if _, err := Round(1.80, -3); err == nil {
		t.Errorf("Excepted an exception!")
	}

	_, excRound := Round(1.80, -3)
	if err := a.AssertException(errors.New("Some exception returned"), excRound); err != nil {
		t.Error(err)
	}
}
