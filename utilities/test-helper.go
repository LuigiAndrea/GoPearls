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
		if v := arrays.Equals(i); !v {
			arrays.PrintError(t, i)
		}
	}
}

//CheckArrays interface defines method for CheckArraySameValues
type CheckArrays interface {
	SameLength() bool
	Equals(i int) bool
	Size() int
	PrintError(t *testing.T, i int)
}

// StringArrays attaches the methods of CheckArrays to []string
type StringArrays struct {
	Expected []string
	Actual   []string
}

// SameLength checks if the two arrays have the same length
func (s StringArrays) SameLength() bool { return len(s.Expected) == len(s.Actual) }

// Equals checks if the two arrays have the same value at position i
func (s StringArrays) Equals(i int) bool { return s.Expected[i] == s.Actual[i] }

// Size is the length of StringArrays
func (s StringArrays) Size() int { return len(s.Expected) }

// PrintError displays an error message when the values at position i are different
func (s StringArrays) PrintError(t *testing.T, i int) {
	t.Errorf("\nExpected '%s' - Actual '%s'", s.Expected[i], s.Actual[i])
}
