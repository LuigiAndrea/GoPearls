package column1

import (
	bv "GoPearls/column1-oyster/2-bit-vectors"
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

func sortFile(path string) {

	fileRandomIntegers, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file %s: %s", path, err)
		return
	}
	defer fileRandomIntegers.Close()

	bitVector, err := bv.NewBitVector(bitVectorSize)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(fileRandomIntegers)
	for {
		if value, err := reader.ReadString(' '); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Error reading %s: %s", path, err)
			}
		} else {
			if intValue, err := strconv.Atoi(strings.TrimSpace(value)); err != nil {
				log.Fatal(err)
			} else {
				bitVector.Set(intValue)
			}
		}
	}

	fileSorted, err := os.Create(filenameResult)
	if err != nil {
		log.Fatalf("Unable to create file '%s': %s", filenameResult, err)
	}

	writer := bufio.NewWriter(fileSorted)

	defer fileSorted.Close()

	for i := 0; i < bitVectorSize; i++ {
		if r, err := bitVector.Get(i); err != nil {
			log.Fatal(err)
		} else if r == 1 {
			writer.WriteString(fmt.Sprintf("%d ", i))
		}
	}

	writer.Flush()
}
