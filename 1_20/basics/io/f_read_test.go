// 从控制台读取输入:
package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  string
	inputReader            *bufio.Reader
	err                    error
)

func TestReadDemo() {
	// scan方式读取用户输入
	// scan()
	// bufio方式读取用户输入
	// bufior()
	// 读取命令行参数
	// args()
	// 使用flag包解析命令参数
	// flagParse()
}

func args() {
	who := "Alice "
	println(os.Args[0]) // 文件绝对路径
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}

func scan() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	// scan方式赋值给变量
	fmt.Sscanf("56.12 / 5212 / Go", "%f / %d / %s", &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}
func bufior() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}
