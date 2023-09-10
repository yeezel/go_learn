// gob1.go
package libraw

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

type Address1 struct {
	Type    string
	City    string
	Country string
}

type VCard1 struct {
	FirstName string
	LastName  string
	Addresses []*Address1
	Remark    string
}

var content string

/*
Gob 通常用于远程方法调用参数和结果的传输，以及应用程序和机器之间的数据传输
只有可导出的字段会被编码，零值会被忽略
*/
func GobDemo() {
	// demo()
	demo1()
}

func demo1() {
	pa := &Address1{"private", "Aartselaar", "Belgium"}
	wa := &Address1{"work", "Boom", "Belgium"}
	vc := VCard1{"Jan", "Kersschot", []*Address1{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}

// 以字节缓冲模拟网络传输的简单例子
// Output:   "Pythagoras": {3,4}
func demo() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)
}
