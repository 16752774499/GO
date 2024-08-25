package main

import "fmt"

func printMap(m map[string]string) {
	// 引用传递
	m["广州"] = "GuangZhou"
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
}

func main() {
	cityMap := make(map[string]string)
	cityMap["北京"] = "BeiJing"
	cityMap["上海"] = "ShangHai"
	cityMap["深圳"] = "ShenZhen"

	// 删除key
	delete(cityMap, "深圳")
	// 修改key对应的value
	cityMap["北京"] = "BeiJing1111111111111"

	// 遍历map
	printMap(cityMap)
	for key, value := range cityMap {
		fmt.Println("key:", key, "value:", value)
	}

}
