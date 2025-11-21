package piscine

// import "github.com/01-edu/z01"

func CheckNumber(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
