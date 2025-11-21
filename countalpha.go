package piscine

func CountAlpha(arg string) int {
	count := 0
	for _, s := range arg {
		if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
			count++
		}
	}
	return count
}
