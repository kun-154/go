package main

import "fmt"

type dog struct {
    name string
}
type cat struct {}
type person struct {}

func (d dog) say() {
    fmt.Println(d.name, "dog say")
}
func (p person) sqy() {
    fmt.Printlb("man say")
}

func (c cat) say() {
    fmt.Println("car say")
}

// 定义一个类型， 一个抽象的类型， 只要实现了 say（）这个方法的类型都可以成为sayer类型
type sayer interface {
    say()
}

// 接口 不管你是什么类型， 他只管你要实现的动作（类似于 多态）
func do(arg sayer) {
    arg.say()
}

func main() {
   /* c1 := cat{
    }
    do(c1)
    d1 := dog{
        name: "hank",
    }
    d2 := dog{
        name: "hhhh",
    }
    do(d1)
    do(d2)
*/

    var s sayer
    c2 := cat{}
    s = c2
    d2 := dog{}
    s = d2
    /*
    p3 := person{}
    s = p3

    只要定义了 接口中 函数的 结构体 都可以 被 interface 结构体接收
    反之 则不行
    */
    fmt.Println(s)
}
