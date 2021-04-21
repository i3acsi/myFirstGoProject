package main

import (
	"fmt"
	"sort"
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

type personSlice []person

func (ps personSlice) Len() int           { return len(ps) }
func (ps personSlice) Less(i, j int) bool { return ps[i].age < ps[j].age }
func (ps personSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func main() {
	ps := personSlice{
		person{name: "Vasiliy", age: 34},
		person{name: "Petr", age: 31},
		person{name: "Semen", age: 28},
	}
	fmt.Println(ps)
	sort.Sort(ps)
	fmt.Println(ps)
}
