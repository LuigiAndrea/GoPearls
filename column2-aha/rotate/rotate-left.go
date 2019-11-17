package rotate

import (
	"github.com/LuigiAndrea/GoPearls/utilities"
	"fmt"
	"sort"
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
	strToReverse := sort.StringSlice(str)
	lastElement := len(strToReverse) - 1
	utilities.Reverse(strToReverse, 0, shiftLeft-1)
	utilities.Reverse(strToReverse, shiftLeft, lastElement)
	utilities.Reverse(strToReverse, 0, lastElement)
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

func swapRange(str sort.StringSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		str.Swap(a+i, b+i)
	}
}

type juggleNum int

func (j *juggleNum) subtractIfGreaterOrEqualThan(quantityToSubtract int) {
	qToSub := juggleNum(quantityToSubtract)
	if *j >= qToSub {
		*j -= qToSub
	}
}

func (j *juggleNum) juggleNumToInt() int {
	return int(*j)
}

func (j *juggleNum) setJuggleNum(value int) {
	*j = juggleNum(value)
}

func rotateLeftJuggling(str []string, shiftLeft int) error {
	if err := validateShiftLeft(str, shiftLeft); err != nil {
		return err
	}
	n := len(str)
	steps := gcd(shiftLeft, n)

	for i := 0; i < steps; i++ {
		temp := str[i]
		j := i
		k := juggleNum(shiftLeft + j)
		k.subtractIfGreaterOrEqualThan(n)

		for i != k.juggleNumToInt() {
			str[j] = str[k]
			j = k.juggleNumToInt()
			k.setJuggleNum(shiftLeft + j)
			k.subtractIfGreaterOrEqualThan(n)
		}
		str[j] = temp
	}

	return nil
}

func gcd(x, y int) int {
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}
	if x == 0 {
		return y
	}
	return x
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
