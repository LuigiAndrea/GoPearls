//Package setbit count number of 1s in binary representation of an integer.
package setbit

import (
	"bufio"
	"index/suffixarray"
	"os"
	"strconv"
)

func countSetBitsSequence(countFunc func(uint64) int) int {

	var count int
	file, err := os.Open("./bitSequence.data")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(splitAt64)
	for scanner.Scan() {
		u64, err := strconv.ParseUint(scanner.Text(), 2, 64)
		if err != nil {
			panic(err)
		}
		count += countFunc(u64)
	}

	return count
}

func splitAt64(data []byte, atEOF bool) (advance int, token []byte, err error) {

	dataLen := len(data)

	// If we're at EOF
	if atEOF {
		if dataLen == 0 { //no data passed
			return 0, nil, nil
		}
		return dataLen, data, nil
	}

	if dataLen > 64 {
		return 64, data[0:64], nil
	}

	return 0, nil, nil
}

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
