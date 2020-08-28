package searchtuning

import "math"

var x = [8]int{1, 5, 8, 2, 6, 43, 21, math.MaxInt32} //fix size array with 7 elements and one sentinel
var xOrig = []int{1, 5, 8, 2, 6, 43, 21}

func seqSearchOriginal(value int) int {
	size := len(xOrig)
	for i := 0; i < size; i++ {
		if xOrig[i] == value {
			return i
		}
	}
	return -1
}

func seqSearch2(value int) int {
	i, size := 0, len(x)
	x[size-1] = value //sentinel

	for i = 0; ; i++ {
		if x[i] == value {
			break
		}
	}

	if i == size-1 {
		return -1
	}

	return i

}

func seqSearch3(value int) int {
	i, size := 0, len(x)
	x[size-1] = value //sentinel

	for i = 0; ; i += size - 1 {
		if x[i] == value {
			break
		}
		if x[i+1] == value {
			i++
			break
		}
		if x[i+2] == value {
			i += 2
			break
		}
		if x[i+3] == value {
			i += 3
			break
		}
		if x[i+4] == value {
			i += 4
			break
		}
		if x[i+5] == value {
			i += 5
			break
		}
		if x[i+6] == value {
			i += 6
			break
		}
	}

	if i == size-1 {
		return -1
	}

	return i

}
