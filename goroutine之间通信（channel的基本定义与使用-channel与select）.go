package main

//channel与select(单流程下一个go只能监控一个channel的状态，select可以完成监控多个channel的状态)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case v := <-ch1:
				println("ch1:", v)
			case v := <-ch2:
				println("ch2:", v)
			case ch1 <- 1:
				println("ch1")
			case ch2 <- 2:
				println("ch2")
			default:
				//println("default")
			}
		}
	}()
	ch1 <- 1
	ch1 <- 3
	ch2 <- 2
	ch2 <- 4
	v := <-ch1
	v2 := <-ch2
	println(v, v2)
}
