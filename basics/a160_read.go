// 从控制台读取输入:
package main

import (
	"bufio"
	"flag"
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

func main() {
	// scan方式读取用户输入
	// scan()
	// bufio方式读取用户输入
	// bufior()
	// 读取命令行参数
	// args()
	// 使用flag包解析命令参数
	flagParse()
}

// 命令行 (Windows) 中执行：echo.exe A B C，将输出：A B C；执行 echo.exe -n A B C 将输出一行一个字母
func flagParse() {
	//flag.Bool() 定义了一个默认值是 false 的 flag
	var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
	const (
		Space   = " "
		Newline = "\n"
	)
	flag.PrintDefaults() //打印 flag 的使用帮助信息
	flag.Parse()         // 扫描参数列表（或者常量列表）并设置 flag
	var s string = ""
	//flag.Narg() 返回参数的数量。解析后 flag 或常量就可用了
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i) //表示第 i 个参数
	}
	os.Stdout.WriteString(s)
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
