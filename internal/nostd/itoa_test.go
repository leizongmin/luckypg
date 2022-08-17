package nostd

import (
	"strconv"
	"testing"

	"github.com/leizongmin/luckypg/internal/testutil"
)

func TestItoa(t *testing.T) {
	test := func(n int) {
		println(strconv.Itoa(n) == Itoa(n), "\t|\t", strconv.Itoa(n), "\t|\t", Itoa(n))
		testutil.Assert(strconv.Itoa(n) == Itoa(n))
	}
	test(12345)
	test(-12345)
	test(0)
}
