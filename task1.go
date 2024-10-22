package main

import (
	"fmt"
	"sort"
)

func main() {
	// task1
	myArray := [10]int{1, 22, 41, 56, 78, 90, 15, 18, 75, 38}
	mySlice := make([]int, len(myArray))
	copy(mySlice, myArray[:])
	sort.Ints(mySlice)
	fmt.Println(mySlice)
	fmt.Println(myArray)
}
