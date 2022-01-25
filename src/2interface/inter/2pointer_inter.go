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

func (p person) say() {
    fmt.Println("man say")
}

// 值接受者 实现接口， 类型的指针和类型的值 都能保存到接口变量中 
func (c cat) say() {
    fmt.Println("car say")
}

func (d dog) move() {
    fmt.Println( "dog say")
}

func (c cat) move() {
    fmt.Println(c.name, "cat say")
}

func (p person) move() {
    fmt.Println(p.name, "person say")
}
// 指针接受者 实现接口， 只能接受 类型的指针 保存到接口变量中 
/*
func (c *cat) say() {
    fmt.Println("car say")
}
*/

// 定义一个类型， 一个抽象的类型， 只要实现了 say（）这个方法的类型都可以成为sayer类型

type sayer interface {
    say()
}

type mover interface {
    move()
}

// 接口的 嵌套 
type animal interface {
    sayer
    mover
}

func main() {
/*
    var s sayer
    c1 := cat{}
    s = c1
    s.say()
    c2 := &cat{}
    s = c2
    s.say()
    /*
    p3 := person{}
    s = p3
    */

    //只有定义了 接口中 函数的 结构体 才可以 被 interface 结构体接收
    //cat dog 都实现了say  和move  所以 cat 和dog 就可以成为 animal类接受的对象
    var a animal
    c1 := cat{}
    d1 := dog{}

    a = c1
    a.say()
    a.move()


}
