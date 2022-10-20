package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subtractProductAndSumV1(n int) int {
	var sum int
	product := 1

	for n > 0 {
		remainder := n % 10
		sum += remainder
		product *= remainder
		n = n / 10
	}

	return product - sum
}
func subtractProductAndSum(n int) int {
	product := 1
	var sumDigit int
	slice := strings.Split(strconv.FormatInt(int64(n), 10), "")
	for _, v := range slice {
		integerConvert, _ := strconv.Atoi(v)
		product *= integerConvert
		sumDigit += integerConvert
	}
	return product - sumDigit
}

func main() {
	fmt.Println(subtractProductAndSum(2145))
}
