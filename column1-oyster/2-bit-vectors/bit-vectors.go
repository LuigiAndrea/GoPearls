package bitvector

import (
	"fmt"
)

//BitVector Type BitVector
type BitVector []int

var bitsPerWord = 32
var size = 0

//NewBitVector create a new BitVector
func NewBitVector(n int) (bv *BitVector, err error) {
	if n < 0 {
		return bv, fmt.Errorf("Negative n '%d' argument in NewBitVector", n)
	}

	size = (n / bitsPerWord) + 1
	bitVector := make(BitVector, size)
	return &bitVector, err
}

func (bit BitVector) Get(index int) (byte, error) {

	if err := bit.validateInput(index); err != nil {
		return 0, err
	}

	word := bit[index/bitsPerWord]

	if word&(1<<uint(index%bitsPerWord)) > 0 {
		return 1, nil
	}

	return 0, nil
}

func (bit BitVector) Set(index int) error {

	if err := bit.validateInput(index); err != nil {
		return err
	}

	bit[index/bitsPerWord] |= 1 << uint(index%bitsPerWord)

	return nil
}

func (bit BitVector) Clear(index int) error {
	if err := bit.validateInput(index); err != nil {
		return err
	}

	bit[index/bitsPerWord] &= ^(1 << uint(index%bitsPerWord))
	return nil
}

func (bit BitVector) validateInput(index int) error {
	if index < 0 || index >= size*bitsPerWord {
		return fmt.Errorf("Index '%d' is Out of Range", index)
	}
	return nil
}
