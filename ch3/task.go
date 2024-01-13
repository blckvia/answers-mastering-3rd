package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"sort"
)

// arrayToMap converts an array to a map
func arrayToMap(a interface{}) (map[int]bool, error) {
	valA := reflect.ValueOf(a)

	if valA.Kind() != reflect.Array {
		return nil, fmt.Errorf("a is not an array")
	}

	newMap := make(map[int]bool)

	for i := 0; i < valA.Len(); i++ {
		newMap[valA.Index(i).Interface().(int)] = true
	}

	return newMap, nil

}

// mapToSlices converts a map to 2 slices
func mapToSlices(a interface{}) (interface{}, interface{}, error) {
	valA := reflect.ValueOf(a)

	if valA.Kind() != reflect.Map {
		return nil, nil, fmt.Errorf("a is not a map")
	}

	keys := valA.MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Int() < keys[j].Int()
	})

	firstSlice := make([]interface{}, len(valA.MapKeys()))
	secondSlice := make([]interface{}, len(valA.MapKeys()))

	for i, v := range keys {
		firstSlice[i] = v.Interface()
		secondSlice[i] = valA.MapIndex(v).Interface()
	}

	return firstSlice, secondSlice, nil
}

// Task: Change csvData.go to separate record fields based on
// # symbol.
func saveCSVFileTask(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	csvwriter.Comma = '#' // jast change thi
	//for _, row := range myData {
	//	temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
	//	_ = csvwriter.Write(temp)
	//}
	csvwriter.Flush()
	return nil
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arrayToMap(arr))

	newMap := map[int]bool{
		1: true,
		2: false,
		3: false,
	}

	fmt.Println(mapToSlices(newMap))
}
