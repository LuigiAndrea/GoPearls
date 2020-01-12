// +build all column2 anagram

package anagram

import (
	"fmt"
	"testing"

	goth "github.com/LuigiAndrea/test-helper/assertions"
)

type testData struct {
	keySignature       string
	lengthSignature    int
	expectedSignatures []string
}

func TestAnagram(t *testing.T) {
	squashSignatures := ListAnagrams([]string{"marines", "remains", "reductions", "seminar", "rattles", "now",
		"won", "starlet", "startle", "era", "top", "tar", "rat", "pot", "east", "introduces",
		"are", "own", "opt", "eats", "sate", "seat", "teas", "arm", "ear", "discounter",
	})

	tests := []testData{
		testData{keySignature: "now", lengthSignature: 3, expectedSignatures: []string{"now", "won", "own"}},
		testData{keySignature: "aest", lengthSignature: 5, expectedSignatures: []string{"east", "eats", "sate", "seat", "teas"}},
		testData{keySignature: "boh", lengthSignature: 0, expectedSignatures: []string{}},
		testData{keySignature: "boh", lengthSignature: 0, expectedSignatures: nil},
	}

	for _, test := range tests {
		if lengthSign := len(squashSignatures[test.keySignature]); lengthSign != test.lengthSignature {
			t.Errorf("\nExpected '%d' - Actual '%d'", test.lengthSignature, lengthSign)
		}

		if err := goth.AssertSlicesEqual(goth.StringSlicesMatch{
			Expected: test.expectedSignatures, Actual: squashSignatures[test.keySignature]}); err != nil {
			t.Errorf(err.Error())
		}
	}

	printAnagram(squashSignatures)

}

func printAnagram(squashSignatures map[string][]string) {
	for key, el := range squashSignatures {
		fmt.Printf("%s --> %s\n", key, el)
	}
}
