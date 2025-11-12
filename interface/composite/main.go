package main

import (
	"fmt"
	"math/rand"
)

type ICounter interface {
	count() int
}

type Employee struct {
}

// Employee is a case of Icounter
func (e *Employee) count() int {
	return rand.Intn(100)
}

type Department struct {
	Name     string
	Counters []ICounter
}

func (e *Department) count() int {
	ans := 0
	for _, counter := range e.Counters {
		ans += counter.count()
	}
	return ans
}

func NewEmplyee() *Employee {
	return &Employee{}
}

func NewDepartment() *Department {
	return &Department{
		Name: "DDM-R",
	}
}

func main() {
	dept := NewDepartment()
	emp1 := NewEmplyee()
	emp2 := NewEmplyee()

	fmt.Println(emp1.count())
	fmt.Println(emp2.count())
	dept.Counters = append(dept.Counters, emp1)
	dept.Counters = append(dept.Counters, emp2)

	total := dept.count()
	println("Total employees:", total)
}
