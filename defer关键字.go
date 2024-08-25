package main

import "fmt"

/*
func main() {
	// 压栈的顺序（执行按出栈）
	defer fmt.Println("hello world defer end111111111")
	defer fmt.Println("hello world defer end222222222")

	fmt.Println("hello world 11111111111")

	fmt.Println("hello world 22222222222")

}
*/

func deferTest() int {
	fmt.Println("deferTest ----------------")
	return 0
}

func returnTest() int {
	fmt.Println("returnTest ----------------")
	return 0
}

func returnAndDefer() int {
	defer deferTest()
	return returnTest() //return先于defer执行
}

func main() {
	returnAndDefer()
}
