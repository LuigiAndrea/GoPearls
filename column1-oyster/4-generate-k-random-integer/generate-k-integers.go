package kintegers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var integers []int

//Filename of the file created with k random integers
const Filename = "./kIntegers.data"

func swap(x, y int) {
	temp := integers[x]
	integers[x] = integers[y]
	integers[y] = temp
}

//MinMaxInterval contains the range of numbers to save in the file
type MinMaxInterval struct {
	Min int
	Max int
}

//CreateFileWithRandomIntegers create a file with k integers in random order
func CreateFileWithRandomIntegers(mmi ...MinMaxInterval) error {
	var k int

	for _, v := range mmi {
		if v.Max <= v.Min || v.Max < 0 || v.Min < 0 {
			return fmt.Errorf("Out of Range parameters (nMin:%d, nMax:%d)", v.Min, v.Max)
		}

		for i, value := k, v.Min; value < v.Max; i++ {
			integers = append(integers, value)
			value++
		}

		k += v.Max - v.Min
	}

	rand.Shuffle(k, swap)

	file, err := os.Create(Filename)
	if err != nil {
		return fmt.Errorf("Unable to create file '%s': %s", Filename, err)
	}

	writer := bufio.NewWriter(file)

	defer file.Close()

	for _, v := range integers {
		writer.WriteString(fmt.Sprintf("%d ", v))
	}
	writer.Flush()

	return nil
}
