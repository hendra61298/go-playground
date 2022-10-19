package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatPassword(pw string) string {
	stringVocalUpper := []string{"A", "E", "I", "O", "U"}
	stringVocalLower := []string{"a", "e", "i", "o", "u"}

	reverseString := reverse(pw)
	spliceString := strings.Split(reverseString, "")
	for index, res := range spliceString {
		for i := 0; i < len(stringVocalLower); i++ {
			if i == len(stringVocalLower)-1 {
				if res == stringVocalUpper[i] {
					spliceString[index] = strings.ToLower(stringVocalUpper[0])
				} else if res == stringVocalLower[i] {
					spliceString[index] = strings.ToUpper(stringVocalLower[0])
				}
			} else {
				if res == stringVocalUpper[i] {
					spliceString[index] = strings.ToLower(stringVocalUpper[i+1])
				} else if res == stringVocalLower[i] {
					spliceString[index] = strings.ToUpper(stringVocalLower[i+1])
				}
			}
			if strings.ToUpper(res) != stringVocalUpper[i] {
				if IsUpper(res) {
					spliceString[index] = strings.ToLower(spliceString[index])
				} else {
					spliceString[index] = strings.ToUpper(spliceString[index])
				}
			}
		}
	}

	return strings.Join(spliceString, "")
}

func main() {
	fmt.Println(formatPassword("HUlluoPw"))
}
