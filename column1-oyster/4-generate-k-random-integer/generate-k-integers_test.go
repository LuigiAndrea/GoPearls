// +build all column1 kintegers

package column1

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestCreateFileWithKIntegers(t *testing.T) {
	if err := createFileWithRandomIntegers(
		minmaxInterval{min: 1000000, max: 1000010},
		minmaxInterval{min: 1000050, max: 1000054},
		minmaxInterval{min: 1000500, max: 2000000}); err != nil {
		t.Errorf("\nError during creation of the file %v", err)
	} else {
		t.Logf("File '%s' succesfully created", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("Unable to open file %s: %s", filename, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var numbersInFile = 0
	for ; ; numbersInFile++ {
		if _, err := reader.ReadString(' '); err != nil {
			if err == io.EOF {
				break
			} else {
				t.Errorf("Error reading %s: %s", filename, err)
			}
		}
	}
	expectedValue := 999514
	if numbersInFile != expectedValue {
		t.Errorf("\nExpected value %d - Actual value %d", expectedValue, numbersInFile)
	}

	if err := os.Remove(filename); err != nil {
		t.Logf("Error deleting the file")
	}
}

func TestCreateFileWithKIntegersWrongValues(t *testing.T) {

	valuesTest := []minmaxInterval{{min: -3, max: 4}, {min: 2, max: -8}, {min: 100, max: 99}}

	for i, v := range valuesTest {
		if err := createFileWithRandomIntegers(minmaxInterval{v.min, v.max}); err == nil {
			t.Error("Expected an parameter error")
		} else {
			t.Logf("%d: %s", i, err)
		}
	}
}
