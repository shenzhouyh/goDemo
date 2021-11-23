package main

import (
	"fmt"
	"reflect"
)

/**
变量声明的四种方式
*/
func main() {
	// 第一种：
	var name string = "你好"
	fmt.Printf("value of name = %s type of name=%T\n", name, name)
	fmt.Println("name的数据类型为：", reflect.TypeOf(name))
	fmt.Println(name)

	//第二种：冒等
	//只能在方法体中使用，不可以用于声明全局变量
	age := 21
	fmt.Println("age=", age)

	//第三种：隐式赋值
	var size int
	fmt.Println("size的默认值为：", size)

	//第四种：省去数据类型
	var length = 10
	fmt.Println("length的长度是：", length)

	//多变量声明

	var a, b, c = 100, 200, 300
	fmt.Println("a=", a, "b=", b, "c=", c)
	var (
		e = 100
		f = 200
		g = 300
	)

	fmt.Println("e=", e, "f=", f, "g=", g)
}
