package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor()
	GetType()
}

// 具体的类
type Cat struct {
	color string
	type_ string
}

func (this *Cat) Sleep() {
	println("猫睡觉")
}
func (this *Cat) GetColor() {
	fmt.Printf("猫的颜色是%s\n", this.color)
}
func (this *Cat) GetType() {

	println("猫是猫")
}

type Dog struct {
	color string
	type_ string
}

func (this *Dog) Sleep() {
	println("狗睡觉")
}
func (this *Dog) GetColor() {
	fmt.Printf("狗的颜色是%s\n", this.color)
}
func (this *Dog) GetType() {
	println("狗是狗")
}

func AnimalIFSleep(this AnimalIF) {
	this.Sleep() //多态
}

func main() {
	var animal, animal1 AnimalIF // 声明一个接口(本质是一个指针)
	animal = &Cat{color: "黄色", type_: "猫"}
	//animal.Sleep()
	//animal.GetColor()
	//animal.GetType()
	animal1 = &Dog{color: "黑色", type_: "狗"}
	//animal1.Sleep()
	//animal1.GetColor()
	//animal1.GetType()
	AnimalIFSleep(animal)
	AnimalIFSleep(animal1)
}
