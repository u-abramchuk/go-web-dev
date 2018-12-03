package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}

type secretAgent struct {
	person
	id string
}

func (sa secretAgent) speak() {
	fmt.Printf("Agent %v (%v %v)\n", sa.id, sa.first, sa.last)
}

func info(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "John",
		last:  "Smith",
	}
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		id: "007",
	}

	info(p)
	info(sa)
}
