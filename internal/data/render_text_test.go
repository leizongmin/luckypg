package data

import (
	"fmt"
	"testing"

	"github.com/leizongmin/luckypg/internal/nostd"
	"github.com/leizongmin/luckypg/internal/testutil"
)

func TestRenderText(t *testing.T) {
	today := nostd.TimeNow()

	test := func(s string) {
		r := RenderText(today, s)
		fmt.Println(r)
		testutil.Assert(s != r)
	}

	test("v=%v")
	test("t=%t")
	test("l=%l")
	test("v=%v t=%t l=%l")
}
