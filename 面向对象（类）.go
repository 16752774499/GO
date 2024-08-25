package main

import "fmt"

// 类名首字母大写，表示该类可以被其他包访问，否则只能在本包中访问
type Person struct {
	Name string // 首字母大写，表示该字段可以被其他包访问
	age  int
}

/*
func (this Person) getName() string {
	//fmt.Println("Hello,", this.Name)
	return this.Name
}
func (this Person) setName(newName string) {
	// 当前this是调用该方法的对象的一个副本
	this.Name = newName
}

func (this Person) isShow() {

	fmt.Printf("%v\n", this)
}
*/

func (this *Person) getName() string {
	//fmt.Println("Hello,", this.Name)
	return this.Name
}
func (this *Person) setName(newName string) {
	// 当前this是调用该方法的对象的一个副本
	this.Name = newName
}

func (this *Person) isShow() {

	fmt.Printf("%v\n", *this)
}

func main() {
	p := Person{"Tom", 18}
	p.setName("Jerry")
	p.isShow()
	fmt.Println(p.age)
}
