package main

import (
	"fmt"
	"sort"
)

func main() {
	testMap := map[string]int{"яблоко": 777, "банан": 2, "вишня": 7, "арбуз": 11}
	sorting(testMap)
}

func sorting(sortMap map[string]int) []string {
	sortedSlice := []string{}

	keys := make([]string, 0, len(sortMap))
	for k := range sortMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return sortMap[keys[i]] > sortMap[keys[j]]
	})

	for _, key := range keys {
		fmt.Println(key)
		sortedSlice = append(sortedSlice, key)
	}
	return sortedSlice
}
