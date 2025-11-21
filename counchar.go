package piscine

func CountChar(s string, r rune) int {
	count := 0
	for _, char := range s {
		if char == r {
			count++
		}
	}
	return count
}
