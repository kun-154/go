package main

import "fmt"

// 结构体的指针

type animal struct {
    name string
}

func (a *animal) move() {
    fmt.Printf("%smove", a.name)
}

type dog struct {
    feet int8
    *animal //  匿名嵌套， 并且是结构体指针
}

func (d *dog)say() {
    fmt.Printf("%s iii", d.name)
}


func main() {
    d1 := dog {
        feet: 4,
        animal: &animal{
            name: "ll",
        },
    }
    d1.move();
    d1.say();



}
