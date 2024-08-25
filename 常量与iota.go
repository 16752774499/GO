package main

import "fmt"

// const定义枚举类型
// iota只能出现在const中
const (
	//可以在const（）添加关键字iota,iota默认从0开始,每行增加1

	shanxi = 10*iota + 1 //
	beijing
	shandong
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f = iota * 2, iota * 3
	g, h
)

func name(i int) {
	fmt.Println(i)
}

func main() {
	// 常量（只读）
	const a int = 10
	fmt.Println(a)

	fmt.Println(shanxi)
	fmt.Println(shandong)
	fmt.Println(beijing)
}
