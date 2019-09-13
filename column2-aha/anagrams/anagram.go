package anagram

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(dictionary []string) {
	for _, v := range buildSignature(dictionary) {
		fmt.Println(v)
	}
}

func buildSignature(dic []string) signature {
	signatures := make(signature, len(dic))
	signatures.initialize()
	for i, v := range dic {
		sign := sort.StringSlice(strings.Split(string(v), ""))
		sort.Stable(sign)
		signatures[i][0] = strings.Join(sign, "")
		signatures[i][1] = v
	}
	return signatures
}

//signature used to create a nx2 signature structure
type signature [][]string

//initialize the signature structure
func (s signature) initialize() {
	for row := range s {
		s[row] = make([]string, len(s))
	}
}
