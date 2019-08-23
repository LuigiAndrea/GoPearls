// +build all column1 sortFile

package sortfile

import (
	kintegers "GoPearls/column1-oyster/generate-k-random-integer"
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSortFile(t *testing.T) {
	if err := kintegers.CreateFileWithRandomIntegers(
		kintegers.MinMaxInterval{Min: 200, Max: 300},
		kintegers.MinMaxInterval{Min: 500, Max: 1200}); err != nil {
		log.Fatal(err)
	}

	if err := sortFile(kintegers.Filename); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filenameResult)
	if err != nil {
		log.Fatalf("Unable to open file '%s': %s", filenameResult, err)
		return
	}
	defer file.Close()

	var expectedEOFIndex = 800
	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		if v, err := reader.ReadString(' '); err != nil {
			if err == io.EOF {
				if i != expectedEOFIndex {
					t.Errorf("Expected EOF at position %d - Actual value %d", expectedEOFIndex, i)
				}
				break
			}
		} else {
			var expectedValue string
			switch i {
			case 0:
				expectedValue = "200"
				if expectedValue != strings.TrimSpace(v) {
					t.Errorf("Expected value %s - Actual value %s", expectedValue, v)
				}
			case 99:
				expectedValue = "299"
				if expectedValue != strings.TrimSpace(v) {
					t.Errorf("Expected value %s - Actual value %s", expectedValue, v)
				}
			case 100:
				expectedValue = "500"
				if expectedValue != strings.TrimSpace(v) {
					t.Errorf("Expected value %s - Actual value %s", expectedValue, v)
				}
			case 300:
				expectedValue = "700"
				if expectedValue != strings.TrimSpace(v) {
					t.Errorf("Expected value %s - Actual value %s", expectedValue, v)
				}
			case 799:
				expectedValue = "1199"
				if expectedValue != strings.TrimSpace(v) {
					t.Errorf("Expected value %s - Actual value %s", expectedValue, v)
				}
			}
		}
	}
}
