package bitvector

import "fmt"

//BitVector Type BitVector
type BitVector []int

var bitsPerWord = 32
var size = 0

//NewBitVector create a new BitVector
func NewBitVector(n int) (bv *BitVector, err error) {
	if n < 0 {
		return bv, NegativeBitVectorSize(n)
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

	if word&(1<<uint(index%bitsPerWord)) > 0 {
		return 1, nil
	}

	return 0, nil
}

//Set the bit at position index
func (bit BitVector) Set(index int) error {

	if err := bit.validateInput(index); err != nil {
		return err
	}

	bit[index/bitsPerWord] |= 1 << uint(index%bitsPerWord)

	return nil
}

//Clear the bit at position index
func (bit BitVector) Clear(index int) error {
	if err := bit.validateInput(index); err != nil {
		return err
	}

	bit[index/bitsPerWord] &= ^(1 << uint(index%bitsPerWord))
	return nil
}

func (bit BitVector) validateInput(index int) error {
	if index < 0 || index >= size*bitsPerWord {
		return IndexBitVectorOutOfRange(index)
	}
	return nil
}

//IndexBitVectorOutOfRange fires when the user provide a wrong index value as parameter for Get, Set, and Clear functions
type IndexBitVectorOutOfRange int

func (i IndexBitVectorOutOfRange) Error() string {
	return fmt.Sprintf("Index '%d' for BitVector is Out of Range", i)
}

//NegativeBitVectorSize fires when n<0 for NewBitVector
type NegativeBitVectorSize int

func (n NegativeBitVectorSize) Error() string {
	return fmt.Sprintf("Size '%d' for NewBitVector must be zero or positive", n)
}
