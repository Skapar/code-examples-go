package main

type Person struct {
	Name	string
	Surname	string
	Age		int
}

func (p Person) GetFullName() string {
	return p.Name + " " + p.Surname
}

func main() {
	person := Person{Name: "John", Surname: "Doe", Age: 25}
	println(person.GetFullName())
}