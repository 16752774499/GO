package main

import "fmt"

func main() {
	// 切片 声明slice是一个切片，并初始化，默认值为10个元素，长度为10
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 切片声明 但是不初始化，默认值为0个元素，长度为0
	//var slice []int
	//// 通过make开辟空间，初始化slice
	//slice = make([]int, 10)
	//slice[0] = 1

	// 声明切片，同时给slice分配空间，初始化slice，初始值为0
	//var slice []int = make([]int, 5)
	slice := make([]int, 5)

	fmt.Printf("slice:%v\n", slice)

	// 判断slice是否为空
	if slice != nil {
		println("slice is not nil")
	} else {
		println("slice is nil")
	}

}
