package column1

import (
	"errors"
	"fmt"
)

type bitVector []int

var bitsPerWord = 32
var size = 0

func newBitVector(n int) (bv *bitVector, err error) {
	if n < 0 {
		return bv, errors.New("Negative n argument in newBitVector")
	}

	size = (n / bitsPerWord) + 1
	bitVector := make(bitVector, size)
	return &bitVector, err
}

func (bit bitVector) get(index int) (byte, error) {

	if err := bit.validateInput(index); err != nil {
		return 0, err
	}

	word := bit[index/bitsPerWord]

	if word&(1<<uint(index%bitsPerWord)) > 0 {
		return 1, nil
	}

	return 0, nil
}

func (bit bitVector) set(index int) {
	bit[index/bitsPerWord] |= 1 << uint(index%bitsPerWord)
}

func (bit bitVector) clear(index int) {
	bit[index/bitsPerWord] &= ^(1 << uint(index%bitsPerWord))
}

func (bit bitVector) validateInput(index int) error {
	if index < 0 || index >= size*bitsPerWord {
		return fmt.Errorf("Out of Range index argument")
	}
	return nil
}
