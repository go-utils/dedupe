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
// Basic usage
sliceString := []string{"Go", "V", "C++", "Java", "Python", "Go", "Ruby", "C++", "Go", "V"}
dedupe.Do(sliceString)
fmt.Println(sliceString) // change original slice

// Extract duplicates without changing the original slice
// But: If it is a structure, you must cast it yourself
sliceFloat64 := []float64{0.1, 0.1, 0.2, 0.2, 0.3, 0.3}
dup.Do(&sliceFloat64)
fmt.Println(dup.Float64())
fmt.Println(sliceFloat64)
```
### Result
```
[C++ Go Java Python Ruby V]
```
Check [Go Playground](https://play.golang.org/p/vyphUn0Lx1E)

# Support
`[]bool` `[]float32` `[]float64` `[]int` `[]int64` `[]uint` `[]uint64` `[]string`

# License
MIT
