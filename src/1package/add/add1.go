package add
import "fmt"

var Name = "init"

func init() {
    fmt.Println(Name)
    fmt.Println("add import")
}

// 包在导入时就已完成全局函数的初始化， 也在init函数执行之前 

func Add(x, y int) int{
    return x + y
}
