package main

import (
	_ "package_test/lib1" // 匿名导包，仅用于初始化包。可以在不适用该包内的接口下导入

	mylib2 "package_test/lib2" // 导包别名
	. "package_test/lib3"      //将该包内的接口导入到当前包，相当于在当前包中定义了该包内的接口
)

func main() {

	//lib1.LibTest1()
	mylib2.LibTest2()

	LibTest3()
}
