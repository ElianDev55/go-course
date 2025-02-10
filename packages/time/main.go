package main

import (
	"fmt"
	"time"
)

func main () {
	p := fmt.Println
	now := time.Now()
	p("Hello World")
	p(now)

	createTime := time.Date(2019, 10, 31, 14, 58, 0, 0, time.UTC)
	p(createTime)

	addHour := createTime.Add(time.Hour * 1)
	p(addHour)

	res := createTime.Add(time.Hour * -1)
	p(res)

	p("Year:")
	p(createTime.Year())
	p("Month:")
	p(createTime.Month())
	p("Day:")
	p(createTime.Day())
	p("Hour:")
	p(createTime.Hour())
	p("Minute:")
	p(createTime.Minute())
	p("Second:")
	p(createTime.Second())
	p("Nanosecond:")
	p(createTime.Nanosecond())

	p("createTime is after now:", createTime.After(now))
	p("createTime is before now:", createTime.Before(now))
	p("createTime is equal to now:", createTime.Equal(now))

}
