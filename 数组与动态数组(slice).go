package main

import "fmt"

// 数组
// 数组是值类型，是值传递的
/*
func main() {
	//固定长度数组
	var myArray [10]int
	for i := 0; i < len(myArray); i++ {
		myArray[i] = i
		println(myArray[i])
	}

	MyArray2 := [5]int{1, 2, 3, 4}
	for index, val := range MyArray2 {
		println(index, val)
	}
	//查看数组的数据类型
	fmt.Printf("%T\n", MyArray2)
	fmt.Printf("%T\n", MyArray2)
}
*/

// 切片 此时函数传参为引用传递
func printArray(array []int) {
	// _表示匿名变量 忽略索引
	for _, val := range array {
		fmt.Printf("val:%d\n", val)
	}
	array[0] = 10000
}

func main() {
	//动态数组
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", array)
	printArray(array)
	println(array[0])
}
