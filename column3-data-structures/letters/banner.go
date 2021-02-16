package banner

import (
	"fmt"
	"strconv"
	"strings"
)

//encoded as number of characters to print for each row, --> "4b1*4b" means 4 blanks, 1 asterisk and four blanks for the first row
var letters = map[rune][]string{'A': {"4b1*4b", "3b1*1b1*3b", "2b5*2b", "1b1*5b1*1b", "1*7b1*"},
	'E': {"4*", "1*4b", "1*4b", "4*", "1*4b", "1*4b", "4*"},
	'I': {"6*", "2b2*2b", "2b2*2b", "2b2*2b", "6*"},
	'O': {"1b4*1b", "1*4b1*", "1*4b1*", "1*4b1*", "1b4*1b"},
	'U': {"1*4b1*", "1*4b1*", "1*4b1*", "1*4b1*", "1b4*1b"},
}

//Banner takes a Capital Letter as input and produces as output the letter depicted graphically as string
func Banner(letter rune) string {
	return decodeLetter(letters[letter])
}

//decodeLetter takes an array of encoded strings and returns the concatenation of their decoded value
func decodeLetter(letterEncode []string) string {
	var letter strings.Builder

	for _, c := range letterEncode {
		fmt.Fprintf(&letter, "%s\n", decodeString(c))
	}

	return letter.String()
}

//decodeString takes an encoded string and returns its expanding version. Ex: "2*1b1*" -> "** *"
func decodeString(code string) string {
	var result strings.Builder
	var num int
	var char string

	for i := 0; i < len(code); i = i + 2 {
		num, _ = strconv.Atoi(string(code[i]))
		char = getChar(code[i+1])
		result.WriteString(buildString(num, char))
	}

	return result.String()
}

//getChar return space string if the input is a 'b' character, otherwise return the string representation of the byte
func getChar(c byte) string {
	if c == 98 { //b char
		return " "
	}

	return string(c)
}

//buildString writes the character 'char' n times
func buildString(n int, char string) string {
	var result strings.Builder

	for i := 0; i < n; i++ {
		result.WriteString(char)
	}

	return result.String()
}
