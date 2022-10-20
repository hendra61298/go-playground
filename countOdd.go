package main

import "fmt"

func countOdds(low int, high int) int {
	var result int
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(countOdds(10, 15))
}
