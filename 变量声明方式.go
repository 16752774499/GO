package main

import (
	"fmt"
)

func main() {
	//方法一、声明int变量，不给默认值则为0
	var a int
	fmt.Println(a)
	fmt.Printf("a的类型是：%T\n", a)
	//方法二、声明int变量，并赋值
	var b int = 10
	fmt.Println(b)
	fmt.Printf("b的类型是：%T\n", b)
	//方法三、在初始化时，可以省去数据类型，通过值自动推导
	var c = "121212"
	fmt.Println(c)
	fmt.Printf("c的类型是：%T\n", c)

	//方法四、简写形式(推荐)省去var关键字，直接自动匹配----只能在局部变量中使用
	d := b
	fmt.Println(d)
	fmt.Printf("d的类型是：%T\n", d)
}
