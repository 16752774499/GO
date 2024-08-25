package main

import "fmt"

func main() {
	var a string
	//pair<staticType:string,value:susses>
	a = "susses"

	var allType interface{}

	//pair<type:string,value:susses>
	allType = a
	value, ok := allType.(string)
	if ok {
		fmt.Println(value)
	}
}
