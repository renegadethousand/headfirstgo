package main

import (
	"fmt"
	"log"
)

var metersPerLiter float64 = 10.0

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	var total float64
	metersPerLiter = 10.0
	amount, err := paintNeeded(4.2, 3.0)
	fmt.Println(err)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount, err := paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount, err := paintNeeded(5.0, 3.3)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / metersPerLiter, nil
}