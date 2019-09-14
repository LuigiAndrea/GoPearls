package anagram

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(dictionary []string) {
	mySign := buildSignature(dictionary)
	mySign.Sort()
	for _, el := range mySign {
		fmt.Println(el)
	}
}

func buildSignature(dic []string) Signature {
	signatures := make(Signature, len(dic))
	signatures.Initialize()
	for i, v := range dic {
		sign := sort.StringSlice(strings.Split(string(v), ""))
		sort.Stable(sign)
		signatures[i][0], signatures[i][1] = strings.Join(sign, ""), v
	}
	return signatures
}

//Signature used to create a nx2 signature structure
type Signature [][]string

//Initialize the signature structure
func (s Signature) Initialize() {
	for row := range s {
		s[row] = make([]string, len(s))
	}
}

//Sort the signature structure using the signature key
func (s Signature) Sort() {
	sort.Stable(s)
}

func (s Signature) Len() int           { return len(s) }
func (s Signature) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s Signature) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
