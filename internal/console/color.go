package console

import (
	"fmt"
	"strconv"
)

var (
	Black   = ansiColor("\033[1;30m%s\033[0m")
	Red     = ansiColor("\033[1;31m%s\033[0m")
	Green   = ansiColor("\033[1;32m%s\033[0m")
	Yellow  = ansiColor("\033[1;33m%s\033[0m")
	Purple  = ansiColor("\033[1;34m%s\033[0m")
	Magenta = ansiColor("\033[1;35m%s\033[0m")
	Teal    = ansiColor("\033[1;36m%s\033[0m")
	White   = ansiColor("\033[1;37m%s\033[0m")
)

func ansiColor(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func XTermColor(fg int, bg int, s string) string {
	return "\x1B[48;5;" + strconv.Itoa(bg) + "m\x1B[38;5;" + strconv.Itoa(fg) + "m" + s + "\x1B[39m\x1B[49m"
}

type RenderColorFunc func(s string) string

func XTermColorFunc(fg int, bg int) RenderColorFunc {
	return func(s string) string {
		return "\x1B[48;5;" + strconv.Itoa(bg) + "m\x1B[38;5;" + strconv.Itoa(fg) + "m" + s + "\x1B[39m\x1B[49m"
	}
}

func NoColorRenderFunc(s string) string {
	return s
}
