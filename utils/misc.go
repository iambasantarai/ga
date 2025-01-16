package utils

// Devanagari characters live in the U+0900 to U+097F block
func IsDevanagariCharacter(ch rune) bool {
	return ch >= '\u0900' && ch <= '\u097F'
}
