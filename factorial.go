package main

import "fmt"

func kthFactor(n int, k int) int {
	var list []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			list = append(list, i)
		}
	}
	if len(list) < k {
		return -1
	}
	return list[k-1]
}
func main() {
	fmt.Println(kthFactor(12, 3))
}
