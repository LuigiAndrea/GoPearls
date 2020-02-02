package kintegers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var integers []int

func swap(x, y int) {
	integers[x], integers[y] = integers[y], integers[x]
}

//MinMaxInterval contains the range of numbers to save in the file
type MinMaxInterval struct {
	Min int
	Max int
}

//CreateFileWithRandomIntegers create a file with k integers in random order
func CreateFileWithRandomIntegers(filename string, mmi ...MinMaxInterval) error {
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
	rand.Seed(time.Now().Unix())
	rand.Shuffle(k, swap)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Unable to create file '%s': %s", filename, err)
	}

	writer := bufio.NewWriter(file)

	defer file.Close()

	for _, v := range integers {
		writer.WriteString(fmt.Sprintf("%d ", v))
	}
	writer.Flush()

	return nil
}
