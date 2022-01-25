package main

import "fmt"

// 空接口 ： 接口中 没有定位任何需要实现的方法时，该接口就是一个空接口
// 故 空接口都实现了 空接口 ： 空接口 可以存储 任何值
type xxx interface {}

// 空接口的应用  
 // 1. 空接口类型  作为函数的参数

/*
func Println(a ...interface{}) (n int, err error) {
    reutrn........
}
上面是 打印函数的函数原型 ， 使用了 空接口 承接任意参数
*/

/*
2.  空接口 可以作为map 的value 拓展 map 的类型选择
*/

func main() {
    
    var x interface{}
    x = "hello"
    x = 100
    x = false
    //fmt.Println(x)
/*
    var m = make(map[string]interface{}, 16)
    m["name"] = "kun"
    m["age"] = 18
    m["killed"] = []string{"ball", "bick"}
    fmt.Println(m)
    */

    // 类型断言 判断 该类型， 返回bool 类型， 正确为true
    ret, ok := x.(bool)
    if ok {
        fmt.Println(ret)
    }else {
        fmt.Println("not bool")
    }

    // 使用switch 进行类型断言
    switch t := x.(type) {
    case string:
        fmt.Println("string")
    case bool:
        fmt.Println("bool")
    case int:
        fmt.Println("int")
    default:
        fmt.Println("type not found")
    }



}
