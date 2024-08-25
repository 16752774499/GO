package main

import "fmt"

type Book struct {
	author string
}

// interface{}表示万能类型
func myFunc(arg interface{}) {
	fmt.Println("MyFunc ----")
	fmt.Printf("%v\n", arg)
	fmt.Printf("参数类型：%T\n", arg)
	//如何区分interface{}此时引用的到底是什么数据类型
	//golang给 interface{}提供“断言”的机制

	value, ok := arg.(string)
	if !ok {
		fmt.Printf("arg is not string\n")
	} else {
		fmt.Printf("arg is string,value = ‘%s’\n", value)
	}

}

func main() {

	book := Book{"xaiohua"}
	myFunc(book)
	myFunc("1222")
	//myFunc(1222)
	//myFunc(12.22)

}
