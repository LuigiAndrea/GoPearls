// +build all column3 hyphenations

package hyphenations

import (
	"testing"

	assert "github.com/LuigiAndrea/test-helper/assertions"
	message "github.com/LuigiAndrea/test-helper/messages"
)

func TestHyphenations(t *testing.T) {
	type testData struct {
		in  string
		out string
	}

	tests := []testData{
		testData{in: "ethnic", out: "h-nic"},
		testData{in: "realistic", out: "al-is-tic"},
		testData{in: "5*2b", out: ""},
		testData{in: "clinic", out: "n-ic"},
		testData{in: "alcoholic", out: "l-ic"},
		testData{in: "terrific", out: "f-ic"},
		testData{in: "public", out: "b-lic"},
		testData{in: "analytic", out: "-lyt-ic"},
		testData{in: "metallic", out: "l-lic"},
		testData{in: "bicyclic", out: "-clic"},
	}

	for _, test := range tests {
		res := hyphenation(test.in)
		if err := assert.AssertDeepEqual(res, test.out); err != nil {
			t.Errorf(message.ErrorMessage(test.out, res))
		}
	}
}
