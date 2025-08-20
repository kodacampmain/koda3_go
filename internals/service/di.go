package service

import "log"

type Cheese struct {
	Type string
}
type Tomato struct {
	Brand string
}

func NewCheese(cheeseType string) Cheese {
	return Cheese{Type: cheeseType}
}
func NewTomato(brand string) Tomato {
	return Tomato{Brand: brand}
}

type Pizza struct {
	Cheese Cheese
	Tomato Tomato
}

func MakePizza() Pizza {
	mozarella := NewCheese("Mozarella")
	// cheddar := NewCheese("Cheddar")

	heinz := NewTomato("Heinz")
	// ayam := NewTomato("Ayam")

	return Pizza{Cheese: mozarella, Tomato: heinz}
	// return Pizza{Cheese: cheddar, Tomato: ayam}
}

func BetterMakePizza(cheese Cheese, tomato Tomato) Pizza {
	return Pizza{Cheese: cheese, Tomato: tomato}
}

func LogPizza(pizza Pizza) {
	log.Printf("Pizza berhasil dibuat dengan keju %s dan saus tomat %s", pizza.Cheese.Type, pizza.Tomato.Brand)
}
