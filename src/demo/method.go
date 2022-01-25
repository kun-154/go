package main

import "fmt"


type Person struct {
    name string 
    age int8
}

// Person 的构造 函数
func NewPerson(name string, age int8) *Person{
    return &Person{
        name: name,
        age: age,
    }
}

//Dream 是为person 类型定义的方法
func (p Person)Dream() {
    fmt.Printf("gogogo")
}


// 指针类型的接受者
// 指针类型接受者  就是指 接受者的类型 为指针类型
func (p *Person)SetAge(newAge int8) {
    p.age = newAge
}

// 值类型的接受者
// 值类型接受者  就是指 接受者的类型 为值类型
// 当想要改变 年龄的时候 不能使用 值接受者， 必须使用 指针接受者 
func (p Person)SetAge2(newAge int8) {
    p.age = newAge
}

// 方法  和接收 者 示例
func main() {
    p1 := NewPerson("df", 34)
    (*p1).Dream()
    p1.Dream()
    p1.SetAge(18)
    fmt.Println(p1.age)

    p1.SetAge2(20)
    fmt.Println(p1.age)

}
