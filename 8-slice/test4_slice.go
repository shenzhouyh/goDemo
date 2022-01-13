package main

import "fmt"

func main() {
	//声明一个slice，int类型，长度为3，容量为5
	var numbers = make([]int, 3, 5)
	//len=3,cap=5,[0,0,0]
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
}
