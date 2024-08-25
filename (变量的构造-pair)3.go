package main

import "fmt"

type Reader interface {
	ReaderBook()
}
type Writer interface {
	WriterBook()
}
type Book struct {
}

func (this *Book) ReaderBook() {
	fmt.Println("ReaderBook")
}

func (this *Book) WriterBook() {
	fmt.Println("WriterBook")
}

func main() {
	b := &Book{}
	var r Reader = b

	// 反射, 接口的动态类型和动态值
	r.ReaderBook()

	value, ok := r.(Writer) //此处的断言为什么会成功呢？（因为w,r的type是一致的）
	if ok {
		println(value)

		value.WriterBook() //此时的value是*Book类型，所以可以调用WriterBook方法
	}

}
