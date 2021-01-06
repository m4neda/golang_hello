package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./study.go")
	if err != nil {
		log.Fatal("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	// initialize count, overwrite err
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))
	// can not use ':=' because err was already initilalized
	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error")
	}
}
