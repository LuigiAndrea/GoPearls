package segments

import "fmt"

const spaceBetweenDigit, widthSegment, heightSegment = 3, 8, 6

type display [2 * heightSegment][5 * (spaceBetweenDigit + widthSegment)]rune

var table display
var char = 'X'

func fillDisplay(number []byte) {
	segmentsToCheck := []byte{segments[2], segments[3], segments[4], segments[1], segments[5], segments[0], segments[6]}
	for _, seg := range segmentsToCheck {
		for offset, digit := range number {
			if isSegmentActive(digit, seg) {
				fillSegmentTable(seg, (spaceBetweenDigit+widthSegment)*offset)
			}
		}
	}
}

func fillSegmentTable(segment byte, offset int) {
	switch segment {
	case segments[2]:
		buildHorizontalSegment(0, offset)
	case segments[1]:
		buildHorizontalSegment(heightSegment-1, offset)
	case segments[0]:
		buildHorizontalSegment(2*(heightSegment-1), offset)
	case segments[3]:
		buildVerticalSegment(0, offset)
	case segments[5]:
		buildVerticalSegment(heightSegment-1, offset)
	case segments[4]:
		buildVerticalSegment(0, offset+widthSegment-1)
	case segments[6]:
		buildVerticalSegment(heightSegment-1, offset+widthSegment-1)
	}
}

func buildVerticalSegment(row, column int) {
	for i := 0; i < heightSegment; i++ {
		table[row+i][column] = char
	}
}

func buildHorizontalSegment(row, column int) {
	for i := 0; i < widthSegment; i++ {
		table[row][column+i] = char
	}
}

func isSegmentActive(number byte, segment byte) bool {
	return number&segment == segment
}

func printDisplay(number uint16) {
	numberEncoded := displayNumber(number)
	fillDisplay(numberEncoded)
	for _, row := range table {
		for _, c := range row {
			if c == char {
				fmt.Print(string(c))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	cleanDisplay()
}

func cleanDisplay() {

	for i, row := range table {
		for j := range row {
			table[i][j] = ' '
		}
	}
}
