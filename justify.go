package justify

import (
	"strings"
)

func Justify(text string, width int) string {
	words := strings.Fields(text)
	lines := [][]string{}
	line := []string{}
	lineLen := 0

	// Break words into lines
	for _, word := range words {
		if lineLen+len(word)+len(line) > width {
			lines = append(lines, line)
			line = []string{}
			lineLen = 0
		}
		line = append(line, word)
		lineLen += len(word)
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}

	var result strings.Builder
	for i, l := range lines {
		if i == len(lines)-1 || len(l) == 1 { // Last line or single word, left-align
			result.WriteString(strings.Join(l, " "))
			result.WriteString("\n")
			continue
		}
		// Justify the line
		totalSpaces := width
		for _, word := range l {
			totalSpaces -= len(word)
		}
		gaps := len(l) - 1
		space := totalSpaces / gaps
		extra := totalSpaces % gaps

		for j, word := range l {
			result.WriteString(word)
			if j < gaps {
				result.WriteString(strings.Repeat(" ", space))
				if extra > 0 {
					result.WriteString(" ")
					extra--
				}
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}
