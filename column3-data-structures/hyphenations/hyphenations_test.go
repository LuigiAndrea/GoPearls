// +build all column3 hyphenations

package hyphenations

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestHyphenations(t *testing.T) {
	type testData struct {
		in  string
		out string
	}

	tests := []testData{
		{in: "ethnic", out: "h-nic"},
		{in: "realistic", out: "al-is-tic"},
		{in: "5*2b", out: ""},
		{in: "clinic", out: "n-ic"},
		{in: "alcoholic", out: "l-ic"},
		{in: "terrific", out: "f-ic"},
		{in: "public", out: "b-lic"},
		{in: "analytic", out: "-lyt-ic"},
		{in: "metallic", out: "l-lic"},
		{in: "bicyclic", out: "-clic"},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.out, hyphenation(test.in)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
