package main

import (
	"fmt"
	"time"
)

/****************************************************
// 协程
调度器的设计策略
1.复用线程
2.利用并向
3.抢占
4.全局G队列
*/
//从
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("task:%d\n", i)
		time.Sleep(1 * time.Second)
	}

}

// 主
func main() {
	// 创建一个go程，去执行newTask函数
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main:%d\n", i)
		time.Sleep(1 * time.Second)
	}

}
