// +build all column1 sortFile

package sortfile

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	kintegers "github.com/LuigiAndrea/GoPearls/column1-oyster/generate-k-random-integer"
	a "github.com/LuigiAndrea/test-helper/assertions"
)

const Filename = "./kIntegers.data"
const FilenameResult = "./kIntegersSorted.data"

func TestSortFile(t *testing.T) {

	if err := a.AssertException(nil, kintegers.CreateFileWithRandomIntegers(Filename,
		kintegers.MinMaxInterval{Min: 200, Max: 300}, kintegers.MinMaxInterval{Min: 500, Max: 1200})); err != nil {
		t.Error(err)
	}

	file, e := SortFile(Filename, FilenameResult)
	if err := a.AssertException(nil, e); err != nil {
		t.Error(e)
	}

	defer cleanWorkSpace(t, Filename, FilenameResult)
	defer file.Close()

	var expectedEOFIndex = 800
	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		if v, err := reader.ReadString(' '); err != nil {
			if err == io.EOF {
				if err := a.AssertDeepEqual(i, expectedEOFIndex); err != nil {
					t.Errorf("EOF: %v", err)
				}
				break
			}
		} else {
			assertValueAtIndex(t, i, v)
		}
	}
}

func TestNotFoundFile(t *testing.T) {
	_, e := SortFile("NotExistingFilename", FilenameResult)

	if err := a.AssertException(&os.PathError{}, e); err != nil {
		t.Error(err)
	}
}

func TestWrongFileName(t *testing.T) {

	if err := a.AssertException(nil, kintegers.CreateFileWithRandomIntegers(Filename,
		kintegers.MinMaxInterval{Min: 200, Max: 300}, kintegers.MinMaxInterval{Min: 500, Max: 1200})); err != nil {
		t.Error(err)
	}

	_, e := SortFile(Filename, "///")
	if err := a.AssertException(&os.PathError{}, e); err != nil {
		t.Error(err)
	}

	cleanWorkSpace(t, Filename)
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

	_, e := SortFile(wrongDataFile, FilenameResult)
	if err := a.AssertException(&strconv.NumError{}, e); err != nil {
		t.Error(err)
	}

	cleanWorkSpace(t, wrongDataFile)
}

func cleanWorkSpace(t *testing.T, path ...string) {
	for _, p := range path {
		if err := os.Remove(p); err != nil {
			log.Fatal(err)
		}
	}
}

func assertValueAtIndex(t *testing.T, idx int, v string) {
	switch idx {
	case 0:
		assertValue(t, "200", v)
	case 99:
		assertValue(t, "299", v)
	case 100:
		assertValue(t, "500", v)
	case 300:
		assertValue(t, "700", v)
	case 799:
		assertValue(t, "1199", v)
	}
}

func assertValue(t *testing.T, value1, value2 string) {
	if err := a.AssertDeepEqual(value1, strings.TrimSpace(value2)); err != nil {
		t.Error(err)
	}
}
