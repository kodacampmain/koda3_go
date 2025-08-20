package minitask1

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
