package formgenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func formLetterGenerator(filename string, placeholders ...string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Unable to open file '%s': %s", filename, err)
	}

	defer file.Close()
	var form strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(&form, scanner.Text())
	}

	formString := form.String()
	for i, v := range placeholders {
		formString = strings.ReplaceAll(formString, fmt.Sprintf("#%v", i+1), v)
	}

	return formString, nil
}
