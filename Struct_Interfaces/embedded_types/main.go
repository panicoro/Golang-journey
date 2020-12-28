package main

import "fmt"

// Person type represents a person with his or her name.
type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Android type represents an android with the person that copies
// and the model
type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "Number 18"
	a.Person.talk()
	a.talk()
}
