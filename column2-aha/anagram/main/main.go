package main

import (
	"fmt"

	"github.com/LuigiAndrea/GoPearls/column2-aha/anagram"
)

func main() {
	squshSignatures := anagram.ListAnagrams([]string{"marines", "remains", "reductions", "won", "seminar", "rattles", "now",
		"starlet", "startle", "era", "top", "tar", "rat", "pot", "east", "introduces",
		"are", "own", "opt", "eats", "sate", "seat", "teas", "arm", "ear", "discounter",
	})

	for key, el := range squshSignatures {
		fmt.Printf("%s --> %s\n", key, el)
	}
}
