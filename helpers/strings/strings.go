package strings

func CountOfRunesInString(s string, r rune) int {
	res := 0
	for _, char := range s {
		if char == r {
			res++
		}
	}
	return res
}

func StringsContainSameCharsAndLen(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	setA := make(map[rune]struct{})
	setB := make(map[rune]struct{})

	exists := struct{}{}

	for _, char := range a {
		setA[char] = exists
	}
	for _, char := range b {
		setB[char] = exists
	}

	for k := range setA {
		_, ok := setB[k]
		if !ok {
			return false
		}
	}

	for k := range setB {
		_, ok := setA[k]
		if !ok {
			return false
		}
	}

	return true
}
