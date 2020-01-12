// +build all column3 segments sevenSegments

package segments

import (
	"testing"

	goth "github.com/LuigiAndrea/test-helper/assertions"
)

type testData struct {
	number        uint16
	expectedValue []byte
}

func TestBuildSegmentsAndEncoding(t *testing.T) {
	tests := []byte{125, 80, 55, 87, 90, 79, 111, 84, 127, 94}
	buildSegments()
	for i := 0; i <= 9; i++ {
		value := encodeDigit(uint16(i))
		if value != tests[i] {
			t.Errorf("Expected '%d' - Actual '%d'", tests[i], value)
		}
	}
}

func TestDispalyNumbers(t *testing.T) {
	tests := []testData{
		testData{number: 65535, expectedValue: []byte{111, 79, 79, 87, 79}},
		testData{number: 0, expectedValue: []byte{125}},
		testData{number: 984, expectedValue: []byte{94, 127, 90}},
	}

	for _, test := range tests {
		result := displayNumber(test.number)
		if err := goth.AssertSlicesEqual(goth.ByteSlicesMatch{Expected: test.expectedValue, Actual: result}); err != nil {
			t.Error(err.Error())
		}
	}
}

func TestPrintNumber(t *testing.T) {
	printDisplay(50948)
	printDisplay(61723)
}
