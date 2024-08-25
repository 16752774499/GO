package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	//channel与rabge
	//range ch 循环读取channel中的数据
	for data := range ch {
		fmt.Println(data)
	}

	//for循环读取channel中的数据
	for {
		//ok如果为true表示channel有数据，否则为false表示已经关闭
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}

	}
	fmt.Println("Main end")
}
