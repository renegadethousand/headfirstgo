package main

import (
	"fmt"
	"math"
)

func main() {
	// var notes []string
	// notes = make([]string, 7)
	// notes[0] = "do"
	// notes[1] = "re"
	// notes[2] = "mi"
	// fmt.Println(notes[0])
	// fmt.Println(notes[1])

	// primes := make([]int, 5)
	// primes[0] = 2
	// primes[1] = 3
	// fmt.Println(primes[0])

	// fmt.Println(len(notes))
	// fmt.Println(len(primes))

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}

	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	primes := []int{
		2,
		3,
		5,
	}
	fmt.Println(primes[0], primes[1], primes[2])

	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[:]
	fmt.Println(slice1)

	// slice := []string{"a", "b"}
	// fmt.Println(slice, len(slice))
	// slice = append(slice, "c")
	// fmt.Println(slice, len(slice))
	// slice = append(slice, "d", "e")
	// fmt.Println(slice, len(slice))

	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])

	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v\n", intSlice)

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v\n", slice)

	severalInts(1)
	severalInts(1, 2, 3)

	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()

	mix(1, true, "a", "b")
	mix(2, false, "a", "b", "c", "d")
	mix(3, true)

	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))

	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))

	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))

	intSlice2 := []int{1, 2, 3}
	severalInts(intSlice2...)
	stringSlice2 := []string{"a", "b", "c", "d"}
	mix(1, true, stringSlice2...)
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max;
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}