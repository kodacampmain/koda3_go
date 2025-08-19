package collection

import "fmt"

type Movies struct {
	Title  string
	Genres []string
	Id     int
}

func MapAndStruct() {
	genres := make(map[int]string)

	genres[1] = "Action"
	genres[10] = "Adventure"
	genres[100] = "Fantasy"

	genress := map[int]string{
		1:   "Action",
		10:  "Adventure",
		100: "Fantasy",
	}
	genres[5] = "Comedy"
	genress[5] = "Comedy"
	fmt.Println(genres)
	fmt.Println(genress)

	spiderman := Movies{
		Title:  "Spiderman",
		Genres: []string{"Action", "Fantasy"},
		Id:     125,
	}
	fmt.Println(spiderman)
}
