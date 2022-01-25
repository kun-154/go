package main

import "fmt"

// 嵌套结构体 

type Address struct {
    province string
    city string
}


type Person struct {
    name string
    age int
    address Address
}

type Person1 struct {
    name string
    age int
    Address // 匿名 结构体， 只用写 结构体名称， 不需要命名
}

func main() {
    p1 := Person{
        name: "df",
        age: 18,
        address:Address{
            province:"dd",
            city:"df",
        },
    }

    fmt.Printf("%#v\n", p1)
    fmt.Println(p1.name, p1.age, p1.address)

    p2 := Person1{
        name: "df",
        age: 18,
        Address: Address{
            province:"dd",
            city:"df",
        },
    }

    fmt.Printf("%#v\n", p2);
    fmt.Println(p2.name, p2.age, p2.Address,province) // 通过嵌套的匿名结构体字段访问其内部的字段
    fmt.Println(p2.name, p2.age, p2.province) // 直接访问匿名结构体中的字段
}
