package column1

import (
	"bufio"
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
