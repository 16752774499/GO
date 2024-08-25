package lib1

import "fmt"

// 大写开头表示导出
func LibTest1() {
	fmt.Println("Func lib1")
}

func init() {
	fmt.Println("init lib1")
}
