package main

import ("fmt"
    "reflect"
)

func reflectType(x interface{}) {
    //接收空接口类型（任意类型的数据）
    // 此时， 我们不知道传入的变量 的类型
    // 判断类型 1. 通过类型断言  2. 反射工具
    t := reflect.TypeOf(x)
    fmt.Println(t, t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
    v := reflect.ValueOf(x)
    fmt.Printf("%v, %T\n", v, v) 
}

type person struct{}

func (p person) say() {
    fmt.Println("say")
}

type sayer interface{
    say()
}

func main() {
    /*
    var a int
    reflectType(a)
    var p = person{}
    reflectType(p)
    var ser sayer
    //reflectType(ser)
    ser = p
    ser.say()
    reflectType(ser)
    */


}
