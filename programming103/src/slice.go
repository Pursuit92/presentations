package main

func main() {
	// slice OMIT
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	// end slice OMIT

	// slice2 OMIT
	// slice2 == [3,4]
	slice2 := slice[2:4]
	// end slice2 OMIT

	// slice3 OMIT
	// slice3 == [1,2,3,4]
	slice3 := slice[:4]
	// end slice3 OMIT
}
