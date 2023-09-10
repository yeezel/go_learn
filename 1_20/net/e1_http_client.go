package net

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
这个结构会保存解析后的返回数据。
他们会形成有层级的 XML，可以忽略一些无用的数据
*/
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func HttpClient() {
	resp, err := http.Head("https://baidu.com/")
	checkError1(err)
	fmt.Println("baidu: ", resp.Status)

	res, err := http.Get("http://localhost:8088/test1")
	checkError1(err)
	data, err := io.ReadAll(res.Body)
	checkError1(err)
	fmt.Printf("Got: %q", string(data))

}

func checkError1(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
