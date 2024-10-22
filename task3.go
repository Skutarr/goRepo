package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{}
	addProd(myMap, "Egg", 15)
	deleteProd(myMap, "Egg")

}

func addProd(prodTable map[string]int, prodName string, prodAmout int) {
	prodTable[prodName] = prodAmout
	fmt.Println(prodTable)
}

func deleteProd(prodTable map[string]int, prodName string) {
	delete(prodTable, prodName)
	fmt.Println(prodTable)

}
