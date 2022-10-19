package main

import (
	"fmt"
	"regexp"
	"sort"
)

func filter(names string) string {
	var regex, err = regexp.Compile(`[a-zA-Z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	name := regex.FindAllString(names, -1)
	sort.Strings(name)
	sort.Slice(name, func(i, j int) bool {
		return len(name[i]) < len(name[j])
	})
	return name[0]

}
func main() {
	fmt.Println(filter("Joko Tio Tia,Danar      "))
}
