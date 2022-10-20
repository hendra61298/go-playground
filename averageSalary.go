package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {
	sort.Slice(salary, func(i, j int) bool {
		return salary[i] < salary[j]
	})
	var totalSalary int
	for i := 1; i < len(salary)-1; i++ {
		totalSalary += salary[i]
	}
	var result float64
	result = float64(totalSalary) / float64(len(salary)-2)

	return result
}

func main() {
	fmt.Println(average([]int{3000, 2000, 4000}))
}
