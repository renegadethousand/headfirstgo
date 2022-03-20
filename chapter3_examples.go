package main

import (
	"fmt"
	"errors"
	"log"
	"math"
	"reflect"
)

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
	var myBool bool = true
	printPointer(&myBool)
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
	myInt := 3456
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)

	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myFloat int
	var myFloatPointer *int
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
	amount := 60
	fmt.Println(amount)
	fmt.Println(&amount)
	root, err := squareRoot(-9.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f", root)
	}

	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)

	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)

	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error())
	fmt.Println(err)

	err = errors.New("height can't be negative")
	fmt.Println(err)
	log.Fatal(err)

	fmt.Println(status(60.1))
	fmt.Println(status(59))

	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))

	sayHi();

	stringHello := "Hello"
	repeatLine(stringHello, 5)

	fmt.Println(stringHello)
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Println(resultString)

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n", 3.1415)
	fmt.Println()

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")

	fmt.Println("-----------------------------")

	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
	line = "qeqwe"
}

func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade < 60.0 {
		return "failing"
	}
	return "passing"
}

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(number *int) {
	*number *= 2
}
