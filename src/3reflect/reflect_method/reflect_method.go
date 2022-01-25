package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Score int
}

func (s student) Study() string {
	msg := "study"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "sleep"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("methodname:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {

}
