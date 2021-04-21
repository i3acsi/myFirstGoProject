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
	return p.name + strconv.Itoa(p.age)
}

func methods() {
	p := person{
		name: "Vasiliy",
		age:  34,
	}
	fmt.Println(p.NickName())
	fmt.Println(person.NickName(p))

}

func main() {
	methods()
}
