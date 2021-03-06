package utilities

//CummulativeArray type
type CummulativeArray []CummulativeValue

func (ca CummulativeArray) Len() int           { return len(ca) }
func (ca CummulativeArray) Less(i, j int) bool { return ca[i].Value < ca[j].Value }
func (ca CummulativeArray) Swap(i, j int)      { ca[i], ca[j] = ca[j], ca[i] }

// CummulativeValue type
// Value, the value in the vector
// OriginalIndex, contains the index before applying a sort or change the positions of the values in the vector
type CummulativeValue struct {
	Value         float64
	OriginalIndex int
}

//GetDistance calculate the difference between two values in the CummulativeArray
func (ca *CummulativeArray) GetDistance(from, to int) float64 {
	size := len(*ca)
	if from < 0 || to < 0 || from >= size || to >= size {
		return -1
	}
	return GetDistance((*ca)[from].Value, (*ca)[to].Value)
}

//CalculateCummulativeArray return a cummulative array
func CalculateCummulativeArray(vector []float64) CummulativeArray {
	size := len(vector)
	cArray := make(CummulativeArray, size)
	sum := 0.0

	for i := 0; i < size; i++ {
		sum, _ = Round(sum+vector[i], 2)
		cArray[i] = CummulativeValue{Value: sum, OriginalIndex: i}
	}

	return cArray
}
