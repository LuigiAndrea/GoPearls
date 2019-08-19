package column1

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

var integers []int

func swap(x, y int) {
	temp := integers[x]
	integers[x] = integers[y]
	integers[y] = temp
}

func createFileWithRandomIntegers(nMin int, nMax int) error {
	if nMax <= nMin || nMax < 0 || nMin < 0 {
		return fmt.Errorf("Out of Range parameters (nMin:%d, nMax:%d)", nMin, nMax)
	}

	k := nMax - nMin
	integers = make([]int, k)
	var bufferIntegers bytes.Buffer

	for i, value := 0, nMin; value < nMax; i++ {
		integers[i] = value
		value++
	}

	rand.Shuffle(k, swap)

	for _, v := range integers {
		bufferIntegers.WriteString(fmt.Sprintf("%d ", v))
	}

	file, err := os.Create("./kIntegers.data")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return err
	}
	defer file.Close()
	bufferIntegers.WriteTo(file)

	return nil
}
