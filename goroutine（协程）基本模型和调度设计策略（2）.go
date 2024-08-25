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

// 主
func main() {

	//go func() {
	//	defer fmt.Println("A.defer")
	//
	//	func() {
	//		defer fmt.Println("B.defer")
	//
	//		// 退出当前go程
	//		runtime.Goexit()
	//
	//		fmt.Println("B")
	//	}()
	//
	//	fmt.Println("A")
	//}()

	go func(a int, b int) bool {
		if a == b {
			fmt.Println("a == b")
			return true
		} else {
			fmt.Println("a != b")
			return false
		}
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}

}
