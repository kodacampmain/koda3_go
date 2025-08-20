package main

import (
	"fmt"
	"log"

	"github.com/kodacampmain/koda3_go/internals/intermediate"
	minitask1 "github.com/kodacampmain/koda3_go/internals/minitask_1"
	"github.com/kodacampmain/koda3_go/internals/service"
)

func main() {
	// utils.Greet()
	// // inference
	// // str := "Hello World"
	// // bt := []byte("golang")
	// // fmt.Printf("%s\n%v\n", str, bt)
	// // rn := []rune("ひ")
	// // rnj := []rune(" ꦈꦴ")
	// // fmt.Println(rn)
	// // fmt.Println(rnj)
	// // manifest
	// // var numb uint = 1
	// // fmt.Println(numb)
	// // angka menjadi hari
	// day, err := utils.NumToDay(0)
	// // explicit error handling
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(day)
	// }
	// total := numbers.Sum(1, 2, 3, 4, 5)
	// fmt.Println(total)
	// // for dengan range of int
	// // for i := range 3 {
	// // 	fmt.Println(i)
	// // }

	// collection.ArrayAndSlice()

	// collection.MapAndStruct()

	// batman := collection.Movies{
	// 	Title:  "Batman",
	// 	Genres: []string{"Action", "Romance"},
	// 	Id:     68,
	// }
	// fmt.Println(batman)
	smk := minitask1.NewSchool("SMK Sumber Rejeki", "Electrical")
	ithb := minitask1.NewSchool("Institut Teknologi Harapan Bangsa", "Electrical Engineering")
	user := minitask1.NewPerson("Andi", "andi.webp", "andi@mail.com", "+6285768904329", 25, false, []minitask1.School{smk, ithb})
	// user := minitask1.Person{
	// 	Name:        "Andi",
	// 	Photo:       "andi.webp",
	// 	Email:       "andi@mail.com",
	// 	Age:         25,
	// 	PhoneNumber: "+6285768904329",
	// 	IsMarried:   false,
	// 	Schools: []minitask1.School{
	// 		{Name: "SMK Sumber Rezeki", Major: "Electrical"},
	// 		{Name: "Institut Teknologi Harapan Bangsa", Major: "Electrical Engineering"},
	// 	},
	// }
	// fmt.Println(user)
	user.GetBiodata()
	fmt.Printf("Sudah punya KTP: %t\n", user.IsEligibleForKTP())
	user.UpdateName("Candra")
	user.GetBiodata()

	// email := "a@mail.com"
	// pass := "pass"
	// success, err := minitask1.Login(email, pass)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(success)
	// }
	// fmt.Println(minitask1.BigNumber())

	// fromUnit := "C"
	// targetUnit := "F"
	// var startTemp float32 = 25
	// resultTemp, err := minitask1.TemperatureConversion(startTemp, fromUnit, targetUnit)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("%.2f%s -> %.2f%s\n", startTemp, fromUnit, resultTemp, targetUnit)
	// }
	// fmt.Println(minitask1.InsertIntoSlice([]int{50, 75, 66, 20, 32, 90}, 66, 88))
	// fmt.Println(minitask1.InsertIntoSlice([]string{"Hello", "World"}, "Hello", "Halo"))
	// defer fmt.Println("Program Selesai 1")
	// defer fmt.Println("Program Selesai 2")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("panic recovered")
	// 		fmt.Println(r)
	// 	}
	// }()
	// intermediate.DeferredFunction()
	// intermediate.MustDo(false)
	// err := errors.New("terjadi error")
	// if err != nil {
	// 	intermediate.Exit(1, err)
	// }
	// fmt.Println("Proses")

	var day int = 20
	var ptrDay *int = &day

	*ptrDay = 10
	// fmt.Println(day)
	day = 30
	// fmt.Println(*ptrDay)
	// fmt.Println(&day, ptrDay)
	name := "Andi"
	// fmt.Println(name)
	intermediate.ChangeName(&name)
	// fmt.Println(name)
	// setter method

	samsung := intermediate.TV{}
	gree := intermediate.AC{}
	miyako := intermediate.ElectricFan{}

	intermediate.UniversalRemote("ON", &samsung)
	intermediate.UniversalRemote("ON", &gree)
	intermediate.UniversalRemote("ON", &miyako)
	intermediate.UniversalRemote("OFF", &gree)
	intermediate.UniversalRemote("OFF", &samsung)

	result := service.MakePizza()
	log.Printf("Pizza berhasil dibuat dengan keju %s dan saus tomat %s", result.Cheese.Type, result.Tomato.Brand)

	mozarella := service.NewCheese("Mozarella")
	cheddar := service.NewCheese("Cheddar")
	heinz := service.NewTomato("Heinz")
	ayam := service.NewTomato("Ayam")
	betterResult1 := service.BetterMakePizza(mozarella, heinz)
	betterResult2 := service.BetterMakePizza(mozarella, ayam)
	betterResult3 := service.BetterMakePizza(cheddar, heinz)
	betterResult4 := service.BetterMakePizza(cheddar, ayam)
	service.LogPizza(betterResult1)
	service.LogPizza(betterResult2)
	service.LogPizza(betterResult3)
	service.LogPizza(betterResult4)
}
