package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}

func (p person) NickName() string {
	return p.name + " " + strconv.Itoa(p.age)
}

type employee struct {
	person
	jobTitle string
}

func methods() {
	emp := employee{person{name: "Vasiliy", age: 34}, "developer"}
	fmt.Println(emp.person.name, emp.person.age, emp.jobTitle)
	fmt.Println(emp.name, emp.age, emp.jobTitle)
	fmt.Println(emp.NickName())
	p := emp.person
	fmt.Println(p.name, p.age)
	fmt.Println(p.NickName())

}

func main() {
	methods()
}
