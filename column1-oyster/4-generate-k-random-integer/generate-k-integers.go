package column1

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

var integers []int

const filename = "./kIntegers.data"

func swap(x, y int) {
	temp := integers[x]
	integers[x] = integers[y]
	integers[y] = temp
}

type minmaxInterval struct {
	min int
	max int
}

func createFileWithRandomIntegers(mmi ...minmaxInterval) error {
	var bufferIntegers bytes.Buffer
	var k int

	for _, v := range mmi {
		if v.max <= v.min || v.max < 0 || v.min < 0 {
			return fmt.Errorf("Out of Range parameters (nMin:%d, nMax:%d)", v.min, v.max)
		}

		for i, value := k, v.min; value < v.max; i++ {
			integers = append(integers, value)
			value++
		}

		k += v.max - v.min

	}

	rand.Shuffle(k, swap)

	for _, v := range integers {
		bufferIntegers.WriteString(fmt.Sprintf("%d ", v))
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return err
	}
	defer file.Close()
	bufferIntegers.WriteTo(file)

	return nil
}
