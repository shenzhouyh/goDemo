package main

import "fmt"

/**
动态数组在作为参数传递时，是引用传递
而固定数组作为参数传递时，是值传递
*/
func main() {
	//声明一个动态数组
	myarray := []int{1, 2, 3, 4}

	fmt.Printf("myarray type is %T\n", myarray)
	printArray(myarray)
	fmt.Println("=======")
	//再次输出数组信息
	for i, value := range myarray {
		//myarray[0]=100,表明函数的动态数组传递是引用传递
		fmt.Println("index= ", i, " value= ", value)
	}
}

func printArray(myarray []int) {
	//数组的引用传递
	for _, value := range myarray {
		fmt.Println("value= ", value)
	}
	myarray[0] = 100
}
