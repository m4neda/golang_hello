package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	var xf64 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", xf64)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
