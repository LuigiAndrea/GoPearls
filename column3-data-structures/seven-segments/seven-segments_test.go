// +build all column3 segments sevenSegments

package segments

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	number        uint16
	expectedValue []byte
}

var runPrintNumber = true

func TestBuildSegmentsAndEncoding(t *testing.T) {
	tests := []byte{125, 80, 55, 87, 90, 79, 111, 84, 127, 94}
	buildSegments()
	for i := 0; i <= 9; i++ {
		value := encodeDigit(uint16(i))
		if err := a.AssertDeepEqual(tests[i], value); err != nil {
			t.Error(err)
			runPrintNumber = false
		}
	}
}

func TestDispalyNumbers(t *testing.T) {
	tests := []testData{
		testData{number: 65535, expectedValue: []byte{111, 79, 79, 87, 79}},
		testData{number: 0, expectedValue: []byte{125}},
		testData{number: 984, expectedValue: []byte{94, 127, 90}},
	}

	for i, test := range tests {
		result := displayNumber(test.number)
		if err := a.AssertSlicesEqual(a.ByteSlicesMatch{Expected: test.expectedValue, Actual: result}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
			runPrintNumber = false
		}
	}
}

func TestPrintNumber(t *testing.T) {
	if runPrintNumber {
		printDisplay(50948)
		printDisplay(61723)
	}
}
