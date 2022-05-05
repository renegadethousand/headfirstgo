package main

import (
	"fmt"
	// "log"
)

func Socialize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eh?")
	return nil
}

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func freakOut() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
	// fmt.Println(recover())
	// err := Socialize()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}