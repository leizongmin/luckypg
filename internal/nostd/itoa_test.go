package nostd

import (
	"testing"
)

func TestItoa(t *testing.T) {
	println(12345, Itoa(12345))
	println(-12345, Itoa(-12345))
	println(0, Itoa(0))
}
