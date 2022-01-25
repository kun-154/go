package main

import "fmt"
import "encoding/json"

// 结构体 字段的可见性  一定 的命名规范为 字段 大写开头的表示可以公开访问， 小写的表示私有

// 如果一个 go语言 包中定义的标识符是首字符大写的， 那么就是 对外可见得




// json  序列化

type student struct {
    Id int
    Name string
}

type class struct {
    title string
    Students []student
}

type class1 struct {
    Title string `json:"title" db:"class_title" xml:"tt"` // 表示 在json， 数据库， xml 中的tag， 多个tag 用空格分割
    //表示 在json包中， 使用 class1 结构体时， 将名称 改为 title 
    // 这是 go 中 tag（标签）的规则  形式为  key:"value"
    Students []student
}

func newStudent(id int, name string) student {
    return student{
        Id: id,
        Name: name,
    }
}




func main() {
    c1 := class{
        title: "101",
        Students: make([]student, 0, 20),
    }

    for i := 0; i < 10; i++ {
        c1.Students = append(c1.Students, newStudent(i, fmt.Sprintf("stu%02d", i)))
        //tmpstu := newStudent(i, fmt.Sprintf("stu%02d", i))
        //c1.Students = append(c1.Students, tmpstu)
    }

    fmt.Printf("%#v\n", c1)


    // JSON 序列化， go语言中的 数据， -》json格式的 字符串
    data, err := json.Marshal(c1)
    if err != nil {
        fmt.Println("json marshal fiailed, err", err)
        return 
    }
    fmt.Printf("%T\n", data)
    fmt.Printf("%s\n", data)

    // 反序列化  
    var c2 class
    err = json.Unmarshal(data, &c2)
    if err != nil {
        fmt.Println("json uncoding fail, err:", err)
        return 
    }
    fmt.Printf("%#v\n", c2)
    



}
