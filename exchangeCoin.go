package main

import "fmt"

var pecahan = []int{
	1000, 500, 200, 50, 20, 10, 5, 1,
}

func exchange(amount int) []int {
	var arrayPecahan []int
	var pembagi int
	for _, res := range pecahan {
		if amount >= res {
			pembagi = amount / res
			amount -= res * pembagi
			for i := 1; i <= pembagi; i++ {
				arrayPecahan = append(arrayPecahan, res)
			}

		}
	}
	return arrayPecahan
}

func main() {
	fmt.Println(exchange(5000))
}
