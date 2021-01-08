package main

import "fmt"

func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// type assertion
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)
	// type conversion
	var i int = 10
	ii := float64(i)
	fmt.Println(i, ii)
}
