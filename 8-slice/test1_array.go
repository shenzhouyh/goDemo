package main

import "fmt"

func main() {
	//1.声明固定长度的数组
	var myarray1 [10]int
	//遍历输出数组内容
	for i := 1; i < len(myarray1); i++ {
		//默认值均为0
		fmt.Println(myarray1[i])
	}
	//2.声明并初始化部分数组
	myarray2 := [10]int{0, 1, 2, 3, 4}

	for index, value := range myarray2 {
		//未赋值的默认值仍为0
		fmt.Println("index=", index, "  vaule=", value)
	}
	//3.声明数组并且初始化
	myarray3 := [4]int{1, 2, 3, 4}
	fmt.Printf("myarray3的地址为：%p", &myarray3)
	fmt.Println()
	printArrayAndChange(myarray3)
	//myarray3的地址和形参的myarray3地址不同
	for i := range myarray3 {
		//在printArrayAndChange中更改了myarray中的值，原始的myarray并没有受到影响，所以，对应的也不是同一个数组信息
		fmt.Println("index=", i, "  value=", myarray3[i])
	}
	//查看数组的数据类型
	fmt.Printf("myarray1 types=%T\n", myarray1)
	fmt.Printf("myarray2 types=%T\n", myarray2)
	fmt.Printf("myarray3 types=%T\n", myarray3)
}

func printArrayAndChange(myarray3 [4]int) {
	fmt.Printf("传入的形参myarray3的地址为：%p", &myarray3)
	for _, value := range myarray3 {
		fmt.Println("对应的数组值为：", value)
	}
	myarray3[0] = 100
	myarray3[1] = 200
	myarray3[2] = 300

}
