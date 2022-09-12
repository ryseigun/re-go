package basic

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())
	fmt.Println("Day :", now.Day())
	fmt.Println("Hour :", now.Hour())
	fmt.Println("Minute :", now.Minute())
	fmt.Println("Nanosecond :", now.Nanosecond())

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-12-01")
	fmt.Println("Time: ", parse)
}
