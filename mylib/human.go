package mylib

import "fmt"

var Public string = "Public"
var Private string = "private"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human")
}
