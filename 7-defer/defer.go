package main

import "fmt"

/**
  defer 关键字
*/
func main() {
	defer fmt.Println("defer main end1")

	defer fmt.Println("defer main end2")

	fmt.Println("main end1")

	fmt.Println("main end2")
}
