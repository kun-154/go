package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "li",
		Score: 12,
	}

	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	//遍历结构体遍历变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引去取字段
		filedObj := t.Field(i)
		fmt.Printf("name:%v, type:%v, tag:%v\n", filedObj.Name, filedObj.Type, filedObj.Tag)
		// 取指定的 tag 的值
		fmt.Println(filedObj.Tag.Get("json"))
	}

	// 根据名字去结构体中 取字段信息
	filed, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n")
	}

}
