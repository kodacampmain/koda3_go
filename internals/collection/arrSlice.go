package collection

import "fmt"

func ArrayAndSlice() {
	hobbies := [3]string{"reading", "swimming", "sleeping"}
	var myHobbies []string = hobbies[1:]
	// myHobbies = append(myHobbies, "gaming")
	// myHobbies[1] = "running"
	// hobbies[1] = "running"
	fmt.Println(hobbies)
	fmt.Println(myHobbies)

	ages1 := []int{}
	var ages2 []int
	agesArray := [3]int{}
	fmt.Println(ages1)
	fmt.Println(ages2)
	fmt.Println(agesArray)
}
