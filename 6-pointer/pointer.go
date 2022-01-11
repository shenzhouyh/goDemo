package main

import "fmt"

/**
指针

*/
func main() {
	var a int = 10
	var b int = 20

	var c int
	fmt.Println("c的初始地址为：", &c)
	c = a
	fmt.Println("赋值之后c的地址为：", &c)
	fmt.Println("c的值为：", c)
	//以上可以说明，go中赋值是值传递
	swap(&a, &b)
	fmt.Println("a=", a, " b=", b)

	var p *int
	p = &a
	fmt.Println("&a的值为：", &a)
	fmt.Println(" p的值为：", p)
	// 二级指针
	var pp **int
	//将指针的地址赋值为pp
	pp = &p

	fmt.Println("&p的值为：", &p)

	fmt.Println("pp的值为：", pp)

}

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}
