package twice

//find in a signed sequential list (or file) an integer that appears at least twice

func atleastTwice() {

}

func populateSequentialList(list []int, n int) {
	odd := n % 2
	midPoint := n / 2
	for i := -midPoint - odd; i < midPoint; i++ {
		list[i+midPoint+odd] = i
	}
}
