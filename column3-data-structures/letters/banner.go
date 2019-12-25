package banner

import (
	"strconv"
	"strings"
)

//encoded as number of character to print for each row, --> "4b1*4b" means 4 blank, 1* and four blank for the first row
var letters = map[rune][]string{'A': []string{"4b1*4b", "3b1*1b1*3b", "2b5*2b", "1b1*5b1*1b", "1*7b1*"},
	'E': []string{"5*", "1*4b", "1*4b", "5*", "1*4b", "1*4b", "5*"},
}

//Banner take a Capital Letter as input and produces as output the letter depicted graphically as string
func Banner(letter rune) string {
	return printAsString(letters[letter])
}

func printAsString(letterEncode []string) string {
	var letter strings.Builder

	for _, c := range letterEncode {
		letter.WriteString(decodeLetter(c))
		letter.WriteString("\n")
	}

	return letter.String()
}

func decodeLetter(code string) string {
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

func buildString(n int, char string) string {
	var result strings.Builder

	for i := 0; i < n; i++ {
		result.WriteString(char)
	}

	return result.String()
}
