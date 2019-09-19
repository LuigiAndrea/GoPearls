package ksubset

import "sort"

//Given a set of n real numbers, a real number t, and an integer k,
// how quickly can you determine whether there exists a k-element subset of the set that sums to at most t?

// n*log(n), true if the k-element subset exists
func existKSubsetSumToT(list []int, t, k int) bool {
	listToCheck := sort.IntSlice(list)
	listToCheck.Sort()
	result := sumElements(listToCheck[:k])
	if result <= t {
		return true
	}

	return false
}

func sumElements(elements []int) int {
	sum := 0
	for _, v := range elements {
		sum += v
	}
	return sum
}
