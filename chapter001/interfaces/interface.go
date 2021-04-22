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

type NickNamed interface {
	NickName() string
}

func joinChat(nn NickNamed) {
	p, ok := nn.(person)
	if ok {
		_ = p.name
	}
	fmt.Println(nn.NickName() + " has joined.")
}

func interfaces() {
	p := person{name: "Vasiliy", age: 34}
	joinChat(p)

}

func main() {
	interfaces()
}
