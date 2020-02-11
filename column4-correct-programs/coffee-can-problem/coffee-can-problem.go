// You are initially given a coffee can that contains some black beans and some white beans and a large pile of “extra” black beans.
// You then repeat the following process until there is a single bean left in the can.
// Randomly select two beans from the can. If they are the same color, throw them both out and insert an extra black bean.
// If they are different colors, return the white bean to the can and throw out the black.

package coffeebean

import (
	"math/rand"
	"time"
)

type bean struct {
	color string
}

func coffeeCanBeans(coffeeCan []bean) bean {

	for size := len(coffeeCan); size > 2; size-- {
		rand.Seed(time.Now().UnixNano())
		posA, posB := rand.Intn(size), rand.Intn(size-1)
		if posB == posA {
			posB++
		}
		removeBeans(&coffeeCan, posA, posB)
	}
	removeBeans(&coffeeCan, 0, 1)
	return coffeeCan[0]
}

func haveSameColor(A, B bean) bool {
	return A == B
}

func removeBeans(coffeeCan *[]bean, posA, posB int) {
	beanA, beanB := (*coffeeCan)[posA], (*coffeeCan)[posB]
	if haveSameColor(beanA, beanB) {
		*coffeeCan = append(*coffeeCan, bean{color: "Black"})
		delete(coffeeCan, posA)
		delete(coffeeCan, posB)
	} else if beanA.color == "Black" {
		delete(coffeeCan, posA)
	} else {
		delete(coffeeCan, posB)
	}
}

func delete(coffeeCan *[]bean, pos int) {
	size := len(*coffeeCan) - 1
	(*coffeeCan)[pos] = (*coffeeCan)[size]
	*coffeeCan = (*coffeeCan)[:size]
}
