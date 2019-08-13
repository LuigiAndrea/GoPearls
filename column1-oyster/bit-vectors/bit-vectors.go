package column1

import "errors"

type bitVector []int

var bitsPerWord = 32

func newBitVector(n int) (bv *bitVector, err error) {
	if n < 0 {
		return bv, errors.New("Negative n argument in newBitVector")
	}

	size := (n / bitsPerWord) + 1
	bitVector := make(bitVector, size)
	return &bitVector, err
}

func (bit bitVector) get(index int) byte {
	word := bit[index/bitsPerWord]

	if word&(1<<uint(index%bitsPerWord)) > 0 {
		return 1
	}

	return 0
}

func (bit bitVector) set(index int) {
	bit[index/bitsPerWord] |= 1 << uint(index%bitsPerWord)
}

func (bit bitVector) clear(index int) {
	bit[index/bitsPerWord] &= ^(1 << uint(index%bitsPerWord))
}
