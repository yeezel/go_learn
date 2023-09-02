// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

/*
不是所有的数据都可以编码为 JSON 类型，只有验证通过的数据结构才能被编码：

	JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map[string]T（T 是 json 包中支持的任何类型）
	Channel，复杂类型和函数类型不能被编码
	不支持循环数据结构；它将引起序列化进入一个无限循环
	指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
*/
func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	// 在 web 应用中最好使用 json.MarshalforHTML() 函数，会对html标签转移
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

	//解码json数据
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		return
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
}
