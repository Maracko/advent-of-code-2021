package maps

func StringInMap(m map[int]string, s string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}
