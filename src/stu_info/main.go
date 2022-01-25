package main

import (
    "fmt"
)


func showMenu() {
    fmt.Println("welecome")
    fmt.Println("1. add")
    fmt.Println("2. dele")
    fmt.Println("3. show")
    fmt.Println("4. exit")
}

func getinput() *student {
    var (
        id int
        name string
        class string
    )
    fmt.Scanf("%d\n", &id)
    fmt.Scanf("%s\n", &name)
    fmt.Scanf("%s\n", &class)
    stu := newStudent(id, name, class)
    return stu
}



func main () {
    // 打印菜单
    sm := NewStudentMar()
    // 等待用户输入命令
    for {
        showMenu();
        var input int
        fmt.Scanf("%d\n", &input)
        switch input {
            case 1:
                //添加
                info := getinput()
                sm.addStudent(info)
                break
            case 2:
                // 编辑
                info := getinput()
                sm.modifyStudent(info)
            case 3:
                sm.showStudent()
            case 4:
                return
        }
    }
    return

}
