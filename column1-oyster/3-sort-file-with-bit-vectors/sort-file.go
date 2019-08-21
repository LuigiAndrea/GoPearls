package column1

import (
	bv "GoPearls/column1-oyster/2-bit-vectors"
	kintegers "GoPearls/column1-oyster/4-generate-k-random-integer"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const bitVectorSize = 20000
const filenameResult = "./kIntegersSorted.data"

func sortFile() {
	if err := kintegers.CreateFileWithRandomIntegers(
		kintegers.MinMaxInterval{Min: 200, Max: 300},
		kintegers.MinMaxInterval{Min: 500, Max: 1200}); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(kintegers.Filename)
	if err != nil {
		log.Fatalf("Unable to open file %s: %s", kintegers.Filename, err)
		return
	}
	defer file.Close()

	bitVector, err := bv.NewBitVector(bitVectorSize)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	for {
		if value, err := reader.ReadString(' '); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Error reading %s: %s", kintegers.Filename, err)
			}
		} else {
			if intValue, err := strconv.Atoi(strings.TrimSpace(value)); err != nil {
				log.Fatal(err)
			} else {
				bitVector.Set(intValue)
			}
		}
	}

	file, err = os.Create(filenameResult)
	if err != nil {
		log.Fatalf("Unable to create file '%s': %s", filenameResult, err)
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
}
