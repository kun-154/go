package main

import "fmt"

type student struct {
    id int
    name string
    class string
}

type studentMar struct {
    allStudents []*student
}

func NewStudentMar() *studentMar {
    return &studentMar {
        allStudents: make([]*student, 0, 100),
    }
}


func newStudent(id int, name, class string) *student {
    return &student{
        id : id,
        name : name,
        class : class,
    }
}



func (s *studentMar) addStudent(newStu *student) {
    fmt.Println("add")
    s.allStudents = append(s.allStudents, newStu)
}


func (s *studentMar) modifyStudent(newStu *student) {
    fmt.Println("modify")
    for i, v := range s.allStudents {
        if newStu.id == v.id {
            s.allStudents[i] = newStu
            return 
        }
    }
    fmt.Printf("not FOUNF\n")

}


func (s *studentMar) showStudent() {
    fmt.Println("show")
    for _, v := range s.allStudents {
        fmt.Printf("id: %d name: %s class: %s\n", v.id, v.name, v.class)
    }
}
