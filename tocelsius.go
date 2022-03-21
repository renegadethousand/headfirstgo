// pass_fail сообщает, сдал ли пользователь экзамен
package main

import (
	"fmt"
	"github.com/headfirstgo/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f defrees Celsius\n", celsius)
}