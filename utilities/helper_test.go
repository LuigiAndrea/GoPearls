// +build all utilities helper

package utilities

import (
	"errors"
	"math"
	"reflect"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestPreAppend(t *testing.T) {
	type testData struct {
		expectedValues, newElements, testValues []interface{}
	}

	tests := []testData{
		{testValues: []interface{}{1, 2, -6, 111},
			expectedValues: []interface{}{12, 3, 1, 2, -6, 111},
			newElements:    []interface{}{3, 12}},
		{testValues: []interface{}{"Ciao", "Hello", "Car"},
			expectedValues: []interface{}{"NewCiao", "Ciao", "Hello", "Car"},
			newElements:    []interface{}{"NewCiao"}},
		{testValues: []interface{}{"Ciao", "Hello", "Car"},
			expectedValues: []interface{}{123, "Ciao", "Hello", "Car"},
			newElements:    []interface{}{123}},
	}

	for i, test := range tests {

		data := PreAppend(test.testValues, test.newElements...)
		if err := a.AssertSlicesEqual(a.DataSlicesMatch{Expected: test.expectedValues, Actual: data}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestReverse(t *testing.T) {
	type testData struct {
		in, out []byte
	}

	tests := []testData{
		{in: []byte{125, 55, 90, 127}, out: []byte{127, 90, 55, 125}},
		{in: []byte{}, out: []byte{}},
		{in: nil, out: nil},
	}

	for i, test := range tests {
		Reverse(ByteSlice(test.in), 0, len(test.in)-1)
		if err := a.AssertSlicesEqual(a.ByteSlicesMatch{Expected: test.out, Actual: test.in}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
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
		{in: []byte{125, 55, 90, 112}, lengthOut: 4, lessOut: false, swapLessIndexes: swapLessIndex{i: 0, j: 3}},
		{in: []byte{125, 55, 90}, lengthOut: 3, lessOut: true, swapLessIndexes: swapLessIndex{i: 1, j: 2}},
	}

	for i, test := range tests {

		byteSlice := ByteSlice(test.in)

		if err := a.AssertDeepEqual(test.lengthOut, byteSlice.Len()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}

		if err := a.AssertDeepEqual(test.lessOut, byteSlice.Less(test.swapLessIndexes.i, test.swapLessIndexes.j)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}

		byteSlice.Swap(test.swapLessIndexes.i, test.swapLessIndexes.j)
		if err := a.AssertSlicesEqual(a.ByteSlicesMatch{Expected: byteSlice, Actual: test.in}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestMax(t *testing.T) {
	type testData struct {
		in  []float64
		out float64
	}

	tests := []testData{
		{in: []float64{18.2, 12.4, 3.5}, out: 18.2},
		{in: []float64{}, out: math.Inf(-1)},
		{in: nil, out: math.Inf(-1)},
		{in: []float64{1.2, -12.4, 3.5}, out: 3.5},
		{in: []float64{3.75}, out: 3.75},
		{in: []float64{3.75, math.Inf(-1)}, out: 3.75},
		{in: []float64{3.75, math.Inf(1)}, out: math.Inf(5)},
		{in: []float64{-3.75, math.Inf(-1)}, out: -3.75},
		{in: []float64{math.Inf(1), math.Inf(-1)}, out: math.Inf(0)},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.out, Max(test.in...)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
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
		{in: 183.467, in2: 2, out: 183.47},
		{in: 183.467, in2: 5, out: 183.46700},
		{in: 146.7032, in2: 0, out: 147},
		{in: -12.455787, in2: 2, out: -12.46},
		{in: 1.79, in2: 308, out: 1.79},
	}

	for i, test := range tests {
		num, _ := Round(test.in, test.in2)
		if err := a.AssertDeepEqual(test.out, num); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestRoundEdgeCases(t *testing.T) {
	type testException struct {
		number        float64
		decimalPlaces int
		expected      error
	}

	errorString := "parameters too big"

	tests := []testException{
		{number: 1.80, decimalPlaces: 308, expected: &RoundError{Err: errorString}},
		{number: 180.43, decimalPlaces: 3108, expected: &RoundError{Err: errorString}},
		{number: 1.80, decimalPlaces: -3, expected: &RoundError{Err: "utilities.Round: decimalPlace parameter must be a positive number"}},
	}

	for i, test := range tests {
		_, e := Round(test.number, test.decimalPlaces)
		if err := a.AssertDeepException(test.expected, e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestRoundWrongException(t *testing.T) {
	type testException struct {
		number        float64
		decimalPlaces int
		expected      error
	}
	errorString := "utilities.Round: decimalPlace parameter must be a positive number"
	tests := []testException{
		{number: 1.80, decimalPlaces: -3, expected: errors.New(errorString)},
	}

	for _, test := range tests {
		_, e := Round(test.number, test.decimalPlaces)
		if err := a.AssertDeepException(test.expected, e); err == nil { //The expected exception should be RoundError
			t.Error("Expected exception!")
		} else {
			a.AssertDeepEqual(errorString, e.Error()) //check the message returned
		}
	}
}

//Just run the function and check the type of the function returned
func TestElapse(t *testing.T) {

	var funcExpected func()
	funcReturned := Elapse(m.GetFuncName(Elapse))

	//do some work
	for i := 0; i < 1000000; i++ {
		math.Cos(float64(i))
	}

	funcReturned()

	if err := a.AssertDeepEqual(reflect.TypeOf(funcExpected), reflect.TypeOf(funcReturned)); err != nil {
		t.Error(err.Error())
	}
}

func TestGetDistance(t *testing.T) {

	type testData struct {
		value1, value2, result float64
	}

	tests := []testData{
		{value1: 12, value2: -1, result: 13},
		{value1: 0, value2: -1, result: 1},
		{value1: 12, value2: 12, result: 0},
		{value1: 12, value2: math.Inf(0), result: math.Inf(0)},
		{value1: math.Inf(-1), value2: 12, result: math.Inf(0)},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(GetDistance(test.value1, test.value2), test.result); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
