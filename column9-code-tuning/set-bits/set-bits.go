//Write an efficient program to count number of 1s in binary representation of an integer.

package setbit

import (
	"index/suffixarray"
	"strconv"
)

func countSetBits(number uint64) int {
	sa := suffixarray.New([]byte(strconv.FormatUint(number, 2)))
	return len(sa.Lookup([]byte("1"), -1))
}

//Brian Kernighan’s Algorithm
func countSetBitsKernighan(number uint64) int {
	count := 0
	for number > 0 {
		number &= number - 1
		count++
	}
	return count
}

//Lookup table algorithms
var lookupBitsTable [256]byte

func init() {
	for i := range lookupBitsTable {
		lookupBitsTable[i] = lookupBitsTable[i/2] + byte(i&1)
	}
}

func countSetBitsLookup(number uint64) int {
	var sum byte
	const length = 8

	for i := 0; i < length; i++ {
		sum += lookupBitsTable[byte(number>>(i*length))]
	}

	return int(sum)

}

//Lookup table using Nibble (4 bits)
var lookupNibbleTable = [16]byte{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

func countSetBitsNibble(number uint64) int {
	var sum byte
	const nibblesize, maskNibble = 4, 0xf

	for i := 0; i < 16; i++ {
		sum += lookupNibbleTable[(number>>(i*nibblesize))&maskNibble]
	}

	return int(sum)

}
