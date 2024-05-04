package libraw

import (
	"flag"
	"fmt"
	"os"
)

func FlagDemo() {
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

	fmt.Println("newline: ", *NewLine)
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
	os.Stdout.WriteString(s + Newline)
}
