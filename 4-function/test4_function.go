package main

import "fmt"

/**
创建多种函数
*/

// 函数1：无参数无返回值
func test1() {
	fmt.Println("test1 这是一个无返回值， 无参数的函数")
}

//函数2：有参数，无返回值
func test2(name string) {
	fmt.Printf("test2 这是一个有参数无返回值的函数，对应的参数为 name = %v\n", name)
}

//函数3：有参数，有返回值
func test3(a string, b string) int {
	if len(a) > 0 && len(b) > 0 {
		return 10
	} else {
		return 0
	}
}

// 函数4：有参数，多返回值
func main() {
	test1()
	test2("学习GoLang")
	result := test3("a", "b")
	fmt.Printf("test3 这是一个有参数有返回值的函数，对应的返回值为：%v", result)

}
