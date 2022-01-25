package main

import "fmt"

func main() {
    count := 0
    s := "中无色粉哦osdfnaof"
    for _, c := range s {
        if c > 128 {
            count = count + 1
        }
    }
    fmt.Printf("%d\n", count)
}
