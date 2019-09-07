package utilities

import (
	"testing"
)

//CheckArraySameValues Check if two arrays have the same values and in the same order
func CheckArraySameValues(t *testing.T, arrays CheckArrays) {
	if !arrays.SameLength() {
		t.Errorf("\nDifferent number of elements between the two arrays")
		return
	}

	for i := 0; i < arrays.Size(); i++ {
		if v := arrays.AreEqual(i); !v {
			arrays.PrintError(t, i)
		}
	}
}

//CheckArrays interface is used for comparing two arrays in testing
type CheckArrays interface {
	SameLength() bool
	AreEqual(i int) bool
	Size() int
	PrintError(t *testing.T, i int)
}

// StringArrays attaches the methods of CheckArrays to struct StringArray
type StringArrays struct {
	Expected []string
	Actual   []string
}

// SameLength checks if the two arrays have the same length
func (s StringArrays) SameLength() bool { return len(s.Expected) == len(s.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (s StringArrays) AreEqual(i int) bool { return s.Expected[i] == s.Actual[i] }

// Size is the length of StringArrays
func (s StringArrays) Size() int { return len(s.Expected) }

// PrintError displays an error message when the values at position i are different
func (s StringArrays) PrintError(t *testing.T, i int) {
	t.Errorf("\nExpected '%s' - Actual '%s'", s.Expected[i], s.Actual[i])
}

// IntArrays attaches the methods of CheckArrays to struct IntArrays
type IntArrays struct {
	Expected []int
	Actual   []int
}

// SameLength checks if the two arrays have the same length
func (s IntArrays) SameLength() bool { return len(s.Expected) == len(s.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (s IntArrays) AreEqual(i int) bool { return s.Expected[i] == s.Actual[i] }

// Size is the length of IntArrays struct
func (s IntArrays) Size() int { return len(s.Expected) }

// PrintError displays an error message when the values at position i are different
func (s IntArrays) PrintError(t *testing.T, i int) {
	t.Errorf("\nExpected '%d' - Actual '%d'", s.Expected[i], s.Actual[i])
}

//Float64Arrays attaches the methods of CheckArrays to struct Float64Arrays
type Float64Arrays struct {
	Expected []float64
	Actual   []float64
}

// SameLength checks if the two arrays have the same length
func (s Float64Arrays) SameLength() bool { return len(s.Expected) == len(s.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (s Float64Arrays) AreEqual(i int) bool { return s.Expected[i] == s.Actual[i] }

// Size is the length of Float64Arrays struct
func (s Float64Arrays) Size() int { return len(s.Expected) }

// PrintError displays an error message when the values at position i are different
func (s Float64Arrays) PrintError(t *testing.T, i int) {
	t.Errorf("\nExpected '%0.3f' - Actual '%0.3f'", s.Expected[i], s.Actual[i])
}
