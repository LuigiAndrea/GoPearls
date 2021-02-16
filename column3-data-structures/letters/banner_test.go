// +build all column3 banner letters

package banner

import (
	"fmt"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

var printBanner = true

func TestBanner(t *testing.T) {
	type testData struct {
		in  string
		out string
	}

	tests := []testData{
		{in: "2b5*3b1*", out: "  *****   *"},
		{in: "5*2b", out: "*****  "},
		{in: "8*", out: "********"},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.out, decodeString(test.in)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
			printBanner = false
		}
	}

	if err := a.AssertDeepEqual(decodeLetter(letters['A']), Banner('A')); err != nil {
		t.Error(err)
		printBanner = false
	}
}

func TestBannerNotSupportedCharacter(t *testing.T) {
	if err := a.AssertDeepEqual(0, len(Banner('$'))); err != nil {
		t.Error(err)
		printBanner = false
	}
}

func TestPrintBanner(t *testing.T) {
	if printBanner {
		fmt.Println(Banner('E'))
	}
}
