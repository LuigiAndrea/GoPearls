package charclassification

func isUpper(r rune) bool {
	return (charstable[r] & upper) != 0
}

func isLower(r rune) bool {
	return (charstable[r] & lower) != 0
}

func isAlpha(r rune) bool {
	return (charstable[r] & (upper | lower)) != 0
}

func isAlphaNum(r rune) bool {
	return (charstable[r] & (upper | lower | digit)) != 0
}

func isAlphaNum2(r rune) bool {
	return isDigit2(r) || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isPunct(r rune) bool {
	return (charstable[r] & punct) != 0
}

func isDigit(r rune) bool {
	return (charstable[r] & digit) != 0
}

func isDigit2(r rune) bool {
	return r >= '0' && r <= '9'
}

func isBlank(r rune) bool {
	return (charstable[r] & blank) != 0
}
