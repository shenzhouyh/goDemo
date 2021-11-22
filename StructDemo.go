package main

import "fmt"

type Cat struct {
	ID   int
	Name string
	age  int
}

func Hello() {
	var cat Cat
	cat.age = 12
	cat.ID = 1
	cat.Name = "旺财"
	fmt.Println("hello", cat)
}

func main() {
	fmt.Printf("hello world")
}
