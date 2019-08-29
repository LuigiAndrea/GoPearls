package rotate

//using temp slices
func rotateLeftSlice(str []string, shiftLeft int) {
	tempStr := make([]string, shiftLeft)
	diff := len(str) - shiftLeft
	copy(tempStr, str[:shiftLeft])
	copy(str[:diff], str[shiftLeft:])
	copy(str[diff:], tempStr)
}

//without uaing temp slice
func rotateLeftReverse(str []string, shiftLeft int) {
	lastElement := len(str) - 1
	reverse(str, 0, shiftLeft-1)
	reverse(str, shiftLeft, lastElement)
	reverse(str, 0, lastElement)
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
