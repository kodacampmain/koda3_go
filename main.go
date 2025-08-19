package main

import (
	"fmt"

	"github.com/kodacampmain/koda3_go/internals/collection"
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
}
