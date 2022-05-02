package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year int
	Month int
	Day int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
}


func (d *Date) SetDay(day int) {
	if day < 1 || day 31 {
		return errors.New("invalid day")
	}
	d.Day = day
}

func main() {
	// date := Date{Year: 2019, Month: 5, Day: 27}
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	date.SetMonth(5)
	date.SetDay(27)
	fmt.Println(date)
}