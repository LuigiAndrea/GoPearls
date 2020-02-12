// You are initially given a coffee can that contains some black beans and some white beans and a large pile of “extra” black beans.
// You then repeat the following process until there is a single bean left in the can.
// Randomly select two beans from the can. If they are the same color, throw them both out and insert an extra black bean.
// If they are different colors, return the white bean to the can and throw out the black.

package coffeebean

import (
	"fmt"
	"math/rand"
	"time"
)

type bean struct {
	color colorBean
}

type colorBean int

const unknown = "Unknown"
const (
	black = iota + 1
	white
)

//ToString for colorBean type
func (color colorBean) String() string {
	names := [...]string{
		unknown,
		"Black",
		"White"}

	if color < black || color > white {
		return unknown
	}

	return names[color]
}

func coffeeCanBeans(coffeeCan []bean) (bean, error) {
	if err := validateCanBean(coffeeCan); err != nil {
		return bean{}, err
	}
	size := len(coffeeCan)
	if size == 1 {
		return coffeeCan[0], nil
	}

	rand.Seed(time.Now().UnixNano())

	for ; size > 2; size-- {
		posA, posB := rand.Intn(size), rand.Intn(size-1)
		if posB == posA {
			posB++
		}
		removeBeans(&coffeeCan, posA, posB)
	}

	removeBeans(&coffeeCan, 0, 1)
	return coffeeCan[0], nil
}

func haveSameColor(A, B bean) bool {
	return A == B
}

func removeBeans(coffeeCan *[]bean, posA, posB int) {
	beanA, beanB := (*coffeeCan)[posA], (*coffeeCan)[posB]
	if haveSameColor(beanA, beanB) {
		*coffeeCan = append(*coffeeCan, bean{color: black})
		delete(coffeeCan, posA)
		delete(coffeeCan, posB)
	} else if beanA.color == black {
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

func validateCanBean(beans []bean) error {
	if len(beans) < 1 {
		return fmt.Errorf("Coffee Can must have at least one bean")
	}

	for _, v := range beans {
		if v.color.String() == unknown {
			return fmt.Errorf("Bean color %s", unknown)
		}
	}
	return nil
}
