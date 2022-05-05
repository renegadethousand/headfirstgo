package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func one() {
	two()
}
func two() {
	defer fmt.Println("deferred is one()")
	three()
}
func three() {
	defer fmt.Println("deferred in two()")
	panic("This call stack's too deep for me!")
}

func main() {
	defer reportPanic()
	scanDirectory("my_directory")
	// one()
	// panic("oh, no, we're going down")
	// scanDirectory("my_directory")
	// files, err := ioutil.ReadDir("my_directory")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Println("Directory:", file.Name())
	// 	} else {
	// 		fmt.Println("File:", file.Name())
	// 	}
	// }
}