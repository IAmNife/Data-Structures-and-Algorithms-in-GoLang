package main

import "fmt"

func interpolationSearch(items []int, searchKey int) int {

	min, max := items[0], items[len(items)-1] // Values in the array

	low, high := 0, len(items)-1

	for {
		if searchKey < min {
			return low //If the number to be searched is less than 0 => return 0
		}

		if searchKey > max {
			return high + 1 //If the number to be searched is greater than the highest number is array => length of listitems
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high //if both high and low is 0 => 0
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(searchKey-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if items[guess] == searchKey { //items[4] == 63
			// scan backwards for start of value range
			for guess > 0 && items[guess-1] == searchKey {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if items[guess] > searchKey {
			high = guess - 1
			max = items[high]
		} else {
			low = guess + 1
			min = items[low]
		}
	}
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(items, 63))
}
