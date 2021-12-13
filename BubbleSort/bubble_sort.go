package main

import "fmt"

func main() {
	unsortedArray := []int{-34, 45, 23, 54, -345, 5, 12, -86, 02}
	fmt.Println(BubbleSort(unsortedArray))
}

func BubbleSort(unsortedArray []int) []int {
	for i := 0; i < len(unsortedArray)-1; i++ {
		for j := 0; j < len(unsortedArray)-i-1; j++ {
			if unsortedArray[j] > unsortedArray[j+1] {
				unsortedArray[j], unsortedArray[j+1] = unsortedArray[j+1], unsortedArray[j]
			}
		}
	}
	return unsortedArray
}
