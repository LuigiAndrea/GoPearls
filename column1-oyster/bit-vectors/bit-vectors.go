package bitvector

import (
	"errors"
	"fmt"
)

//BitVector Type BitVector
type BitVector []int

var bitsPerWord = 32
var size = 0

//NewBitVector create a new BitVector
func NewBitVector(n int) (bv *BitVector, err error) {
	if n < 0 {
		return bv, &NegativeBitVectorSize{size: n}
	}

	size = (n / bitsPerWord) + 1
	bitVector := make(BitVector, size)
	return &bitVector, err
}

//Get return 1 if the number is in the bitVector, 0 otherwise
func (bit BitVector) Get(index int) (byte, error) {

	if err := bit.validateInput(index); err != nil {
		return 0, err
	}

	word := bit[index/bitsPerWord]

	if word&(1<<(index%bitsPerWord)) > 0 {
		return 1, nil
	}

	return 0, nil
}

//Set the bit at position index
func (bit BitVector) Set(index int) error {

	if err := bit.validateInput(index); err != nil {
		return err
	}

	bit[index/bitsPerWord] |= 1 << (index % bitsPerWord)

	return nil
}

//Clear the bit at position index
func (bit BitVector) Clear(index int) error {
	if err := bit.validateInput(index); err != nil {
		return err
	}

	bit[index/bitsPerWord] &= ^(1 << (index % bitsPerWord))
	return nil
}

func (bit BitVector) validateInput(index int) error {
	if index < 0 || index >= size*bitsPerWord {
		return &IndexBitVectorOutOfRange{size: index}
	}
	return nil
}

//IndexBitVectorOutOfRange fires when the user provide a wrong index value as parameter for Get, Set, and Clear functions
type IndexBitVectorOutOfRange struct {
	size int
}

func (idxOutRange IndexBitVectorOutOfRange) Error() string {
	return fmt.Sprintf("Index '%d' for BitVector is Out of Range", idxOutRange)
}

// Is Compare the values of IndexBitVectorOutOfRange
func (idxOutRange *IndexBitVectorOutOfRange) Is(e error) bool {
	var err *IndexBitVectorOutOfRange
	if errors.As(e, &err) {
		return err.size == idxOutRange.size && err.Error() == idxOutRange.Error()
	}

	return false
}

//NegativeBitVectorSize fires when n<0 for NewBitVector
type NegativeBitVectorSize struct {
	size int
}

func (n *NegativeBitVectorSize) Error() string {
	return fmt.Sprintf("Size '%d' for NewBitVector must be zero or positive", n.size)
}

// Is Compare the values of NegativeBitVectorSize
func (n *NegativeBitVectorSize) Is(e error) bool {
	var err *NegativeBitVectorSize
	if errors.As(e, &err) {
		return err.size == n.size && err.Error() == n.Error()
	}

	return false
}
