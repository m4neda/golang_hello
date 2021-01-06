package main

import (
	"io"
	"log"
	"os"
)

// LoggingSettings Settting log format
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multilogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multilogFile)
}

func main() {
	LoggingSettings("test.log")

	// _, err := os.Open("faege")
	// if err != nil {
	// 	log.Fatalln("Exit", err)
	// }
	// log.Println("logging!")
	// log.Printf("%T %v", "test", "test")

	// log.Fatalf("%T %v", "test", "test")
	// log.Fatalln("error")
}
