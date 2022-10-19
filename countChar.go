package main

import (
	"fmt"
	"sort"
	"strings"
)

func countChar(sentences string) (int, []byte, []string) {
	sentencesWithoutSpace := strings.ReplaceAll(sentences, " ", "")
	amountSentences := len(sentencesWithoutSpace)

	char := strings.Split(sentencesWithoutSpace, "")
	sort.Strings(char)
	byteArray := []byte(sentencesWithoutSpace)
	sort.Slice(byteArray, func(i, j int) bool {
		return byteArray[i] < byteArray[j]
	})
	return amountSentences, byteArray, char
}

func main() {
	fmt.Println(countChar("Halo Kaka"))
}
