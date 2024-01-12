package main

import (
	"fmt"
	"reflect"
)

// mergingSlices merges two slices
func mergingSlices(first []int, second []int) ([]int, error) {
	if len(first) == 0 || len(second) == 0 {
		return nil, fmt.Errorf("one of the slices is empty")
	}

	newSlice := make([]int, 0, len(first)+len(second))

	newSlice = append(newSlice, first...)
	newSlice = append(newSlice, second...)

	return newSlice, nil
}

// mergingArrays merges two arrays
func mergingArrays(a, b interface{}) ([]int, error) {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() != reflect.Array || bVal.Kind() != reflect.Array {
		return nil, fmt.Errorf("one of the arguments is not an array")
	}

	aSlice := make([]int, aVal.Len())
	bSlice := make([]int, bVal.Len())

	for i := 0; i < aVal.Len(); i++ {
		aSlice[i] = aVal.Index(i).Interface().(int)
	}

	for i := 0; i < bVal.Len(); i++ {
		bSlice[i] = bVal.Index(i).Interface().(int)
	}

	merged := append(aSlice, bSlice...)

	return merged, nil
}

// sliceFromArrays creates an array from two slices
func sliceFromArrays(a, b []int) ([6]int, error) {
	if len(a) == 0 || len(b) == 0 {
		return [6]int{}, fmt.Errorf("one of the slices is empty")
	}

	var result [6]int
	i := 0

	for _, v := range a {
		result[i] = v
		i++
	}

	for _, v := range b {
		result[i] = v
		i++
	}

	return result, nil
}

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	newSlice, err := mergingSlices(first, second)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newSlice)

	firstArray := [3]int{1, 2, 3}
	secondArray := [4]int{4, 5, 6, 7}

	newArray, err := mergingArrays(firstArray, secondArray)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newArray)

	array, err := sliceFromArrays(first, second)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(array)
}
