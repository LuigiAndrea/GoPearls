// +build all column1 kintegers

package column1

import (
	"os"
	"testing"
)

func TestCreateFileWithKIntegers(t *testing.T) {
	if err := createFileWithRandomIntegers(1000000, 2000000); err != nil {
		t.Errorf("\nError during creation of the file %v", err)
	}

	if err := os.Remove("./kIntegers.data"); err != nil {
		t.Logf("Error deleting the file")
	}
}
