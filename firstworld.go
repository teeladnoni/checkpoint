package piscine

func FirstWord(s string) string {
	word := ""
	found := false // The Switch

	for _, r := range s {
		if r != ' ' {
			// It's a letter!
			found = true      // Turn switch ON
			word += string(r) // Add to word
		} else if found {
			// It's a space, AND we have already started the word.
			// This means the word is over.
			break
		}
		// If it's a space and found is false, we do nothing (skip it).
	}

	return word + "\n"
}
