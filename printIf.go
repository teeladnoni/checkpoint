package piscine

func PrintIf(s string) string {
	if s == "" || len(s) >= 3 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
