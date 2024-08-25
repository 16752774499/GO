package main

import "fmt"

func aaa(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 999
	return c
}

// 多返回值(匿名)
func bbb(a string, b string) (int, int, int) {

	fmt.Println(a)
	fmt.Println(b)
	c := 999
	d := 998
	e := 997
	return c, d, e
}

// 多返回值(命名)
func ccc(a string, b string) (r1 int, r2 int) {
	fmt.Println(a)
	fmt.Println(b)
	r1 = 999
	r2 = 998
	return r1, r2
}

// 多返回值(命名+匿名)
func ddd(a string, b string) (r1, r2 int) {
	//fmt.Println(a)
	//fmt.Println(b)
	// r1,r2 是属于ddd函数的局部变量
	fmt.Println(r1)
	fmt.Println(r2)
	r1 = 999
	r2 = 998
	return r1, r2
}

func init() {
	fmt.Println("init")
}

func main() {
	//fmt.Println(bbb("hello", "world"))
	ret1, ret2 := ccc("hello", "world")
	fmt.Println(ret1, ret2)
}
