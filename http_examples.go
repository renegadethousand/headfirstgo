package main

import (
	"log"
	"net/http"
	"fmt"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHi() {
	fmt.Println("Hi")
}
func sayBye() {
	fmt.Println("Bye")
}
func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}
func twice(theFunction func()) {
	theFunction()
	theFunction()
}
func doMath(passedFunction func(int, int) float64) {
	result := passedFunction(10, 2)
	fmt.Println(result)
}
func multiply(a int, b int) float64 {
	return float64(a * b)
}

func main() {
	doMath(divide)
	doMath(multiply)
	
	// var greeterFunction func()
	// var mathFunction func(int, int) float64
	// greeterFunction = sayHi
	// mathFunction = divide
	// greeterFunction()
	// fmt.Println(mathFunction(5, 2))

	// twice(sayHi)
	// twice(sayBye)

	// var myFunction func()
	// myFunction = sayHi
	// myFunction()

	// http.HandleFunc("/hello", englishHandler)
	// http.HandleFunc("/salut", frenchHandler)
	// http.HandleFunc("/namaste", hindiHandler)
	// err := http.ListenAndServe("localhost:8080", nil)
	// log.Fatal(err)
}