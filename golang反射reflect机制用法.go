package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string
	Age  int
}

func (this *MyStruct) getName() {
	fmt.Println("name:", this.Name)
}

func Reflect(arg interface{}) {
	fmt.Printf("arg type is %v\n", reflect.TypeOf(arg))
	fmt.Printf("arg value is %v\n", reflect.ValueOf(arg))

}

func main() {

	m := MyStruct{"xaioHua", 18}

	//var a float64 = 10.2

	//Reflect(m)
	DoFiledAndMethod(m)

	//type_ := reflect.TypeOf(m) //a是reflect.Type类型, 类型是MyStruct
	//
	//value_ := reflect.ValueOf(m)
	//
	//fmt.Printf("%v\n", type_)
	//fmt.Printf("%v\n", value_)

}

func DoFiledAndMethod(arg interface{}) {
	//获取arg的type
	type_ := reflect.TypeOf(arg)
	fmt.Println("type is ", type_)
	//获取arg的value
	value_ := reflect.ValueOf(arg)
	fmt.Println("value is ", value_)

	//通过type获取里边的字段
	//for i := 0; i < type_.NumField(); i++ {
	//	filde := type_.Field(i)
	//	value := value_.Field(i)
	//	fmt.Println("field name is ", filde.Name)
	//	fmt.Println("field value is ", value)
	//	fmt.Println("field type is ", filde.Type)
	//}
	//通过type获取里边的方法
	for i := 0; i < type_.NumMethod(); i++ {
		method := type_.Method(i)
		fmt.Printf("method name is %s\n", method.Name)
	}

}
