package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL string
	Size int
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i ++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// myChannel := make(chan string)
	// go send(myChannel)
	// reportNap("receiving goroutine", 5)
	// fmt.Println(<-myChannel)
	// fmt.Println(<-myChannel)

	// channel1 := make(chan string)
	// channel2 := make(chan string)
	// go abc(channel1)
	// go def(channel2)
	// fmt.Print(<-channel1)
	// fmt.Print(<-channel2)
	// fmt.Print(<-channel1)
	// fmt.Print(<-channel2)
	// fmt.Print(<-channel1)
	// fmt.Print(<-channel2)
	// fmt.Print()

	// myChannel := make(chan string)
	// go greeting(myChannel)
	// fmt.Println(<-myChannel)

	// go a()
	// go b()
	// time.Sleep(time.Second)
	// fmt.Println("end main()")
	// var size int

	pages := make(chan Page)
	urls := []string{"https://example.com/",
					"https://golang.org/",
					"https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	// go responseSize("https://example.com", sizes)
	// go responseSize("https://golang.org", sizes)
	// go responseSize("https://golang.org/doc", sizes)
	// fmt.Println(<-sizes)
	// fmt.Println(<-sizes)
	// fmt.Println(<-sizes)
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Println("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Println("b")
	}
}