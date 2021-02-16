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

var funcToRun = []func(rune) bool{isUpper, isLower, isAlpha, isAlphaNum, isAlphaNum2, isPunct, isDigit, isDigit2, isBlank}

var tests = []testData{
	{char: 'B', upperValue: true, lowerValue: false, alphaValue: true, alphaNumValue: true, punctValue: false, digitValue: false, blankValue: false},
	{char: '%', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: true, digitValue: false, blankValue: false},
	{char: 'z', upperValue: false, lowerValue: true, alphaValue: true, alphaNumValue: true, punctValue: false, digitValue: false, blankValue: false},
	{char: '1', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: true, punctValue: false, digitValue: true, blankValue: false},
	{char: ' ', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: false, digitValue: false, blankValue: true},
	{char: '\n', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: false, digitValue: false, blankValue: false},
	{char: '\t', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: false, digitValue: false, blankValue: true},
	{char: '}', upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: false, punctValue: true, digitValue: false, blankValue: false},
	{char: 48 /*char 0*/, upperValue: false, lowerValue: false, alphaValue: false, alphaNumValue: true, punctValue: false, digitValue: true, blankValue: false},
}

func TestCharClassification(t *testing.T) {

	for i, f := range funcToRun {
		t.Run(m.GetFuncName(f), func(t *testing.T) {
			for j, test := range tests {
				if err := a.AssertDeepEqual(getValue(i, test), f(test.char)); err != nil {
					t.Error(m.ErrorMessageTestCount(j+1, err))
				}
			}
		})
	}
}

func getValue(i int, test testData) bool {
	retValue := false
	switch i {
	case 0:
		retValue = test.upperValue
	case 1:
		retValue = test.lowerValue
	case 2:
		retValue = test.alphaValue
	case 3, 4:
		retValue = test.alphaNumValue
	case 5:
		retValue = test.punctValue
	case 6, 7:
		retValue = test.digitValue
	case 8:
		retValue = test.blankValue
	}

	return retValue
}
