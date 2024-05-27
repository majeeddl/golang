package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	// We’ll start by getting the current time.
	now := time.Now()
	p(now)

	/*
		then we can build a time struct by providing the year, month, day, etc.
		times are always associated with a Location, i.e. time zone.
	*/
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	/*
		You can extract the various components of the time value as expected.
	*/
	p(then.Year())       // 2009
	p(then.Month())      // November
	p(then.Day())        // 17
	p(then.Hour())       // 20
	p(then.Minute())     // 34
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	/*
		The Monday-Sunday Weekday is also available.
	*/
	p(then.Weekday()) // Tuesday

	/*
		These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
	*/
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	/*
		The Sub methods returns a Duration representing the interval between two times.
	*/
	diff := now.Sub(then)
	p(diff) // 28482h25m1.539922763s

	/*
		We can compute the length of the duration in various units.
	*/
	p(diff.Hours())       // 28482.41642775694
	p(diff.Minutes())     // 1708944.9856654164
	p(diff.Seconds())     // 102536699.139925
	p(diff.Nanoseconds()) // 102536699139925000

	/*
		You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	*/
	p(then.Add(diff))  // 2009-11-17 20:34:59.651387237 +0000 UTC
	p(then.Add(-diff)) // 2009-11-17 20:34:57.348612763 +0000 UTC

}
