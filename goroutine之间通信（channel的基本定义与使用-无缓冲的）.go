package main

import "fmt"

func main() {
	//定义一个chanel 无缓冲
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束") //defer 会在num := <-c 执行完毕后执行
		fmt.Println("goroutine 开始")
		//将666写入c中
		c <- 666 //阻塞

	}()
	//从c中读取数据,并赋值给num
	num := <-c //阻塞
	fmt.Println(num)
	fmt.Println("main 结束")
}
