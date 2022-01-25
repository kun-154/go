package main

import "fmt"

func main() {
    forloop1:
    for i := 1; i < 10; i++ {
        for j := 1; j < 10; j++ {
            if j == 2 {
                break forloop1
            }
            fmt.Print(j)
        }
    }
}
