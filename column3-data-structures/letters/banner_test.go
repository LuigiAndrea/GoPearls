// +build all column3 banner letters

package banner

import (
	"fmt"
	"testing"
)

func TestBanner(t *testing.T) {
	type testData struct {
		in  string
		out string
	}

	tests := []testData{
		testData{in: "2b5*3b1*", out: "  *****   *"},
		testData{in: "5*2b", out: "*****  "},
		testData{in: "8*", out: "********"},
	}

	for _, test := range tests {
		if res := decodeString(test.in); res != test.out {
			t.Errorf("Expected value '%s' - Actual value '%s'", test.out, res)
		}
	}

	decodedLetter := decodeLetter(letters['A'])
	if bannerStr := Banner('A'); bannerStr != decodedLetter {
		t.Errorf("Expected value '%s' - Actual value '%s'", decodedLetter, bannerStr)
	} else {
		fmt.Println(bannerStr)
	}

}

func TestBannerNotSupportedCharacter(t *testing.T) {
	if res := Banner('$'); len(res) > 0 {
		t.Errorf("Expected an empty string - Actual value '%s'", res)
	}
}
