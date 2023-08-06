package main

import (
	"fmt"
	"log"
	"github.com/headfirstgo/encapsulation/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	event := calendar.Event{}
	err = event.SetTitle("jgseojgseojgseojgseojgseo")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year(), event.Title())
}