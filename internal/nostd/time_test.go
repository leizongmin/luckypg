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
	fmt.Println(a.Year() == b.Year(), "\t|\t", b.Year(), "\t|\t", a.Year())
	fmt.Println(int(a.Month()) == int(b.Month()), "\t|\t", b.Month(), "\t|\t", a.Month(), "\t|\t", int(a.Month()), "\t|\t", int(b.Month()))
	fmt.Println(a.Day() == b.Day(), "\t|\t", b.Day(), "\t|\t", a.Day())
	fmt.Println(int(a.Weekday()) == int(b.Weekday()), "\t|\t", b.Weekday(), "\t|\t", a.Weekday())
	fmt.Println(a.Hour() == b.Hour(), "\t|\t", b.Hour(), "\t|\t", a.Hour())
	fmt.Println(a.Minute() == b.Minute(), "\t|\t", b.Minute(), "\t|\t", a.Minute())
	fmt.Println(a.Second() == b.Second(), "\t|\t", b.Second(), "\t|\t", a.Second())
}
