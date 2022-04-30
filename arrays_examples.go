package main

import (
	"fmt"
	"time"
)

func main() {
	// var notes [7]string
	// notes[0] = "do"
	// notes[1] = "re"
	// notes[2] = "mi"
	// fmt.Println(notes[0])
	// fmt.Println(notes[1])
	// fmt.Println(notes[3])
	// fmt.Println(notes[6])

	// var primes [5]int
	// primes[0] = 2
	// primes[1] = 3
	// fmt.Println(primes[0])
	// fmt.Println(primes[2])
	// fmt.Println(primes[4])

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Printf("%#v\n", notes)
	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Printf("%#v\n", primes)
	index := 1
	fmt.Println(index, notes[index])
	index = 3
	fmt.Println(index, notes[index])

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}

	for index, value := range notes {
		fmt.Println(index, value)
	}
 
	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"togerher on a single line",

	}
	fmt.Println(text[0])
}
