package main

import "fmt"

/**
  const 用于定义枚举类型
	1、iota的从0开始，逐行递增，中间若有中断
	2、iota只能在const()中使用，不能单独使用，可用于进行累计
*/

const (
	BEIJING  = 10 * iota //iota =0
	TIANJIN  = 1         //iota =1
	SHANGHAI             //iota =2
	SHENZHEN = 10 * iota //iota =3
)
const name string = "GoLang"

func main() {
	fmt.Println("BEIJING", BEIJING)
	fmt.Println("TIANJIN", TIANJIN)
	fmt.Println("SHANGHAI", SHANGHAI)
	fmt.Println("SHENZHEN", SHENZHEN)
	fmt.Println("name is ", name)
}
