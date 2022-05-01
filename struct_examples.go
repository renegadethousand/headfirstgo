package main

import (
	"fmt"
)

type part struct {
	description string
	count int
}

type car struct {
	name string
	topSpeed float64
}

type subscriber struct {
	name string
	rate float64
	active bool
}

type myStruct struct {
	myField int
}

func doublePack(p *part) {
	p.count *= 2
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)

	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)

	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	pointer.myField = 9
	fmt.Println(pointer.myField)

	// var value int = 2
	// var pointer *int = &value
	// fmt.Println(pointer)
	// fmt.Println(*pointer)

	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)

	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)

	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)

	var bolts1 part
	bolts1.description = "Hex bolts"
	bolts1.count = 24
	showInfo(bolts1)

	// var subscriber1 subscriber
	// subscriber1.name = "Aman singh"
	// fmt.Println("Name:", subscriber1.name)
	// var subscriber2 subscriber
	// subscriber2.name = "Beth Ryan"
	// fmt.Println("Name:", subscriber2.name)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	mySlice := make([]car, 0)
	mySlice = append(mySlice, porsche)
	fmt.Println(mySlice)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description", bolts.description)
	fmt.Println("Count:", bolts.count)
	// var myStruct struct {
	// 	number float64
	// 	word string
	// 	toggle bool
	// }
	// fmt.Printf("%#v\n", myStruct)
	// myStruct.number = 3.14
	// myStruct.word = "pie"
	// myStruct.toggle = true
	// fmt.Println(myStruct.number)
	// fmt.Println(myStruct.word)
	// fmt.Println(myStruct.toggle)

	var subscriber struct {
		name string
		rate float64
		active bool
	}
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println("Name:", subscriber.name)
	fmt.Println("Monthly rate:", subscriber.rate)
	fmt.Println("Active?", subscriber.active)

// 	var subscriber1 struct {
// 		name string
// 		rate float64
// 		active bool
// 	}
// 	subscriber1.name = "Aman singh"
// 	fmt.Println("Name:", subscriber1.name)
// 	var subscriber2 struct {
// 		name string
// 		rate float64
// 		active bool
// 	}
// 	subscriber2.name = "Beth Ryan"
// 	fmt.Println("Name:", subscriber2)
}