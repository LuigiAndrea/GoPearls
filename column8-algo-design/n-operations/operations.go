package operations

var sizeOp int
var cumm []float64

//Operation o
type Operation []float64

// NewOperation create a new operation object
func NewOperation(size int) Operation {
	sizeOp = size
	Op := Operation{}
	for i := 0; i < size; i++ {
		Op = append(Op, 0.0)
	}

	cumm = make([]float64, len(Op))
	return Op
}

// LoadOperation load data in operation type
func (o Operation) LoadOperation(data []float64) {
	for i, v := range data {
		o.RunOperation(v, i, i)
	}
	o.ComputeValues()
}

// RunOperation assign value v from position l to u, return true if the assignements are successfully, false otherwise
func (o Operation) RunOperation(v float64, l, u int) bool {
	if l < 0 || l > u || u >= sizeOp {
		return false
	}

	cumm[u] += v
	if l > 0 {
		cumm[l-1] -= v
	}

	return true
}

// ComputeValues Compute the values of Operation type
func (o Operation) ComputeValues() {
	o[sizeOp-1] = cumm[sizeOp-1]
	for i := sizeOp - 2; i >= 0; i-- {
		o[i] = o[i+1] + cumm[i]
	}
}

// Reset Reset Operation type
func (o Operation) Reset() {
	for i := range o {
		o[i] = 0.0
		cumm[i] = 0.0
	}
}
