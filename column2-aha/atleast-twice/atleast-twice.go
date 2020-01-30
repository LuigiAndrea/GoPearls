package twice

// Find in a signed sequential list (or file) an integer that appears at least twice.

//Return the number in the sequential list that appear at least twice, return false if it doesn't exist
//Assume the list is sequential
func searchAtLeastTwiceValue(list []int) (bool, int) {
	n := len(list)
	var res int

	if n != 0 {
		res = searchValue(list, 0, n-1, list[0], n)
	} else {
		res = searchValue(list, 0, n-1, 0, n)
	}

	if res != n {
		return true, res
	}
	return false, res
}

func searchValue(list []int, left, right, startValue, currentTwice int) int {
	if left > right {
		return currentTwice
	}

	i := (left + right) / 2
	if list[i] == startValue+i {
		// check right half
		return searchValue(list, i+1, right, startValue, currentTwice)
	}
	//check left half
	return searchValue(list, left, i-1, startValue, list[i])
}
