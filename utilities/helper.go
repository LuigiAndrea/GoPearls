package utilities

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"
)

//Reverse a slice
func Reverse(data sort.Interface, i, j int) {
	for ; j > i; i++ {
		data.Swap(i, j)
		j--
	}
}

//ByteSlice attaches the methods of Interface to []byte, sorting in increasing order.
type ByteSlice []byte

func (b ByteSlice) Len() int           { return len(b) }
func (b ByteSlice) Less(i, j int) bool { return b[i] < b[j] }
func (b ByteSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

//PreAppend a generic value to a slice
func PreAppend(list []interface{}, elements ...interface{}) []interface{} {
	for _, element := range elements {
		list = append([]interface{}{element}, list...)
	}
	return list
}

//Elapse calculate the time elapsed
func Elapse(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

//Max calculate the max element in the parameters list, return -Inf if no parameters are passed
func Max(elements ...float64) float64 {
	maxElement := math.Inf(-1)
	for _, el := range elements {
		if el > maxElement {
			maxElement = el
		}
	}
	return maxElement
}

//Round a number to n decimal places
func Round(number float64, decimalPlaces int) (float64, error) {
	if decimalPlaces < 0 {
		return 0.0, &RoundError{Err: "utilities.Round: decimalPlace parameter must be a positive number"}
	}

	errorString := "parameters too big"
	normalize := math.Pow10(decimalPlaces)

	if math.IsInf(normalize, 0) {
		return normalize, &RoundError{Err: errorString}
	}
	numberToRound := number * normalize

	if math.IsInf(numberToRound, 0) {
		return numberToRound, &RoundError{Err: errorString}
	}

	return math.Round(numberToRound) / normalize, nil
}

//RoundError records an error for Round function
type RoundError struct {
	Err string
}

func (re *RoundError) Error() string {
	return re.Err
}

// Is Compare the values of RoundError
func (re *RoundError) Is(e error) bool {
	var err *RoundError
	if errors.As(e, &err) {
		return re.Err == err.Err
	}

	return false
}
