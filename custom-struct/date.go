package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) setYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) setMonth(month int) {
	d.Month = month
}

func (d *Date) setDay(day int) {
	d.Day = day
}

func main() {
	date := Date{}
	err := date.setYear(0)
	if err != nil {
		log.Fatal(err)
	}

	date.setMonth(5)
	date.setDay(27)
	fmt.Println(date)
}
