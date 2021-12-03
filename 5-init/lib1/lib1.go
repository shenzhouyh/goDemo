package lib1

import "fmt"

func LibTest() {
	fmt.Print("lib1Test()")
}

func init() {
	fmt.Println("调用lib1的init方法")
}
