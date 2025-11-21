package piscine

func ReacPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	} else {
		return 2 * (w + h)
	}
}
