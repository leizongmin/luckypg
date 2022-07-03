package nostd

import (
	"fmt"
	"testing"
	"time"
)

func TestTestTime(t *testing.T) {
	a := TimeNow()
	b := time.Now()
	fmt.Printf("%+v\n", a)
	fmt.Println(b.Year(), a.Year(), a.Year() == b.Year())
	fmt.Println(b.Month(), a.Month(), int(a.Month()), int(b.Month()), int(a.Month()) == int(b.Month()))
	fmt.Println(b.Day(), a.Day(), a.Day() == b.Day())
	fmt.Println(b.Weekday(), a.Weekday(), int(a.Weekday()) == int(b.Weekday()))
	fmt.Println(b.Hour(), a.Hour(), a.Hour() == b.Hour())
	fmt.Println(b.Minute(), a.Minute(), a.Minute() == b.Minute())
	fmt.Println(b.Second(), a.Second(), a.Second() == b.Second())
}
