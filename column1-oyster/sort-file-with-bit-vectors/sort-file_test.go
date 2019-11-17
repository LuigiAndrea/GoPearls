// +build all column1 sortFile

package sortfile

import (
	kintegers "github.com/LuigiAndrea/GoPearls/column1-oyster/generate-k-random-integer"
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

	file, err := SortFile(kintegers.Filename)
	if err != nil {
		log.Fatal(err)
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
	cleanWorkSpace(t, kintegers.Filename)
	cleanWorkSpace(t, filenameResult)
}

func TestWrongFilename(t *testing.T) {
	if _, err := SortFile("NotExistingFilename"); err == nil {
		t.Errorf("Expected an error: Not Existing Filename")
	} else {
		t.Log(err)
	}
}

func TestWrongDataInFile(t *testing.T) {
	wrongDataFile := "./wrongDataFile"
	file, err := os.Create(wrongDataFile)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("Wrong data")
	writer.Flush()
	file.Close()

	if _, err := SortFile(wrongDataFile); err == nil {
		t.Errorf("Expected an error: Wrong Data in file")
	} else {
		t.Log(err)
	}

	cleanWorkSpace(t, wrongDataFile)
}

func cleanWorkSpace(t *testing.T, path string) {
	if err := os.Remove(path); err != nil {
		t.Logf("Error deleting the file '%s'", path)
	}
}
