package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	for i, v := range l {
		fmt.Println(i, v)
	}
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
}
