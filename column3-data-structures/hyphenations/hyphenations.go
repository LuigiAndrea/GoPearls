package hyphenations

import (
	"strings"
)

//Contains only valid hyphenations of words that end in the letter "c", rules must be applied in this order
var hyphenations = []string{"et-ic", "al-is-tic", "s-tic", "p-tic", "-lyt-ic", "ot-ic",
	"an-tic", "n-tic", "c-tic", "at-ic", "h-nic", "n-ic", "m-ic", "l-lic", "b-lic", "-clic",
	"l-ic", "h-ic", "f-ic", "d-ic", "-bic", "a-ic", "-mac", "i-ac"}

var hyphenationMap map[string]string

func hyphenation(word string) (suffix string) {
	if len(hyphenationMap) == 0 {
		buildHyphenationMap()
	}

	size := len(word)
	var key string
	for i := 0; i < size; i++ {
		key = word[i:size]
		if v, ok := hyphenationMap[key]; ok {
			suffix = v
			break
		}
	}

	return suffix
}

func buildHyphenationMap() map[string]string {
	hyphenationMap = make(map[string]string, len(hyphenations))
	for _, v := range hyphenations {
		hyphenationMap[strings.ReplaceAll(v, "-", "")] = v
	}

	return hyphenationMap
}
