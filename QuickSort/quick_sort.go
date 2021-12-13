package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piece := createPiece(15) //Generate 15 random numbers
	fmt.Println("\n Unsorted Array \n", piece)
	quickSort(piece)
	fmt.Println("\n Sorted \n", piece)
}

func createPiece(size int) []int {
	piece := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		piece[i] = rand.Intn(999) - rand.Intn(999) //Making sure we are generating negative numbers too
	}
	return piece
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	center := rand.Int() % len(a)
	a[center], a[right] = a[right], a[center]
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort((a[:left]))
	quickSort(a[left+1:])

	return a

}
