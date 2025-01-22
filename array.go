package main

import "fmt"

func array() {
	var defaultNumbers [5]int

	fmt.Println("defaultNumbers:", defaultNumbers)

	assinedNumbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println("assinedNumbers:", assinedNumbers)

	arrayLenght := len(assinedNumbers)

	fmt.Printf("assinedNumbers's lenght %d \n", arrayLenght)

	firstElement := assinedNumbers[0] // arrays in Go are zero-based indexing

	fmt.Printf("first element of assignedNumbers is %d \n", firstElement)
}

func slice() {
	var defaultIntSlice []int
	fmt.Println("defaultIntSlice", defaultIntSlice)

	makeSlice := make([]int, 3) // kind of arrays with capacity
	fmt.Println("makeSlice", makeSlice)

	numbers := [5]int{1, 2, 3, 4, 5}
	sliceFromAnExistedArray := numbers[1:3]
	fmt.Println("sliceFromAnExistedArray", sliceFromAnExistedArray)
	sliceFromAnExistedArray = append(sliceFromAnExistedArray, 6)
	fmt.Println("sliceFromAnExistedArray", sliceFromAnExistedArray)

	hardCopied := sliceFromAnExistedArray
	sliceFromAnExistedArray = append(sliceFromAnExistedArray, 7)
	fmt.Println("sliceFromAnExistedArray", sliceFromAnExistedArray)
	fmt.Println("hardCopied", hardCopied)

}
