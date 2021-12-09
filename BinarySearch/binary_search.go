package main

import "fmt"

func BinarySearch(listOfItems []int, valueToBeSearched int) bool {

	low := 0

	high := len(listOfItems) - 1

	for low <= high {
		avg := (low + high) / 2

		if listOfItems[avg] < valueToBeSearched {
			low = avg + 1
		} else {
			high = avg - 1
		}
	}

	if low == len(listOfItems) || listOfItems[low] != valueToBeSearched {
		return false
	}

	return true

}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(BinarySearch(items, 9))
}
