package main

import (
	"errors"
	"fmt"
	"reflect"
)

var a int = 32

const b = iota
const c = iota
const (
	d = iota + 2
	e
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
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

	//调用异常
	handleException()
	//获取长度、容量
	getLen()
	//调用结构体
	TestForDog()
	Hello()
}

//创建切片，相当于java中的数组
func makeSlice() {
	strings := make([]string, 3)
	strings[0] = "第一位"
	strings[1] = "第二位"
	strings[2] = "第三位"
	fmt.Println(strings)
	i := make([]string, 5)
	i[1] = "这是第一位"
	fmt.Println(i, "这是新声明的切片")
}

//创建map，相当于java中的map，具有key和value的键值对
func makeMap() {
	m := make(map[int]string, 5)
	m[11] = "11"
	m[22] = "22"
	m[22] = "33"
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
	fmt.Println("newMap", reflect.TypeOf(newMap))
	//引用类型
	fmt.Println("makeMap", reflect.TypeOf(makeMap))
}

//调用slice的append方法，对切片进行扩充
func appendElementForSlice() {
	strings := make([]string, 2)
	strings[0] = "id-1"
	strings[1] = "id-2"
	fmt.Println("切片扩容之前的长度为：", len(strings))
	strings = append(strings, "id-3")

	fmt.Println(strings)
	fmt.Println("切片扩容之后的长度为：", len(strings))
}

//复制切片内容
func copySlice() {
	sliceSource := make([]string, 3)
	sliceSource[0] = "0"
	sliceSource[1] = "1"
	sliceSource[2] = "2"
	sliceTarget := make([]string, 3)
	copy(sliceTarget, sliceSource)
	fmt.Println(sliceTarget)
	fmt.Println(reflect.TypeOf(sliceTarget))
	fmt.Println(reflect.TypeOf(sliceSource))
}

func deleteMap() {
	m := make(map[int]string)
	m[1] = "1"
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)
}

/**
异常处理：panic   recover
*/

func handleException() {
	//捕获异常信息，保证程序的顺利执行
	/*defer func() {
		message := recover()
		fmt.Println("异常信息为：",message)
	}()*/
	defer coverPanic()
	//panic("I am panic")
	panic(errors.New("this is an error"))
}

func coverPanic() {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message", message)
	case error:
		fmt.Println("error message", message)
	default:
		fmt.Println("unknown message")

	}

}

/**
测试长度和容量
*/
func getLen() {
	//第三个参数为容量大小，初始化容量，如果容量不够，会自动扩充
	sliceSource := make([]string, 2, 5)
	sliceSource[0] = "位置1"
	sliceSource[1] = "位置2"
	//赋值报错，超出长度范围
	//sliceSource[2]="位置3"
	//但是可以使用append的方法追加元素
	sliceSource = append(sliceSource, "位置3")
	fmt.Println("sliceSource的长度为：", len(sliceSource))
	fmt.Println("sliceSource的容量为：", cap(sliceSource))
}

/**
测试关闭chan
*/
func closeChan() {
	//创建chan
	mChan := make(chan int, 1)
	//往chan里面填充数据
	mChan <- 1
	//关闭chan
	close(mChan)
}

//定义结构体，多个数据类型的集合，类似于java中的类对象
type Dog struct {
	ID   int
	Name string
	Age  int
}

//使用结构体
func TestForDog() {
	//创建方式一
	var dog Dog
	dog.Age = 10
	dog.ID = 1
	dog.Name = "旺财"
	//创建方式2
	dog2 := Dog{ID: 1, Name: "花花", Age: 10}
	//创建方式3，指针类型
	dog3 := new(Dog)
	dog3.ID = 2
	dog3.Name = "小白"
	dog3.Age = 10
	fmt.Println("dog的结构体为：", dog)
	fmt.Println("dog2的结构体为：", dog2)
	fmt.Println("dog3的结构体为：", dog3)
}
