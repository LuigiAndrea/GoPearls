package ksubset

import (
	"math/rand"
	"sort"
	"time"
)

// O(n*log(n)), true if the k-element subset exists
func existKSubset(list []int, t, k int) bool {
	if r := areInputsValid(list, t, k); r == false {
		return r
	}

	listToCheck := sort.IntSlice(list)
	listToCheck.Sort()
	result := sumElements(listToCheck[:k])
	if result <= t {
		return true
	}
	return false
}

// O(n), true if the k-element subset exists
func existKSubsetQuickSelect(list []int, t, k int) bool {
	if r := areInputsValid(list, t, k); r == false {
		return r
	}

	subsetKQuickSelect(list, 0, len(list)-1, k)
	result := sumElements(list[:k])

	if result <= t {
		return true
	}
	return false
}

func subsetKQuickSelect(list []int, left, right, k int) {
	index := partition(list, left, right)
	currentSmallestElemnts := index + 1 - left

	if currentSmallestElemnts == k {
		return
	}

	if currentSmallestElemnts > k {
		subsetKQuickSelect(list, left, index-1, k)
	} else {
		subsetKQuickSelect(list, index+1, right, k-currentSmallestElemnts)
	}
}

//Partition the list in half using a random pivot, return the index of the pivot
func partition(list []int, left, right int) int {
	intSlice := sort.IntSlice(list)
	var pivot int
	diff := right - left
	if diff < 3 {
		pivot = list[right]
	} else {
		rand.Seed(time.Now().UnixNano())
		pos := rand.Intn(right-left) + left
		pivot = list[pos]
		intSlice.Swap(pos, right)
	}

	i := left
	for j := left; j < right; j++ {
		if intSlice[j] < pivot {
			intSlice.Swap(i, j)
			i++
		}
	}

	intSlice.Swap(i, right)
	return i
}

func sumElements(elements []int) int {
	sum := 0
	for _, v := range elements {
		sum += v
	}
	return sum
}

func areInputsValid(list []int, t, k int) bool {
	size := 0
	res := false
	if list != nil {
		size = len(list)
		if size == 0 || k > size {
			res = false
		} else {
			res = true
		}
	}

	return res
}
