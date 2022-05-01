package main

import "fmt"

type Liters float64
type Gallons float64
type Title string
type Milliliters float64
type MyType string
type Number int

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264) 
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264) 
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (n *Number) Double() {
	*n *= 2
}

func (m MyType) WithReturn() int {
	return len(m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())

	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)
	value := MyType("MyType value")
	fmt.Println(value.WithReturn())
	value.MethodWithParameters(4, true)
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()

	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
	// var carFuel Gallons
	// var busFuel Liters
	// carFuel = Gallons(10.0)
	// busFuel = Liters(240.0)
	// fmt.Println(carFuel, busFuel)
	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Zodiac"))
	fmt.Println(Title("Alien") > Title("Zodiac"))
	// fmt.Println(Title("Alien") + "s")
	// fmt.Println(Title("Jaws 2") - " 2")

	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Liters(5.5) - 2.2)
	fmt.Println(Liters(1.2) == 1.2)
	fmt.Println(Liters(1.2) < 3.4)
}