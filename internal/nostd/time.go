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

// FIXME: 由于 tinygo 目前对 cgo 支持还有 bug，暂时不能直接返回 struct
int get_local_time_year() { return get_local_time().year; }
int get_local_time_month() { return get_local_time().month; }
int get_local_time_day() { return get_local_time().day; }
int get_local_time_wday() { return get_local_time().wday; }
int get_local_time_hour() { return get_local_time().hour; }
int get_local_time_min() { return get_local_time().min; }
int get_local_time_sec() { return get_local_time().sec; }
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
	return Time{
		year:    int(C.get_local_time_year()),
		month:   int(C.get_local_time_month()),
		day:     int(C.get_local_time_day()),
		weekday: int(C.get_local_time_wday()),
		hour:    int(C.get_local_time_hour()),
		minute:  int(C.get_local_time_min()),
		second:  int(C.get_local_time_sec()),
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
