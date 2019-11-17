package utilities

import (
	"sort"
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
