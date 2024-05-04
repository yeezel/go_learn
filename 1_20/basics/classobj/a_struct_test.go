package classobj

import (
	"fmt"
	"testing"
)

type innerS struct {
	in1     int
	in2     int
	subject float32
}

// 同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）
type A struct{ a int }
type B struct{ a int }

// 框架/工具可以通过反射获取到某个字段定义的属性，采取相应的处理方式。
type Books struct {
	title   string `json: title_name` //tag 可以理解为 struct 字段的注解，可以用来定义字段的一个或多个属性。
	subject string // 外层名字会覆盖内层名字（但是两者的内存空间都保留）
	book_id int
	int     // 匿名字段，类型就是字段的名字
	innerS  // 内嵌结构体
	A       // 同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）
	B       // 同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）
}

// 开头字母小写表示不暴露给外部
type ebook Books

func TestStructDemo(t *testing.T) {
	var book1 Books /* 声明 Book1 为 Books 类型 */
	book1.title = "Go 语言"
	printBook(book1)

	// 其他创建结构体方式
	var book2 *Books = new(Books)
	printBook1(book2)
	book3 := &Books{"Java 教程", "www.java.com", 6412300, 60, innerS{5, 10, 3}, A{1}, B{2}}
	printBook1(book3)
	var book4 = Books{title: "rust 教程", subject: "www.rust.com"}
	printBook(book4)

	//使用工厂方法创建结构体
	ebok := NewBook()
	fmt.Printf("ebook title : %s\n", ebok.title)

	//结构体转换
	book5 := ebook{title: "e 教程", subject: "www.e.com"}
	book6 := Books(book5) // 类型转换
	printBook(book6)
}
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
}

func printBook1(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
}

// 让外部包强制使用该工厂方法获取ebook对象
func NewBook() *ebook {
	//如果是结构体，new(file) 和 &file{} 是等价的
	return &ebook{}
}
