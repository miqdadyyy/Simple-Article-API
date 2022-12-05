package str

import "fmt"

func StringEllipsis(text string, characterCount int) string {
	if characterCount > len(text) {
		return text
	}
	return fmt.Sprintf("%s...", text[:characterCount])
}
