// +build all column1 kintegers

package kintegers

import (
	"bufio"
	"errors"
	"io"
	"os"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

const Filename = "./kIntegers.data"

func TestCreateFileWithKIntegers(t *testing.T) {
	valuesTest := []MinMaxInterval{{Min: 1000000, Max: 1000010}, {Min: 1000050, Max: 1000054}, {Min: 1000500, Max: 2000000}}

	if err := a.AssertException(nil, CreateFileWithRandomIntegers(Filename, valuesTest...)); err != nil {
		t.Error(err)
	}

	file, e := os.Open(Filename)
	if err := a.AssertException(nil, e); err != nil {
		t.Error(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var numbersInFile = 0
	for ; ; numbersInFile++ {
		if _, er := reader.ReadString(' '); er != nil {
			if er == io.EOF {
				break
			} else if err := a.AssertException(nil, er); err != nil {
				t.Error(err)
			}
		}
	}
	expectedValue := 999514

	if err := a.AssertDeepEqual(expectedValue, numbersInFile); err != nil {
		t.Error(err)
	}

	if err := a.AssertException(nil, os.Remove(Filename)); err != nil {
		t.Error(err)
	}
}

func TestCreateFileWithKIntegersWrongValues(t *testing.T) {
	valuesTest := []MinMaxInterval{{Min: -3, Max: 4}, {Min: 2, Max: -8}, {Min: 100, Max: 99}}

	for i, v := range valuesTest {
		if err := a.AssertException(errors.New(""),
			CreateFileWithRandomIntegers(Filename, MinMaxInterval{v.Min, v.Max})); err != nil {
			t.Error(m.ErrorMessageTestCount(i, err))
		}
	}
}

func TestCreateFileWithKIntegersWrongFilename(t *testing.T) {
	if err := a.AssertException(errors.New(""), CreateFileWithRandomIntegers("///",
		MinMaxInterval{Min: 1000000, Max: 1000010}, MinMaxInterval{Min: 1000500, Max: 2000000})); err != nil {
		t.Error(err)
	}
}
