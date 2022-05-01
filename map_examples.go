package main

import (
	"fmt"
	"sort"
)

func main() {
	// var ranks map[string]int
	// ranks = make(map[string]int)
	// ranks := make(map[string]int)
	// ranks["gold"] = 1
	// ranks["silver"] = 2
	// ranks["bronze"] = 3
	// fmt.Println(ranks["bronze"])
	// fmt.Println(ranks["gold"])

	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["Li"] = "Lithium"
	// fmt.Println(elements["Li"])
	// fmt.Println(elements["H"])

	// isPrime := make(map[int]bool)
	// isPrime[4] = false
	// isPrime[7] = true
	// fmt.Println(isPrime[4])
	// fmt.Println(isPrime[7])

	// ranks := map[string]int{"bronze": 3, "silever": 2, "gold": 1}
	// fmt.Println(ranks["gold"])
	// fmt.Println(ranks["bronze"])
	// elements := map[string]string{
	// 	"H": "Hydrogen",
	// 	"Li": "lithium",
	// }
	// fmt.Println(elements["H"])
	// fmt.Println(elements["Li"])

	emptyMap := map[string]float64{}
	fmt.Println(emptyMap)

	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])

	words := make(map[string]string)
	words["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", words["I've been assigned"])
	fmt.Printf("%#v\n", words["I haven't been assigned"])

	// counters := make(map[string]int)
	// counters["a"]++
	// counters["a"]++
	// counters["c"]++
	// fmt.Println(counters["a"], counters["b"], counters["c"])

	// var nilMap map[int]string
	// fmt.Printf("%#v\n", nilMap)
	// nilMap[3] = "three"

	var myMap map[int]string = make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)

	status("Alma")
	status("Carl")

	// counters := map[string]int{"a": 3, "b": 0}
	// var value int
	// var ok bool
	// value, ok = counters["a"]
	// fmt.Println(value, ok)
	// value, ok = counters["b"]
	// fmt.Println(value, ok)
	// value, ok = counters["c"]
	// fmt.Println(value, ok)

	// counters := map[string]int{"a": 3, "b": 0}
	// var ok bool
	// _, ok = counters["a"]
	// fmt.Println(ok)
	// _, ok = counters["b"]
	// fmt.Println(ok)
	// _, ok = counters["c"]
	// fmt.Println(ok)

	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

	mapRange()
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func mapRange() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}