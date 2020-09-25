package sortfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	bv "github.com/LuigiAndrea/GoPearls/column1-oyster/bit-vectors"
)

//Capacity of BitVector
const bitVectorSize = 200000

//SortFile Sort a file with integers, return the sorted File
func SortFile(path string, sortedFilename string) (*os.File, error) {
	bitVector, err := createBitVectorFromFile(path)
	if err != nil {
		return nil, err
	}

	err = createSortedFile(bitVector, sortedFilename)
	if err != nil {
		return nil, err
	}

	return os.Open(sortedFilename)
}

func createBitVectorFromFile(path string) (bv.BitVector, error) {
	fileRandomIntegers, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileRandomIntegers.Close()

	bitVector, _ := bv.NewBitVector(bitVectorSize)

	reader := bufio.NewReader(fileRandomIntegers)
	for {
		if value, err := reader.ReadString(' '); err != nil {
			break
		} else {
			intValue, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				return nil, err
			}

			bitVector.Set(intValue)
		}
	}
	return *bitVector, nil
}

func createSortedFile(bitVector bv.BitVector, sortedFilename string) error {
	file, err := os.Create(sortedFilename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer file.Close()

	for i := 0; i < bitVectorSize; i++ {
		if r, _ := bitVector.Get(i); r == 1 {
			writer.WriteString(fmt.Sprintf("%d ", i))
		}
	}

	writer.Flush()

	return nil
}
