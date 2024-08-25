package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func (this *Dog) Eat() {
	fmt.Println("Dog Eat")
}

func (this *Dog) Run() {
	fmt.Println("Dog Run")
}

type SupperDog struct {
	Dog   // 继承(SupperDog 继承 Dog)
	level int
}

func (this *SupperDog) Eat() {
	fmt.Println("Super Dog Eat")
}

func (this *SupperDog) Show() {
	fmt.Printf("name:%s,age:%d,level:%d\n", this.name, this.age, this.level)
}

func main() {
	dog := Dog{"dog", 1}
	dog.Eat()
	//定义一个子类对象
	//Sdog := SupperDog{Dog{"sdog", 2}, 1}
	var Sdog SupperDog
	Sdog.name = "sdog"
	Sdog.age = 2
	Sdog.level = 1

	Sdog.Eat()
	Sdog.Run()
	Sdog.Show()

}
