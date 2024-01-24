package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name    string
	Surname string
}

type names []Person

func (a names) Len() int           { return len(a) }
func (a names) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a names) Less(i, j int) bool { return a[i].Surname < a[j].Surname }

func main() {
	people := []Person{
		{"Bob", "Smith"},
		{"John", "Doe"},
		{"Jane", "Doe"},
	}

	fmt.Println(people)
	sort.Sort(names(people))
	fmt.Println(people)
}
