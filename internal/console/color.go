package console

import (
	"github.com/leizongmin/luckypg/internal/nostd"
)

func XTermColor(fg int, bg int, s string) string {
	return "\x1B[48;5;" + nostd.Itoa(bg) + "m\x1B[38;5;" + nostd.Itoa(fg) + "m" + s + "\x1B[39m\x1B[49m"
}

type RenderColorFunc func(s string) string

func XTermColorFunc(fg int, bg int) RenderColorFunc {
	return func(s string) string {
		return "\x1B[48;5;" + nostd.Itoa(bg) + "m\x1B[38;5;" + nostd.Itoa(fg) + "m" + s + "\x1B[39m\x1B[49m"
	}
}

func NoColorRenderFunc(s string) string {
	return s
}
