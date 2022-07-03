package nostd

/*
#include <time.h>

typedef struct Time {
	int year;
	int month;
	int day;
	int wday;
	int hour;
	int min;
	int sec;
} Time;

Time get_local_time() {
  	time_t rawtime;
    struct tm timeinfo;
    time(&rawtime);
    localtime_r(&rawtime, &timeinfo);

	Time t;
	t.year = timeinfo.tm_year + 1900;
	t.month = timeinfo.tm_mon + 1;
	t.day = timeinfo.tm_mday;
	t.wday = timeinfo.tm_wday;
	t.hour = timeinfo.tm_hour;
	t.min = timeinfo.tm_min;
	t.sec = timeinfo.tm_sec;
	return t;
}
*/
import "C"

type Time struct {
	year    int
	month   int
	day     int
	weekday int
	hour    int
	minute  int
	second  int
}

func (t *Time) Year() int {
	return t.year
}

func (t *Time) Month() Month {
	return Month(t.month)
}

func (t *Time) Day() int {
	return t.day
}

func (t *Time) Weekday() Weekday {
	return Weekday(t.weekday)
}

func (t *Time) Hour() int {
	return t.hour
}

func (t *Time) Minute() int {
	return t.minute
}

func (t *Time) Second() int {
	return t.second
}

func TimeNow() Time {
	t := C.get_local_time()
	return Time{
		year:    int(t.year),
		month:   int(t.month),
		day:     int(t.day),
		weekday: int(t.wday),
		hour:    int(t.hour),
		minute:  int(t.min),
		second:  int(t.sec),
	}
}

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
