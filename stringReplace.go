package main

import (
	"fmt"
	"strings"
)

func replaceString(name string) string {
	nameReplace := strings.ReplaceAll(name, "h", "")

	return nameReplace
}

func main() {
	fmt.Println(replaceString("hHeloH"))
}
