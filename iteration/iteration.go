package iteration

import (
	"strings"
)

// Repeat takes a character and repeats it as much as repeatCount times
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
