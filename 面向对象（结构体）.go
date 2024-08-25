package main

import "fmt"

// 声明一种新的数据类型，是int的别名
type myInt int

// 结构体
type Book struct {
	title  string
	author string
}

func printBook(book Book) {
	// 传递一个Book的副本（传递）
	book.title = "Go"
	fmt.Printf("%v\n", book)
}

func PrintBook(book *Book) {
	// 传递一个Book的指针（引用）
	book.title = "Go"
	fmt.Printf("%v\n", book)
}

func main() {
	var a myInt = 10
	println(a)

	var myBook Book
	myBook.title = "Go语言"
	myBook.author = "Go语言中文网"

	printBook(myBook)
	//fmt.Printf("%v\n", myBook)

	fmt.Printf("%v\n", myBook)

	PrintBook(&myBook)

}
