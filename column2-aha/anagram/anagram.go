package anagram

import (
	"sort"
	"strings"
)

//ListAnagrams list all the anagrams of a dictionary pass as argument
func ListAnagrams(dictionary []string) map[string][]string {
	mySign := buildSignature(dictionary)
	mySign.Sort()
	return squash(mySign)
}

func buildSignature(dic []string) Signature {
	signatures := NewSignature(len(dic))
	for i, v := range dic {
		sign := sort.StringSlice(strings.Split(v, ""))
		sort.Stable(sign)
		signatures[i][0], signatures[i][1] = strings.Join(sign, ""), v
	}
	return signatures
}

func squash(signatures Signature) map[string][]string {
	signSquash := make(map[string][]string)
	var tempSignSquash signatureSquash
	signature := signatures[0][0]
	tempSignSquash.add(signatures[0][1])

	for i := 1; i < len(signatures); i++ {
		if signature == signatures[i][0] {
			tempSignSquash.add(signatures[i][1])
		} else {
			signSquash[signature] = tempSignSquash
			signature = signatures[i][0]
			tempSignSquash.clean() //empty the slice
			tempSignSquash.add(signatures[i][1])
		}
	}
	signSquash[signature] = tempSignSquash

	return signSquash
}

//NewSignature add a new Signature struct
func NewSignature(size int) Signature {
	signatures := make(Signature, size)
	signatures.Initialize()
	return signatures
}

//Signature used to create a nx2 signature structure
type Signature [][]string

//Initialize the signature structure
func (s Signature) Initialize() {
	for row := range s {
		s[row] = make([]string, 2)
	}
}

//Sort the signature structure using the signature key
func (s Signature) Sort()              { sort.Stable(s) }
func (s Signature) Len() int           { return len(s) }
func (s Signature) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s Signature) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type signatureSquash []string

func (s *signatureSquash) clean()           { *s = nil }
func (s *signatureSquash) add(value string) { *s = append(*s, value) }
