package main

import "fmt"

func recurses() {
	fmt.Println("Oh, no, I'm stuck!")
	recurses()
}

func count(start int, end int) {
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
}

func main() {
	count(1, 3)
	//recurses()
}