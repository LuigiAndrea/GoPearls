// +build all column8 maxsubvector

package maxsubvector

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	vector        []float64
	expectedValue float64
}

func TestMaxSubVector(t *testing.T) {

	tests := []testData{
		testData{vector: []float64{-2.4, 8, 3.3, -5.1, 2, -13}, expectedValue: 11.3},
		testData{vector: []float64{-2.4, -5.1, -13}, expectedValue: 0},
		testData{vector: []float64{2.4, -5.1, 13}, expectedValue: 13},
		testData{vector: []float64{2.4, -1.1, 13}, expectedValue: 14.3},
		testData{vector: []float64{}, expectedValue: 0},
		testData{vector: nil, expectedValue: 0},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.expectedValue, maxSubvector(test.vector)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
