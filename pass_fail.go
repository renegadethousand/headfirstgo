// pass_fail сообщает, сдал ли пользователь экзамен
package main

import (
	"fmt"
	"github/headfirstgo/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}