# What is this?
[]int{1, 2, 3, 3}, []string{"1", "2", "3", "3"}, []float64{0.1, 0.2, 0.3, 0.3} etc.   
You have to implement each type by yourself when removing duplicates from existing slices.  
At that time, I made something that was put together and can be used even a little.

# Installation
```commandline
go get -u github.com/go-utils/dedupe
```

# Usage
```go
import (
    "fmt"

    "github.com/go-utils/dedupe"
)

func main() {
    fmt.Println("[Case1]")
    // Basic usage
    sliceString := []string{"Go", "V", "C++", "Java", "Python", "Go", "Ruby", "C++", "Go", "V"}
    dedupe.Do(&sliceString)                               // Be sure to pass it by address
    fmt.Printf("original  slice -> %#v\n\n", sliceString) // change original slice

    fmt.Println("[Case2]")
    // Extract duplicates without changing the original slice
    // But: If it is a structure, you must cast it yourself
    sliceFloat64 := []float64{0.1, 0.1, 0.2, 0.2, 0.3, 0.3}
    dup := dedupe.NewDeduplication() // or &dedupe.Deduplication{NotChange: true}
    dup.Do(&sliceFloat64)            // Be sure to pass it by address
    if src, err := dup.Float64(); err == nil {
        fmt.Printf("extracted slice -> %#v\n", src)
    }
    fmt.Printf("original  slice -> %#v\n", sliceFloat64)
}
```

### Result
```
[Case1]
original  slice -> []string{"C++", "Go", "Java", "Python", "Ruby", "V"}

[Case2]
extracted slice -> []float64{0.1, 0.2, 0.3}
original  slice -> []float64{0.1, 0.1, 0.2, 0.2, 0.3, 0.3}
```
Check [Go Playground](https://play.golang.org/p/QEUKQciGp4J)  
Check [Other patterns](https://github.com/go-utils/dedupe/blob/master/example/main.go)

# Support
`[]bool` `[]float32` `[]float64` `[]int` `[]int64` `[]uint` `[]uint64` `[]string`  
`[]struct` `[]*struct` `[]AnyType`

### No Support
`[]func`

# License
MIT
