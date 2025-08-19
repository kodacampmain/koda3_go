package numbers

// variadic function
func Sum(numbers ...int) int {
	var total int = 0
	// normal for loop
	// for i := 0; i < len(numbers); i++ {
	// 	total += numbers[i]
	// }
	// for with range of collection
	// for i := range numbers {
	// 	total += numbers[i]
	// }
	// while loop (using for)
	i := 0
	for i < len(numbers) {
		total += numbers[i]
		i++
	}
	return total
}
