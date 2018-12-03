package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func (p person) pSpeak() {
	fmt.Println(p.first, p.last)
}

type secretAgent struct {
	person
	id string
}

func (sa secretAgent) saSpeak() {
	fmt.Printf("Agent %v (%v %v)\n", sa.id, sa.first, sa.last)
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

	fmt.Println("Person's first name is", p.first)

	p.pSpeak()

	fmt.Println("Secret agent's id is", sa.id)

	sa.saSpeak()
	sa.pSpeak()
}
