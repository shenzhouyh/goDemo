package main

import "fmt"
var a int = 32
const b = iota
const c = iota
const(
	d=iota+2
	e
)
func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

fmt.Println("hello world")
	makeSlice()
}
//创建切片
func makeSlice(){
	strings := make([]string, 3)
	strings[0]="第一位"
	strings[1]="第二位"
	strings[2]="第三位"
	fmt.Println(strings)
}
