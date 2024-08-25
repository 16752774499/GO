package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"` //在其他包导用时，可获取标签内容，判断该字段的作用
	Age  int    `info:"age"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() //获取结构体的类型
	//v := reflect.ValueOf(str)       //获取结构体的值
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info") //获取结构体字段的名称
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Printf("info：%s\n", tagInfo)
		fmt.Printf("doc：%s\n", tagDoc)

	}
}

func main() {

	re := resume{
		Name: "张三",
		Age:  18,
	}

	findTag(&re)

}
