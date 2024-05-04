package main

import (
	"fmt"

	"golang.org/x/exp/mmap"
)

func main() {
	at, _ := mmap.Open("./tmp.txt")
	buff := make([]byte, 2)
	_, _ = at.ReadAt(buff, 4)
	_ = at.Close()
	fmt.Println(string(buff))
}
