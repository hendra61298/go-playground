package main

import "fmt"

type school struct {
	Name, Address string
	Grades        []int
}

func (s *school) AddGrade(grades []int) {
	for _, v := range grades {
		s.Grades = append(s.Grades, v)
	}
}

func Analysis(s *school) (float32, int, int) {
	s.AddGrade(s.Grades)
	var average float32
	var min, max = 100, 0
	var sum int

	for _, res := range s.Grades {
		if max < res {
			max = res
		}
		if min > res {
			min = res
		}
		sum += res
	}
	if len(s.Grades) != 0 {
		average = float32(sum / len(s.Grades))

	} else {
		average = 0
		min = 0
	}
	return average, min, max
}

func main() {
	fmt.Println(Analysis(&school{
		Name:    "InternationalSchool",
		Address: "Anywhere",
		Grades:  []int{0, 100, 0, 100},
	}))
}
