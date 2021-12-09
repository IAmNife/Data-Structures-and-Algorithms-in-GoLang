package main

import "fmt"

func linearSearch(listOfItems []int, valueToBeSearched int) bool {
	for _, item := range listOfItems {
		if item == valueToBeSearched {
			return true
		}
	}
	return false
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(linearSearch(items, 2))
}
