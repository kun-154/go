package main

import "fmt"

// 为任意类型添加方法


/*
不能 为其他 不是自己定义的类型 定义方法
func (i int) say() {
    fmt.Println("hi")
}
*/

// 但是 可以 基于内置的基本类型造一个我们自己的类型 

type myInt int

func (i myInt) say() {
    fmt.Println("hi")
}

func main() {
    var n myInt
    n.say();
}
