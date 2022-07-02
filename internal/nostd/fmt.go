package nostd

import (
	"strconv"
)

func Println(s string) {
	println(s)
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}
