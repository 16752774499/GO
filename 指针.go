package main

import "fmt"

// 值传递
/*
func changeValue(p int) {
	p = 10
}

func main() {
	var a int = 1
	changeValue(a)
	print(a)
}
*/

// 指针传递
/*
func changeValue(p *int) {
	*p = 10
	fmt.Printf("p = %T\n", p)
	println(&p)
	var b **int = &p
	fmt.Printf("b = %T\n", b)

	println(b)
}

func main() {
	var a int = 1
	changeValue(&a)
	log.Println(a)
	//println(&a)
}
*/
func main() {
	var a, b int = 100, 200
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(i1 *int, i2 *int) {
	temp := *i1
	*i1 = *i2
	*i2 = temp
}
