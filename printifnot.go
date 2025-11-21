package piscine

func PrintIfNot(s string) string {
	if len(s) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
