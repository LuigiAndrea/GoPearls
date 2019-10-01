package utilities

import "sort"

//Reverse a slice
func Reverse(data sort.Interface, i, j int) {
	for ; j > i; i++ {
		data.Swap(i, j)
		j--
	}
}

//ByteSlice attaches the methods of Interface to []byte, sorting in increasing order.
type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
