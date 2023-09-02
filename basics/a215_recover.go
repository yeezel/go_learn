// panic_recover.go
package main

import (
	"fmt"
	"log"
)

func do(work *Work) {
	panic("bad end")
}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work) // start the goroutine for that work
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Work failed with %s in %v", err, work)
		}
	}()
	do(work)
}

func main() {
	fmt.Printf("Calling test\r\n")
	server()
	fmt.Printf("Test completed\r\n")
}
