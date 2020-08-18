package operations

//Operation o
type Operation struct {
	Values, cumm []float64
	Size         int
}

// NewOperation create a new operation object
func NewOperation(size int) Operation {

	if size < 0 {
		panic("Operation size must be greater than 0")
	}

	Op := Operation{}
	Op.Size = size
	for i := 0; i < size; i++ {
		Op.Values = append(Op.Values, 0.0)
	}

	Op.cumm = make([]float64, len(Op.Values))
	return Op
}

// LoadOperation load data in operation type
//if data slice length is greater than operation length 'lop' then LoadOperation loads the first lop elements from data in operation type
func (o *Operation) LoadOperation(data []float64) {
	for i, v := range data {
		o.RunOperation(v, i, i)
	}

	computeValues(o)
}

// RunOperation assign value v from position l to u, return true if the assignements are successfully, false otherwise
func (o *Operation) RunOperation(v float64, l, u int) bool {

	if l < 0 || l > u || u >= o.Size {
		return false
	}

	o.cumm[u] += v
	if l > 0 {
		o.cumm[l-1] -= v
	}

	computeValues(o)
	return true
}

// computeValues Compute the values of Operation type
func computeValues(o *Operation) {
	o.Values[o.Size-1] = o.cumm[o.Size-1]
	for i := o.Size - 2; i >= 0; i-- {
		o.Values[i] = o.Values[i+1] + o.cumm[i]
	}
}

// Reset Reset Operation type
func (o *Operation) Reset() {
	*o = NewOperation(o.Size)
}
