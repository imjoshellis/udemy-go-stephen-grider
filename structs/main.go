package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "hi@hi.com",
			zip:   23000,
		},
	}

	alex.updateName("lexi")
	alex.print()
}

func (p *person) updateName(f string) {
	p.firstName = f
}

func (p person) print() {
	fmt.Println(p)
}
