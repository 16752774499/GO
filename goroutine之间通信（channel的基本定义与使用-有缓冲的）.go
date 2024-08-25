package main

import (
	"fmt"
	"time"
)

func main() {
	//有缓冲的
	ch := make(chan int, 3)                               //创建了一个长度为3的int类型的channel
	fmt.Println("len(ch):", len(ch), "cap(ch):", cap(ch)) //len表示当前元素的数量，cap表示channel的容量

	go func() {
		defer fmt.Println("子go程 end")
		for i := 0; i < 4; i++ { //当channel已经满了的时候，向channel中写入数据会阻塞
			ch <- i                                                                      //向channel中写入数据
			fmt.Println("子go程正在运行，len(ch)=", len(ch), "cap(ch)=", cap(ch), "当前发送元素：", i) //从channel中读取数据
		}
	}()
	time.Sleep(2 * time.Second)
	//
	for i := 0; i < 4; i++ { //当channel为空的时候，从channel中读取数据会阻塞
		num := <-ch                                                                              //从channel中读取数据
		fmt.Println("main goroutine正在运行，len(ch)=", len(ch), "cap(ch)=", cap(ch), "当前读取元素：", num) //向channel
	}

	fmt.Println("main goroutine end")
}
