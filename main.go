package main

import (
	"fmt"

	"github.com/kodacampmain/koda3_go/internals/collection"
	minitask1 "github.com/kodacampmain/koda3_go/internals/minitask_1"
	"github.com/kodacampmain/koda3_go/internals/numbers"
	"github.com/kodacampmain/koda3_go/internals/utils"
)

func main() {
	utils.Greet()
	// inference
	// str := "Hello World"
	// bt := []byte("golang")
	// fmt.Printf("%s\n%v\n", str, bt)
	// rn := []rune("ひ")
	// rnj := []rune(" ꦈꦴ")
	// fmt.Println(rn)
	// fmt.Println(rnj)
	// manifest
	// var numb uint = 1
	// fmt.Println(numb)
	// angka menjadi hari
	day, err := utils.NumToDay(0)
	// explicit error handling
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(day)
	}
	total := numbers.Sum(1, 2, 3, 4, 5)
	fmt.Println(total)
	// for dengan range of int
	// for i := range 3 {
	// 	fmt.Println(i)
	// }

	collection.ArrayAndSlice()

	collection.MapAndStruct()

	batman := collection.Movies{
		Title:  "Batman",
		Genres: []string{"Action", "Romance"},
		Id:     68,
	}
	fmt.Println(batman)

	user := minitask1.Person{
		Name:        "Andi",
		Photo:       "andi.webp",
		Email:       "andi@mail.com",
		Age:         25,
		PhoneNumber: "+6285768904329",
		IsMarried:   false,
		Schools: []minitask1.School{
			{Name: "SMK Sumber Rezeki", Major: "Electrical"},
			{Name: "Institut Teknologi Harapan Bangsa", Major: "Electrical Engineering"},
		},
	}
	fmt.Println(user)

	email := "a@mail.com"
	pass := "pass"
	success, err := minitask1.Login(email, pass)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(success)
	}
	fmt.Println(minitask1.BigNumber())

	fromUnit := "C"
	targetUnit := "F"
	var startTemp float32 = 25
	resultTemp, err := minitask1.TemperatureConversion(startTemp, fromUnit, targetUnit)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%.2f%s -> %.2f%s\n", startTemp, fromUnit, resultTemp, targetUnit)
	}
	fmt.Println(minitask1.InsertIntoSlice([]int{50, 75, 66, 20, 32, 90}, 66, 88))
	fmt.Println(minitask1.InsertIntoSlice([]string{"Hello", "World"}, "Hello", "Halo"))
}
