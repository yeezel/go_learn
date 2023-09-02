package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type ebook Books

func main() {
	var book1 Books /* 声明 Book1 为 Books 类型 */
	book1.title = "Go 语言"
	book1.author = "www.go.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407
	printBook(book1)

	var book2 *Books = new(Books)
	book2.title = "Python 教程"
	book2.author = "www.python.com"
	book2.subject = "Python 语言教程"
	(*book2).book_id = 6495700
	printBook1(book2)

	book3 := &Books{"Java 教程", "www.java.com", "Java 语言教程", 6412300}
	printBook1(book3)

	var book4 = Books{"rust 教程", "www.rust.com", "rust 语言教程", 2412370}
	printBook(book4)

	//结构体转换
	book5 := ebook{"e 教程", "www.e.com", "e 语言教程", 2846270}
	book6 := Books(book5)
	printBook(book6)
}
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook1(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
