package main

import (
	"fmt"
	"reflect"
)

func main() {
	/**
	使用make创建slice、map、chan
	*/
	makeSlice()
	makeMap()
	makeChan()
	//使用new创建map，带有指针
	newMap()

	appendElementForSlice()

	copySlice()

	deleteMap()
}

//创建切片，相当于java中的数组
func makeSlice(){
	strings := make([]string, 3)
	strings[0]="第一位"
	strings[1]="第二位"
	strings[2]="第三位"
	fmt.Println(strings)
	i := make([]string, 5)
	i[1]="这是第一位";
	fmt.Println(i,"这是新声明的切片")
}
//创建map，相当于java中的map，具有key和value的键值对
func makeMap()  {
	m := make(map[int]string, 5)
	m[11]="11"
	m[22]="22"
	m[22]="33"
	fmt.Println(m)
}
//创建chan，管道，可以向里面写数据
func makeChan() {
	ints := make(chan int)
	//关闭管道
	close(ints)
}
//使用new的方式创建map
func newMap() {
	newMap := new(map[int]string)
	makeMap := make(map[int]string)
	//指针类型
	fmt.Println("newMap",reflect.TypeOf(newMap))
	//引用类型
	fmt.Println("makeMap",reflect.TypeOf(makeMap))
}
//调用slice的append方法，对切片进行扩充
func appendElementForSlice()  {
	strings := make([]string, 2)
	strings[0]="id-1"
	strings[1]="id-2"
	fmt.Println("切片扩容之前的长度为：",len(strings))
	strings = append(strings, "id-3")

	fmt.Println(strings)
	fmt.Println("切片扩容之后的长度为：",len(strings))
}
//复制切片内容
func copySlice() {
	sliceSource :=make([]string,3)
	sliceSource[0]="0"
	sliceSource[1]="1"
	sliceSource[2]="2"
	sliceTarget :=make([]string,3)
	copy(sliceTarget,sliceSource)
	fmt.Println(sliceTarget)
	fmt.Println(reflect.TypeOf(sliceTarget))
	fmt.Println(reflect.TypeOf(sliceSource))
}

func deleteMap() {
	m := make(map[int]string)
	m[1]="1"
	fmt.Println(m)
	delete(m,1)
	fmt.Println(m)
}

