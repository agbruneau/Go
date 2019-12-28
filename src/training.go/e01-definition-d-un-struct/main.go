package main

type Address struct {
	street, city string
}

type Company struct {
	Name string
}

type Person struct {
	Name    string
	Age     int
	Addr    Address
}

func main() {
	var p Person
	p.Name = "Bob"
	p.Addr.city = "Lyon"
}