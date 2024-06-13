package runestuff

import (
	"fmt"
	"unicode/utf8"
)

func Get() {
	// Rune for the ⌘ character (U+2318)
	startRune := '⌘'

	// Display the starting rune and its code point
	fmt.Printf("Starting rune: %c, Code point: U+%04X\n", startRune, startRune)

	// Iterate backwards to get the preceding 100 code points
	for i := 1; i <= 100; i++ {
		prevRune := startRune - rune(i)
		// Check if the rune is valid and printable
		if !utf8.ValidRune(prevRune) {
			fmt.Printf("Code point U+%04X is not a valid rune\n", prevRune)
			continue
		}
		fmt.Printf("Previous rune: %c, Code point: U+%04X\n", prevRune, prevRune)
	}
}
