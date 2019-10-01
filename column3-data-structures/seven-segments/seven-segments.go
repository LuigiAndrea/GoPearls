package segments

import (
	"GoPearls/utilities"
	"math"
)

//segments
//   2_
// 3|1_|4
// 5|0_|6
var segments [7]byte

func displayNumber(number int16) []byte {
	var display []byte
	buildSegments()

	for number > 0 {
		display = append(display, encodeDigit(number%10))
		number /= 10
	}

	size := len(display)
	utilities.Reverse(utilities.ByteSlice(display), 0, size-1)

	return display
}

func buildSegments() {
	for i := 0; i < 7; i++ {
		segments[i] = 1 << uint(i)
	}
}

func encodeDigit(digit int16) byte {
	var encodedDigit byte
	switch digit {
	case 0:
		encodedDigit = allSegmentsON() &^ activeSegments(1)
	case 1:
		encodedDigit = activeSegments(4, 6)
	case 2:
		encodedDigit = allSegmentsON() &^ activeSegments(3, 6)
	case 3:
		encodedDigit = allSegmentsON() &^ activeSegments(3, 5)
	case 4:
		encodedDigit = allSegmentsON() &^ activeSegments(0, 2, 5)
	case 5:
		encodedDigit = allSegmentsON() &^ activeSegments(4, 5)
	case 6:
		encodedDigit = allSegmentsON() &^ activeSegments(4)
	case 7:
		encodedDigit = activeSegments(2, 4, 6)
	case 8:
		encodedDigit = allSegmentsON()
	case 9:
		encodedDigit = allSegmentsON() &^ activeSegments(0, 5)

	}

	return encodedDigit
}

func allSegmentsON() byte {
	return byte(math.Exp2(7) - 1)
}

func activeSegments(segmentNumber ...int) byte {
	var activeSegment byte
	for _, v := range segmentNumber {
		activeSegment |= segments[v]
	}
	return activeSegment
}
