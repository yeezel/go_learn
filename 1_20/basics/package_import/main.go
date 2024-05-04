/*
1. 一个文件夹可以作为 package，同一个 package 内部变量、类型、方法等定义可以相互看到
2. 一个目录只一个包名，不同目录不能使用同一个包名
3. 含有main函数的目录package名必须为main
4. 所有的包名都应该使用小写字母
*/
package main

// 推荐写法
import (
	"fmt"                                  // 默认导入
	. "fmt"                                //可以直接使用函数，不用前缀调用
	out "fmt"                              // 别名导入
	_ "golearn/basics/package_import/init" // 表示不使用该包，而是只是使用该包的init()函数
	"golearn/basics/package_import/pi"
)

func main() {
	pi.Say()
	fmt.Println("googogo")
	out.Println("alias")
	Println("gogog")
}
