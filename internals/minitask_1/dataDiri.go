package minitask1

import "fmt"

type Person struct {
	Name        string
	Photo       string
	Email       string
	Age         uint8
	PhoneNumber string
	IsMarried   bool
	Schools     []School
}

type School struct {
	Name  string
	Major string
}

// Getter Method
func (p *Person) GetBiodata() {
	statusPernikahan := "Belum Menikah"
	if p.IsMarried {
		statusPernikahan = "Sudah Menikah"
	}
	fmt.Printf("Nama: %s\nUmur: %d\nStatus Pernikahan: %s\n", p.Name, p.Age, statusPernikahan)
}

func (p *Person) IsEligibleForKTP() bool {
	return p.Age >= 17
}

// Setter Method
func (p *Person) UpdateName(newName string) {
	p.Name = newName
}

// constructor function
func NewPerson(name, photo, email, phoneNumber string, age uint8, isMarried bool, schools []School) *Person {
	return &Person{
		Name:        name,
		Photo:       photo,
		PhoneNumber: phoneNumber,
		Email:       email,
		Age:         age,
		IsMarried:   isMarried,
		Schools:     schools,
	}
}

func NewSchool(name, major string) School {
	return School{
		Name:  name,
		Major: major,
	}
}
