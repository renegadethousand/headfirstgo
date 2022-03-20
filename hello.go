package main

import (
    "fmt"
    "strings"
    "reflect"
)

func main() {
	broken := "G# r#cks"
    replacer := strings.NewReplacer("#", "o")
    fixed := replacer.Replace(broken)
    fmt.Println(fixed)
    fmt.Println(reflect.TypeOf(replacer))
}
