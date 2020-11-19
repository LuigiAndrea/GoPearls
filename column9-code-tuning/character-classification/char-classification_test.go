// +build all column9 charclassification charmasks

package charclassification

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	char                                                                                  rune
	upperValue, lowerValue, alphaValue, alphaNumValue, punctValue, digitValue, blankValue bool
}

var tests = []testData{
	testData{char: 'B', upperValue: true, lowerValue: false, alphaValue: true, alphaNumValue: true, punctValue: false, digitValue: false, blankValue: false},
	testData{char: '%', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: true, digitValue: false, blankValue: false},
	testData{char: 'z', upperValue: false, lowerValue: true, alphaValue: true, alphaNumValue: true, punctValue: false, digitValue: false, blankValue: false},
	testData{char: '1', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: true, punctValue: false, digitValue: true, blankValue: false},
	testData{char: ' ', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: false, digitValue: false, blankValue: true},
	testData{char: '\n', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: false, digitValue: false, blankValue: false},
	testData{char: '\t', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: false, digitValue: false, blankValue: true},
	testData{char: '}', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: true, digitValue: false, blankValue: false},
	testData{char: 48 /*char 0*/, upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: true, punctValue: false, digitValue: true, blankValue: false},
}

func TestIsUpper(t *testing.T) {

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.upperValue, isUpper(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestIsLower(t *testing.T) {
	for i, test := range tests {
		if err := a.AssertDeepEqual(test.lowerValue, isLower(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestIsAlpha(t *testing.T) {
	for i, test := range tests {
		if err := a.AssertDeepEqual(test.alphaValue, isAlpha(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestIsAlphaNum(t *testing.T) {
	for i, test := range tests {
		if err := a.AssertDeepEqual(test.alphaNumValue, isAlphaNum(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestIsPunct(t *testing.T) {
	for i, test := range tests {
		if err := a.AssertDeepEqual(test.punctValue, isPunct(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestIsDigit(t *testing.T) {
	for i, test := range tests {
		if err := a.AssertDeepEqual(test.digitValue, isDigit(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestIsBlank(t *testing.T) {
	for i, test := range tests {
		if err := a.AssertDeepEqual(test.blankValue, isBlank(test.char)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
