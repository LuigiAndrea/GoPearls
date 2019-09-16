package twice

//find in a signed sequential list (or file) an integer that appears at least twice

//Return the number in the sequential list that appear at least twice, return false if it doesn't exist
func searchAtLeastTwiceValue(list []int) (bool, int) {
	n := len(list)
	midPoint := n / 2
	res := searchValue(list, 0, n-1, midPoint, n)
	if res != n {
		return true, res
	}
	return false, res
}

func searchValue(list []int, left, right, midpoint, currentTwice int) int {
	if left > right {
		return currentTwice
	}

	i := (left + right) / 2
	if i-midpoint == list[i] {
		// check right half
		return searchValue(list, i+1, right, midpoint, currentTwice)
	}
	//check left half
	return searchValue(list, left, i-1, midpoint, list[i])
}
