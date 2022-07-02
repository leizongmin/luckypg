package console

import (
	"strings"

	"github.com/leizongmin/luckypg/internal/nostd"
)

var TerminalWidth = 0
var TerminalHeight = 0

func init() {
	// TODO: get terminal width and height
}

// TextStripAnsi 去除 ANSI 控制字符
func TextStripAnsi(s string) string {
	return nostd.StripAnsi(s)
}

// GetCharDisplayLength 获取字符在控制台的显示宽度
func GetCharDisplayLength(c rune) int {
	if c < 256 {
		return 1
	} else if c >= 0x2600 && c <= 0x26ff {
		// 符号，https://www.compart.com/en/unicode/block/U+2600
		return 1
	} else {
		return 2
	}
}

func TextRemoveExceed(s string, n int) string {
	count := 0
	runes := []rune(TextStripAnsi(s))
	result := make([]rune, 0)
	for _, r := range runes {
		count += GetCharDisplayLength(r)
		if count < n {
			result = append(result, r)
		} else {
			break
		}
	}
	return string(result)
}

// TextDisplayLength 获取文本在控制台的显示宽度，自动忽略 ANSI Color 字符
func TextDisplayLength(s string) int {
	count := 0
	runes := []rune(TextStripAnsi(s))
	for _, r := range runes {
		count += GetCharDisplayLength(r)
	}
	return count
}

// CenteredText 每行文本居中，指定最大宽度，两边自动补空格，超出自动截断
func CenteredText(s string, n int, textColorFunc RenderColorFunc, spaceColorFunc RenderColorFunc) string {
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
			left := ""
			right := ""
			if n1 > 0 {
				left = strings.Repeat(" ", n1)
			}
			if n2 > 0 {
				right = strings.Repeat(" ", n2)
			}
			lines[i] = spaceColorFunc(left) + textColorFunc(line) + spaceColorFunc(right)
		}
		return strings.Join(lines, "\n")
	}
	return s
}

// FixedWidthText 文本左对齐，指定最大宽度，不足自动补空格，超出自动截断
func FixedWidthText(s string, n int, textColorFunc RenderColorFunc, spaceColorFunc RenderColorFunc) string {
	if n > 0 {
		lines := strings.Split(s, "\n")
		maxLength := 0
		for _, line := range lines {
			if len(line) > maxLength {
				maxLength = TextDisplayLength(line)
			}
		}
		for i, line := range lines {
			n1 := n - maxLength
			if n1 > 0 {
				right := strings.Repeat(" ", n1)
				lines[i] = textColorFunc(line) + spaceColorFunc(right)
			} else {
				lines[i] = textColorFunc(line)
			}
		}
		return strings.Join(lines, "\n")
	}
	return s
}
