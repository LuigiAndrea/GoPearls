package rotate

import (
	"fmt"
)

//using temp slices
func rotateLeftSlice(str []string, shiftLeft int) error {
	if err := validateShiftLeft(str, shiftLeft); err != nil {
		return err
	}

	tempStr := make([]string, shiftLeft)
	diff := len(str) - shiftLeft
	copy(tempStr, str[:shiftLeft])
	copy(str[:diff], str[shiftLeft:])
	copy(str[diff:], tempStr)
	return nil
}

//without using temp slice
func rotateLeftReverse(str []string, shiftLeft int) error {
	if err := validateShiftLeft(str, shiftLeft); err != nil {
		return err
	}

	lastElement := len(str) - 1
	reverse(str, 0, shiftLeft-1)
	reverse(str, shiftLeft, lastElement)
	reverse(str, 0, lastElement)
	return nil
}

func rotateLeftSwapRange(str []string, shiftLeft int) error {
	if err := validateShiftLeft(str, shiftLeft); err != nil {
		return err
	}

	n := len(str)

	if shiftLeft == 0 || shiftLeft == n {
		return nil
	}

	var i, j = shiftLeft, n - shiftLeft

	for i != j {
		if i > j {
			swapRange(str, shiftLeft-i, shiftLeft, j)
			i -= j
		} else {
			swapRange(str, shiftLeft-i, shiftLeft+j-i, i)
			j -= i
		}
	}

	swapRange(str, shiftLeft-i, shiftLeft, i)
	return nil
}

func swapRange(str []string, a, b, n int) {
	for i := 0; i < n; i++ {
		swap(str, a+i, b+i)
	}
}

func reverse(str []string, i, j int) {
	for ; j > i; i++ {
		swap(str, i, j)
		j--
	}
}

func swap(str []string, x, y int) {
	temp := str[x]
	str[x] = str[y]
	str[y] = temp
}

//ShiftLeftOutOfRange fires when the user provide a wrong shifteft value as parameter for rotate functions
type ShiftLeftOutOfRange int

func (s ShiftLeftOutOfRange) Error() string {
	return fmt.Sprintf("ShiftLength %d is Out of Range", s)
}

func validateShiftLeft(strSlice []string, value int) error {
	if value < 0 || value > len(strSlice) {
		return ShiftLeftOutOfRange(value)
	}
	return nil
}
