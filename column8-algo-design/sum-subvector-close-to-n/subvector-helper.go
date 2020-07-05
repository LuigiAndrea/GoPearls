package sumsubvectornumber

type subvector struct {
	from, to         int
	distanceToNumber float64
}

//subvector constructor to fill currentCloseNumber
func newSubvector(from, to int, value float64) subvector {
	return subvector{
		from:             from,
		to:               to,
		distanceToNumber: value}
}

func validateAndReturnSize(vector []float64) int {
	var size int

	if vector == nil {
		return size
	}
	if size = len(vector); size == 0 {
		return size
	}

	return size
}
