package main

import (
	"fmt"
	"github/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Postal Code:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joe Carr"}
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)

	// address := magazine.Address{
	// 	Street: "123 Oak St",
	// 	City: "Omaha",
	// 	State: "NE",
	// 	PostalCode: "68111",
	// }
	// subscriber := magazine.Subscriber{Name: "Aman Singh"}
	// subscriber.HomeAddress = address
	// fmt.Println(subscriber.HomeAddress)
	// // var address magazine.Address
	// // address.Street = "123 Oak St"
	// // address.City = "Omaha"
	// // address.State = "NE"
	// // address.PostalCode = "68111"
	// // fmt.Println(address)

	// var employee magazine.Employee
	// employee.Name = "Joy Carr"
	// employee.Salary = 60000
	// fmt.Println(employee.Name)
	// fmt.Println(employee.Salary)

	// subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	// subscriber := magazine.Subscriber{Rate: 4.99}

	// fmt.Println("Name:", subscriber.Name)
	// fmt.Println("Rate:", subscriber.Rate)
	// fmt.Println("Active:", subscriber.Active)

	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}