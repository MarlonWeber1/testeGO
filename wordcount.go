package testeGO

import (
	"strings"
	"unicode/utf8"
)

// CountWords receives a text and returns a map with the frequency of words.
// It converts words to lowercase, removes simple punctuation, and ignores words
// with less than 3 characters (runes).
func CountWords(text string) map[string]int {
	frequencies := make(map[string]int)

	// Split the text into fields (words) using whitespace.
	words := strings.Fields(text)

	for _, word := range words {
		// Convert to lowercase.
		lowerWord := strings.ToLower(word)

		// Remove simple punctuation.
		var cleanedBuilder strings.Builder
		for _, r := range lowerWord {
			// Check if the character is a simple punctuation mark.
			// Simple punctuation include: , . ! ? ; :
			// Let's filter these out.
			if r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == ':' {
				continue
			}
			cleanedBuilder.WriteRune(r)
		}
		cleaned := cleanedBuilder.String()

		// Ignore words with less than 3 characters (runes).
		// We use utf8.RuneCountInString to count runes instead of bytes,
		// ensuring correct handling of accented characters.
		if utf8.RuneCountInString(cleaned) >= 3 {
			frequencies[cleaned]++
		}
	}

	return frequencies
}
