package models

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (p *Person) Print() {
	fmt.Printf("Nama: %s\nAlamat: %s\nNomor Telepon: %s\n", p.Name, p.Address, p.Phone)
}

func (p *Person) Greet() {
	fmt.Printf("Hello %s", p.Name)
}

func (p *Person) SetName(newName string) {
	p.Name = newName
}
