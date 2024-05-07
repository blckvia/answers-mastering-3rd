package main

import (
	"fmt"
	"reflect"
)

func PrintReflection(s interface{}) {
	fmt.Println("** Reflection")
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			fmt.Print(val.Index(i).Interface(), " ")
		}
		fmt.Println()
	} else if val.Kind() == reflect.String {
		fmt.Println(val.String())
	}
}

func PrintSlice[T any](s []T) {
	fmt.Println("** Generics")
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func Print[T any](s T) {
	fmt.Println("** Generics")
	fmt.Println(s)
}

func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.2, -2.33, 4.55})
	Print("hello")
	Print(1)

	PrintReflection([]int{1, 2, 3})
	PrintReflection([]string{"a", "b", "c"})
	PrintReflection([]float64{1.2, -2.33, 4.55})
	PrintReflection("hello")
}
