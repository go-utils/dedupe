package main

import (
	"fmt"

	"github.com/go-utils/dedupe"
)

func main() {
	dup := dedupe.NewDeduplication()
	fmt.Println("Case: []string")
	sliceStr := []string{"Go", "V", "Java", "Python", "Go", "Ruby", "Go", "V"}
	dedupe.Do(&sliceStr)
	fmt.Printf("original  slice -> %#v\n\n", sliceStr)

	fmt.Println("Case: []int")
	sliceInt := []int{1, 1, 2, 2, 3, 3}
	dup.Do(&sliceInt)
	if sl, err := dup.Int(); err == nil {
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf(  "original  slice -> %#v / No change.\n\n", sliceInt)

	fmt.Println("Case: []float64")
	sliceFloat64 := []float64{0.1, 0.1, 0.2, 0.2, 0.3, 0.3}
	dup.Do(&sliceFloat64)
	if sl, err := dup.Float64(); err == nil {
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf(  "original  slice -> %#v / No change.\n\n", sliceFloat64)

	type duplication struct {
		Name string
		Age  int
	}

	fmt.Println("Case: []struct")
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
	fmt.Printf("original  slice -> %#v\n\n", sliceStruct)

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
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf(  "original  slice -> %#v / No change.\n\n", sliceStruct)

	fmt.Println("Case: []Pointer")
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
	fmt.Printf("original  slice -> %#v\n\n", slicePtr)

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
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf(  "original  slice -> %#v / No change.\n\n", slicePtr)

	// TODO FIXME
	fmt.Println("Case: []UniqueType")
	type AgeInt int
	sliceUnique := []AgeInt{1,2,3,3}
	fmt.Println(dedupe.Do(&sliceUnique))
	fmt.Printf("original  slice -> %#v\n\n", sliceUnique)
}
