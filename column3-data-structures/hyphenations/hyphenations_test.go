// +build all column3 hyphenations

package hyphenations

import (
	"testing"

	assert "github.com/LuigiAndrea/test-helper/assertions"
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
		if err := assert.AssertDeepEqual(hyphenation(test.in), test.out); err != nil {
			t.Error(err)
		}
	}
}
