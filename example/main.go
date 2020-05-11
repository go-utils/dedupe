package main

import (
	"fmt"

	"github.com/go-utils/dedupe"
)

func main() {
	dup := dedupe.NewDeduplication()
	fmt.Println("[]string")
	sliceStr := []string{"Go", "V", "Java", "Python", "Go", "Ruby", "Go", "V"}
	dedupe.Do(&sliceStr)
	fmt.Println("original  slice  ->", sliceStr)

	fmt.Println("[]int")
	sliceInt := []int{1, 1, 2, 2, 3, 3}
	dup.Do(&sliceInt)
	if sl, err := dup.Int(); err == nil {
		fmt.Println("extracted slices ->", sl)
	}
	fmt.Println("original  slice  ->", sliceInt)

	fmt.Println("[]float64")
	sliceFloat64 := []float64{0.1, 0.1, 0.2, 0.2, 0.3, 0.3}
	dup.Do(&sliceFloat64)
	if sl, err := dup.Float64(); err == nil {
		fmt.Println("extracted slices ->", sl)
	}
	fmt.Println("original  slice  ->", sliceFloat64)

	type duplication struct {
		Name string
		Age  int
	}

	fmt.Println("[]struct")
	sliceStruct := []duplication{
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "B",
			Age:  2,
		},
	}
	dedupe.Do(&sliceStruct)
	fmt.Println("original  slice  ->", sliceStruct)

	sliceStruct = []duplication{
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "B",
			Age:  2,
		},
	}
	dup.Do(&sliceStruct)
	if sl, err := dup.Struct(); err == nil {
		fmt.Println("extracted slices ->", sl)
	}
	fmt.Println("original  slice  ->", sliceStruct)

	fmt.Println("[]Pointer")
	slicePtr := []*duplication{
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "B",
			Age:  2,
		},
	}
	dedupe.Do(&slicePtr)
	fmt.Println("original  slice  ->", slicePtr)

	slicePtr = []*duplication{
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "B",
			Age:  2,
		},
	}
	dup.Do(&slicePtr)
	if sl, err := dup.Struct(); err == nil {
		fmt.Println("extracted slices ->", sl)
	}
	fmt.Println("original  slice  ->", slicePtr)
}
