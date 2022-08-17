package nostd

import (
	"fmt"
	"testing"
	"time"

	"github.com/leizongmin/luckypg/internal/testutil"
)

func TestTestTime(t *testing.T) {
	a := TimeNow()
	b := time.Now()
	fmt.Printf("%+v\n", a)

	fmt.Println(a.Year() == b.Year(), "\t|\t", b.Year(), "\t|\t", a.Year())
	testutil.Assert(a.Year() == b.Year())

	fmt.Println(int(a.Month()) == int(b.Month()), "\t|\t", b.Month(), "\t|\t", a.Month(), "\t|\t", int(a.Month()), "\t|\t", int(b.Month()))
	testutil.Assert(int(a.Month()) == int(b.Month()))

	fmt.Println(a.Day() == b.Day(), "\t|\t", b.Day(), "\t|\t", a.Day())
	testutil.Assert(a.Day() == b.Day())

	fmt.Println(int(a.Weekday()) == int(b.Weekday()), "\t|\t", b.Weekday(), "\t|\t", a.Weekday())
	testutil.Assert(int(a.Weekday()) == int(b.Weekday()))

	fmt.Println(a.Hour() == b.Hour(), "\t|\t", b.Hour(), "\t|\t", a.Hour())
	testutil.Assert(a.Hour() == b.Hour())

	fmt.Println(a.Minute() == b.Minute(), "\t|\t", b.Minute(), "\t|\t", a.Minute())
	testutil.Assert(a.Minute() == b.Minute())

	fmt.Println(a.Second() == b.Second(), "\t|\t", b.Second(), "\t|\t", a.Second())
	testutil.Assert(a.Second() == b.Second())
}
