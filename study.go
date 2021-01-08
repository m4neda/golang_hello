package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World!")

	b := make([]byte, 8)
	// e := errors.New("EOF")
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// if err == e {
		if err == io.EOF {
			break
		}
	}
}
