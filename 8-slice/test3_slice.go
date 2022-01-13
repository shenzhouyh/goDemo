package main

import "fmt"

func main() {
	var slice0 []int
	//方式一
	//声明slice1是一个切片，但是并没有给slice1分配空间
	var slice1 []int
	//给slice1分配空间,开辟了3个空间，默认值为0
	slice1 = make([]int, 3)
	length := len(slice1)
	fmt.Println("the length of slice1 is", length)
	//方式二 通过冒等进行赋值
	slice2 := make([]int, 3)
	length2 := len(slice2)
	fmt.Println("the length of slice2 is", length2)

	//判断一个slice是否为0
	if slice0 == nil {
		fmt.Println("slice1 是一个空的切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
