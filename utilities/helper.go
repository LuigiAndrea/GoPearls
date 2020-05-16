package utilities

import (
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

//Max calculate the max element in the parameters list
func Max(elements ...float64) float64 {
	maxElement := 0.0
	for _, el := range elements {
		if el > maxElement {
			maxElement = el
		}
	}
	return maxElement
}

//Round a number to n decimal places
func Round(number float64, decimalPlaces int) float64 {
	normalize := math.Pow10(decimalPlaces)
	return math.Round(number*normalize) / normalize
}
