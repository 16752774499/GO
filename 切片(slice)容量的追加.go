package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)

	// 追加元素
	slice = append(slice, 1)
	slice = append(slice, 2)

	fmt.Printf("len(长度)=%d cap(容量)=%d slice=%v\n", len(slice), cap(slice), slice)

	// 切片扩容机制，append时，如果容量不够，则容量增加两倍
	slice = append(slice, 3)
	//fmt.Printf("len(长度)=%d cap(容量)=%d slice=%v\n", len(slice), cap(slice), slice)

	num2 := make([]int, 2)
	fmt.Printf("len(长度)=%d cap(容量)=%d slice=%v\n", len(num2), cap(num2), num2)
	num2 = append(num2, 1)
	num2 = append(num2, 2)
	fmt.Printf("len(长度)=%d cap(容量)=%d slice=%v\n", len(num2), cap(num2), num2)
	num2 = append(num2, 3)
	fmt.Printf("len(长度)=%d cap(容量)=%d slice=%v\n", len(num2), cap(num2), num2)

}
