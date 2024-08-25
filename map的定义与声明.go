package main

import "fmt"

func main() {
	// 第一种声明方式
	//声明myMap是一个一个map，key为string,value为string
	var myMap map[string]string
	// 开辟空间
	myMap = make(map[string]string, 1)

	if myMap == nil {
		fmt.Println("map is nil")
	} else {
		fmt.Println("map is not nil")
		myMap["name"] = "张三"
		myMap["age"] = "20"
		fmt.Printf("len:%d\n", len(myMap))
	}
	// 第二种声明方式
	myMap2 := make(map[string]int)
	myMap2["name"] = 1
	myMap2["age"] = 20
	fmt.Println(myMap2)

	// 第三种声明方式
	dict := map[string]string{
		"name": "张三",
		"age":  "20",
		"sex":  "男",
	}
	fmt.Println(dict)

	fmt.Scanln()

}
