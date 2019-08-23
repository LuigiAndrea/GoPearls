package sortfile

import (
	bv "GoPearls/column1-oyster/2-bit-vectors"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const bitVectorSize = 20000
const filenameResult = "./kIntegersSorted.data"

func sortFile(path string) error {
	bitVector, err := createBitVectorFromFile(path)
	if err != nil {
		return err
	}

	if err := createSortedFile(bitVector); err != nil {
		return err
	}

	return nil
}

func createBitVectorFromFile(path string) (bv.BitVector, error) {
	fileRandomIntegers, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file '%s': %s", path, err)
	}
	defer fileRandomIntegers.Close()

	bitVector, err := bv.NewBitVector(bitVectorSize)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	reader := bufio.NewReader(fileRandomIntegers)
	for {
		if value, err := reader.ReadString(' '); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, fmt.Errorf("Error reading '%s': %s", path, err)
			}
		} else {
			intValue, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				return nil, errors.New(err.Error())
			}

			bitVector.Set(intValue)
		}
	}
	return *bitVector, nil
}

func createSortedFile(bitVector bv.BitVector) error {
	file, err := os.Create(filenameResult)
	if err != nil {
		return fmt.Errorf("Unable to create file '%s': %s", filenameResult, err)
	}

	writer := bufio.NewWriter(file)
	defer file.Close()

	for i := 0; i < bitVectorSize; i++ {
		if r, err := bitVector.Get(i); err != nil {
			log.Fatal(err)
		} else if r == 1 {
			writer.WriteString(fmt.Sprintf("%d ", i))
		}
	}

	writer.Flush()
	return nil
}
