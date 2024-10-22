package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"яблоко", "банан", "яблоко", "вишня"}
	fmt.Println(stat(mySlice))

}

func stat(mySlice []string) map[string]int {
	myMap := map[string]int{}
	for i := 0; i < len(mySlice); i++ {
		el, status := myMap[mySlice[i]]
		if status == true {
			myMap[mySlice[i]] = el + 1
		} else {
			myMap[mySlice[i]] = 1
		}
	}

	return myMap

}
