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

// 获取文本在控制台的显示宽度，自动忽略 ANSI Color 字符
func TextDisplayLength(s string) int {
	count := 0
	runes := []rune(stripansi.Strip(s))
	for _, r := range runes {
		if r < 256 {
			count++
		} else if r >= 0x2600 && r <= 0x26ff {
			// 符号，https://www.compart.com/en/unicode/block/U+2600
			count++
		} else {
			count += 2
		}
	}
	return count
}

// 每行文本居中，指定最大宽度，两边自动补空格，超出自动截断
func CenteredText(s string, n int) string {
	if n > 0 {
		lines := strings.Split(s, "\n")
		maxLength := 0
		for _, line := range lines {
			if len(line) > maxLength {
				maxLength = TextDisplayLength(line)
			}
		}
		for i, line := range lines {
			line := line + strings.Repeat(" ", maxLength-TextDisplayLength(line))
			n1 := (n - maxLength) / 2
			n2 := n - maxLength - n1
			left := strings.Repeat(" ", n1)
			right := strings.Repeat(" ", n2)
			lines[i] = left + line + right
		}
		return strings.Join(lines, "\n")
	}
	return s
}

// 文本左对齐，指定最大宽度，不足自动补空格，超出自动截断
func FixedWidthText(s string, n int) string {
	if n > 0 {
		lines := strings.Split(s, "\n")
		maxLength := 0
		for _, line := range lines {
			if len(line) > maxLength {
				maxLength = TextDisplayLength(line)
			}
		}
		for i, line := range lines {
			right := strings.Repeat(" ", n-maxLength)
			lines[i] = line + right
		}
		return strings.Join(lines, "\n")
	}
	return s
}
