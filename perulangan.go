package main

import (
	"fmt"
	"strings"
)

func reverseV1(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	rns := strings.Split("Hel lo", "")
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	fmt.Println(strings.Join(rns, ""))

}
