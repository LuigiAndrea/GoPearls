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

var size int // size of the coffee can

const (
	unknown = iota
	black
	white
)

//ToString for colorBean type
func (color colorBean) String() string {
	names := [...]string{
		"Unknown",
		"Black",
		"White"}

	if color < black || color > white {
		return names[unknown]
	}

	return names[color]
}

//Prove that the process ends
func coffeeCanBeans(coffeeCan []bean) (bean, error) {
	size = len(coffeeCan)

	if err := validateCanBean(coffeeCan); err != nil {
		return bean{}, err
	}

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
	if size < 1 {
		return fmt.Errorf("Coffee Can must have at least one bean")
	}

	for _, v := range beans {
		if v.color >= black+white || v.color <= unknown {
			return fmt.Errorf("Bean color %s", v.color.String())
		}
	}
	return nil
}
