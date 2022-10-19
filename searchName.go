package main

import (
	"fmt"
	"sort"
	"strings"
)

func findName(names string) string {
	name1 := strings.ReplaceAll(names, ",", " ")
	name2 := strings.ReplaceAll(name1, ";", " ")
	name := strings.Split(name2, " ")
	sort.Strings(name)
	sort.Slice(name, func(i, j int) bool {
		return len(name[i]) < len(name[j])
	})
	return name[0]

}

func main() {
	fmt.Println(findName("Joko Abi Tio Tia Danar"))
}
