package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

/*
 Go1.16 引入了//go:embed功能，可以将资源文件内容直接打包到二进制文件，方便部署
 //go:embed可以将任何文件或者文件夹的内容打包到编译出的可执行文件中
 在embed中，可以将静态资源文件嵌入到三种类型的变量，分别为：字符串、字节数组、embed.FS文件类型
 注意事项：
 1. 需要手动import embed
 2. //go:embed指令只能用在包一级的变量中，不能用在函数或方法级别
 3. 当包含目录时，它不会包含以“.”或“ “开头的文件还有符号链接。但是如果使用通配符，比如dir/*，它将包含所有匹配的文件
*/

// embed会读取当前目录下的version.txt存放到变量里
//
//go:embed version.txt
var version string

//go:embed version.txt
var versionByte []byte

// static是一个目录
//
//go:embed static
var embededFiles embed.FS

func main() {
	fmt.Printf("str version %q\n", version)
	fmt.Printf("byte version %q\n", string(versionByte))

	/*
		验证embed是否依赖目录中的文件
		1. 执行go run main.go live 运行非编译进二进制，然后修改index.html文件内容，刷新页面发现内容改变
		2. 执行go run main.go 运行编译进二进制，然后修改index.html文件内容，刷新页面没有改变内容
	*/
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	http.ListenAndServe(":8888", nil)
}

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}

	log.Print("using embed mode")

	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
