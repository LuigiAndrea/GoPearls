// +build all column1 kintegers

package kintegers

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestCreateFileWithKIntegers(t *testing.T) {
	if err := CreateFileWithRandomIntegers(
		MinMaxInterval{Min: 1000000, Max: 1000010},
		MinMaxInterval{Min: 1000050, Max: 1000054},
		MinMaxInterval{Min: 1000500, Max: 2000000}); err != nil {
		t.Errorf("\nError during creation of the file %v", err)
	} else {
		t.Logf("File '%s' succesfully created", Filename)
	}

	file, err := os.Open(Filename)
	if err != nil {
		t.Errorf("Unable to open file %s: %s", Filename, err)
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
				t.Errorf("Error reading %s: %s", Filename, err)
			}
		}
	}
	expectedValue := 999514
	if numbersInFile != expectedValue {
		t.Errorf("\nExpected value %d - Actual value %d", expectedValue, numbersInFile)
	}

	if err := os.Remove(Filename); err != nil {
		t.Logf("Error deleting the file")
	}
}

func TestCreateFileWithKIntegersWrongValues(t *testing.T) {

	valuesTest := []MinMaxInterval{{Min: -3, Max: 4}, {Min: 2, Max: -8}, {Min: 100, Max: 99}}

	for i, v := range valuesTest {
		if err := CreateFileWithRandomIntegers(MinMaxInterval{v.Min, v.Max}); err == nil {
			t.Error("Expected an parameter error")
		} else {
			t.Logf("%d: %s", i, err)
		}
	}
}
