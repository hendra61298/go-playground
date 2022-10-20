package main

import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	var result int
	s := strconv.FormatUint(uint64(num), 2)
	for _, res := range s {
		if res == '1' {
			result++
		}
	}
	return result
}
func main() {
	fmt.Println(hammingWeight(uint32(00111001)))
}
