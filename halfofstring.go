package piscine

func HalfOfString(s string) string {
	if s == "" {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	halve := len(s) / 2

	return s[:halve]
}
