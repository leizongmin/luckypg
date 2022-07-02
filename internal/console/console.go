package console

import (
	"strings"

	"github.com/acarl005/stripansi"
	"golang.org/x/term"
)

var TerminalWidth = 0
var TerminalHeight = 0

func init() {
	width, height, err := term.GetSize(0)
	if err == nil {
		TerminalWidth = width
		TerminalHeight = height
	}
}

func getStringLength(s string) int {
	count := 0
	runes := []rune(stripansi.Strip(s))
	for _, r := range runes {
		if r > 255 {
			count += 2
		} else {
			count++
		}
	}
	return count
}

func CenteredText(s string, n int) string {
	if n > 0 {
		lines := strings.Split(s, "\n")
		maxLength := 0
		for _, line := range lines {
			if len(line) > maxLength {
				maxLength = getStringLength(line)
			}
		}
		for i, line := range lines {
			line := line + strings.Repeat(" ", maxLength-getStringLength(line))
			left := strings.Repeat(" ", (n-maxLength)/2)
			right := strings.Repeat(" ", (n-maxLength)/2)
			lines[i] = left + line + right
		}
		return strings.Join(lines, "\n")
	}
	return s
}

func FixedWidthText(s string, n int) string {
	if n > 0 {
		lines := strings.Split(s, "\n")
		maxLength := 0
		for _, line := range lines {
			if len(line) > maxLength {
				maxLength = getStringLength(line)
			}
		}
		for i, line := range lines {
			line := line + strings.Repeat(" ", maxLength-getStringLength(line))
			right := strings.Repeat(" ", (n - maxLength))
			lines[i] = line + right
		}
		return strings.Join(lines, "\n")
	}
	return s
}
