package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(mult(a))
	fmt.Println(a)

}
func mult(a []int) []int {
	dSlice := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		dSlice[i] = a[i] * 2
	}
	return dSlice
}
