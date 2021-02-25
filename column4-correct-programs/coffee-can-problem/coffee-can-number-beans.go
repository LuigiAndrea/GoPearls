package coffeebean

//Final remaining bean as a function of the numbers of black and white beans originally in the can
func coffeCanRemainingBean(coffeeCan []bean) (bean, error) {
	size = len(coffeeCan)

	if err := validateCanBean(coffeeCan); err != nil {
		return bean{}, err
	}

	whiteBeans := 0

	for _, bean := range coffeeCan {
		if bean.color == white {
			whiteBeans++
		}
	}

	if whiteBeans%2 == 0 {
		return bean{color: black}, nil
	}

	return bean{color: white}, nil
}
