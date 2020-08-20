package subvector

//Subvector type
type Subvector struct {
	From, To         int
	DistanceToNumber float64
}

//NewSubvector subvector constructor to fill currentCloseNumber
func NewSubvector(from, to int, value float64) Subvector {
	return Subvector{
		From:             from,
		To:               to,
		DistanceToNumber: value}
}

//ValidateAndReturnSize validate vector and return its size
func ValidateAndReturnSize(vector []float64) int {
	var size int

	if vector == nil {
		return size
	}
	if size = len(vector); size == 0 {
		return size
	}

	return size
}
