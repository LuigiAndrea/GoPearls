// +build all column1 kintegers

package column1

import (
	"os"
	"testing"
)

func TestCreateFileWithKIntegers(t *testing.T) {
	if err := createFileWithRandomIntegers(1000000, 2000000); err != nil {
		t.Errorf("\nError during creation of the file %v", err)
	} else {
		t.Logf("File succesfully created")
	}

	if err := os.Remove("./kIntegers.data"); err != nil {
		t.Logf("Error deleting the file")
	}
}

func TestCreateFileWithKIntegersWrongValues(t *testing.T) {
	type minmaxInterval struct {
		min int
		max int
	}

	valuesTest := []minmaxInterval{{min: -3, max: 4}, {min: 2, max: -8}, {min: 100, max: 99}}

	for i, v := range valuesTest {
		if err := createFileWithRandomIntegers(v.min, v.max); err == nil {
			t.Error("Expected an parameter error")
		} else {
			t.Logf("%d: %s", i, err)
		}
	}
}
