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
	fmt.Printf("original  slice -> %#v // No change.\n\n", sliceInt)

	fmt.Println("Case: []float64")
	sliceFloat64 := []float64{0.1, 0.1, 0.2, 0.2, 0.3, 0.3}
	dup.Do(&sliceFloat64)
	if sl, err := dup.Float64(); err == nil {
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf("original  slice -> %#v // No change.\n\n", sliceFloat64)

	fmt.Println("Case: []bool")
	sliceBool := []bool{true, true, true, true, true, true, true, true, true, true, true, true, false, false, false}
	dedupe.Do(&sliceBool)
	fmt.Printf("original  slice -> %#v\n\n", sliceBool)

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
	fmt.Printf("original  slice -> %#v\n", sliceStruct)

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
	if sl, err := dup.Any(); err == nil {
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf("original  slice -> %#v // No change.\n\n", sliceStruct)

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
	fmt.Printf("original  slice -> %#v\n", slicePtr)

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
	if sl, err := dup.Any(); err == nil {
		fmt.Printf("extracted slice -> %#v\n", sl)
	}
	fmt.Printf("original  slice -> %#v // No change.\n\n", slicePtr)

	fmt.Println("Case: []UniqueType")
	type AgeInt int
	sliceUnique := []AgeInt{1, 2, 3, 3}
	dedupe.Do(&sliceUnique)
	fmt.Printf("original  slice -> %#v\n\n", sliceUnique)

	fmt.Println("Case: []UniqueType2")
	type otherType duplication
	sliceOtherType := []otherType{
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
	dedupe.Do(&sliceOtherType)
	fmt.Printf("original  slice -> %#v\n\n", sliceOtherType)

	fmt.Println("Case: []func")
	type Func func(string) bool
	sliceFunc := []Func{
		func(s string) bool {
			return true
		},
		func(s string) bool {
			return false
		},
		func(s string) bool {
			return true
		},
	}
	if err := dedupe.Do(&sliceFunc); err != nil {
		fmt.Println(err)
	}
}
