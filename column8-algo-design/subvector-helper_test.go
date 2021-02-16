// +build all column8 subvectorhelper

package subvector

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	vector []float64
	size   int
}

func TestSubvector(t *testing.T) {
	if err := a.AssertDeepEqual(NewSubvector(2, 4, 13.05), Subvector{From: 2, To: 4, DistanceToNumber: 13.05}); err != nil {
		t.Error(err)
	}

	tests := []testData{
		{vector: []float64{-2.4, 5.7, -3.31, 5, 3, -2, 2}, size: 7},
		{vector: []float64{-2.4}, size: 1},
		{vector: nil, size: 0},
		{vector: []float64{}, size: 0},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.size, ValidateAndReturnSize(test.vector)); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
