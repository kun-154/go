package main

import "fmt"

func main() {
    s := "s水电费"

    fmt.Println(len(s))
    // 组成每个字符串的元素叫做字符 字符包括 两种类型 byte  和 runne类型
    for _, c := range s {
        fmt.Printf("%c\n", c)
    }

    // 字符串修改 不能直接修改， 必须将字符串强制转换成 rune切片 (其中涉及到了内存的重新分配)
    s2 := "dsfdofivm"
    s3 := []rune(s2)
    //s3[0] = "红" 错误
    s3[0] = '红'
    fmt.Printf("%d\n", s3[0])

    //fmt.Printf("%c\n", s3[0])
    fmt.Println(string(s3)) // 将rune切片强制转换成字符串 
}

