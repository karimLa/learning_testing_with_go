package iteration

import (
	"strings"
)

// Repeat will return the provided character repeated n times
// where n is equal to repeatCount
func Repeat(character string, repeatCount uint) string {
	b := strings.Builder{}
	for i := uint(0); i < repeatCount; i++ {
		b.WriteString(character)
	}
	return b.String()
}
