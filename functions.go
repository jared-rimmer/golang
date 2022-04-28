package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     uint
	grades  []int
	subject string
}

func (p *Person) setAge(age uint) {
	p.age = age
}

func (p *Person) setSubject(subject string) {
	p.subject = subject
}

func (p *Person) calculateAverageGrade() float32 {
	var TotalGrades = 0
	for _, v := range p.grades {
		TotalGrades += v
	}
	return float32(TotalGrades) / float32(len(p.grades))
}

func main() {
	a := Person{"Jared", 96, []int{10, 20, 30}, "Golang"}
	fmt.Println(a)
	a.grades = append(a.grades, 90)
	fmt.Println(a)
	a.setAge(30)
	fmt.Println(a)
	a.setSubject("Python")
	fmt.Println(a)
	fmt.Println(a.calculateAverageGrade())
}
