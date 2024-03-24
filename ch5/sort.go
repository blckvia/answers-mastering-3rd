package main

import "fmt"

func firstSort(a, b, c int) (d, e, f int) {
	if a < b {
		if b < c {
			d, e, f = a, b, c
		} else if a < c {
			d, e, f = a, c, b
		} else {
			d, e, f = c, a, b
		}
	} else {
		if a < c {
			d, e, f = b, a, c
		} else if b < c {
			d, e, f = b, c, a
		} else {
			d, e, f = c, b, a
		}
	}

	return
}

func secondSort(a, b, c int) []int {
	values := []int{a, b, c}
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return values
}

func main() {
	sortedValues := secondSort(5, 3, 4)
	first, second, third := firstSort(5, 3, 4)
	fmt.Println(sortedValues)
	fmt.Println(first, second, third)
}
