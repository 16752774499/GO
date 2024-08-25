package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	// 截取切片 [0, 2)
	s1 := s[0:2] // 与s是同一地址

	s[0] = 111111

	fmt.Printf("s1:%v len:%v cap:%v\n", s1, len(s1), cap(s1))

	// copy 可以将数组直接copy一份
	s2 := make([]int, 3)

	copy(s2, s)

	s[0] = 666666666666666

	fmt.Printf("s2:%v len:%v cap:%v\n", s2, len(s2), cap(s2))

}
