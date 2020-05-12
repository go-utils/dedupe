package dedupe_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/go-utils/dedupe"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func deDuplicate1(args []int) []int {
	results := make([]int, 0, len(args))
	encountered := make(map[int]bool, len(args))
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

func deDuplicate2(args []int) []int {
	results := make([]int, 0, len(args))
	for i := 0; i < len(args); i++ {
		dup := false
		for j := 0; j < len(results); j++ {
			if args[i] == results[j] {
				dup = true
				break
			}
		}
		if !dup {
			results = append(results, args[i])
		}
	}
	return results
}

func deDuplicate3(args []int) []int {
	encountered := make(map[int]struct{}, len(args))

	results := make([]int, 0, len(args))

	for _, element := range args {
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			results = append(results, element)
		}
	}

	return results
}

type DeduplicateTest struct {
	Result []interface{}
	Type   reflect.Type
	Error  error
}

func (u DeduplicateTest) DeDuplicate(args interface{}) DeduplicateTest {
	switch reflect.TypeOf(args).Kind() {
	case reflect.Slice, reflect.Array:
		targetValue := reflect.ValueOf(args)
		encountered := make(map[interface{}]struct{}, targetValue.Len())
		u.Type = targetValue.Type()
		for i := 0; i < targetValue.Len(); i++ {
			element := targetValue.Index(i).Interface()
			if _, ok := encountered[element]; !ok {
				encountered[element] = struct{}{}
				u.Result = append(u.Result, element)
			}
		}
		return u
	default:
		u.Error = xerrors.New("invalid type")
		return u
	}
}

func (u DeduplicateTest) checkType(typ reflect.Kind) (err error) {
	if elem := u.Type.Elem(); elem.Kind() != typ {
		err = xerrors.Errorf("invalid Type -> %s", elem.String())
	}
	return
}

func (u DeduplicateTest) Int() (results []int, err error) {
	defer u.Clear()
	if err = u.checkType(reflect.Int); err != nil {
		return
	}
	size := len(u.Result)
	results = make([]int, 0, size)
	for i := 0; i < size; i++ {
		results = append(results, u.Result[i].(int))
	}
	return
}

func (u DeduplicateTest) Str() (results []string, err error) {
	defer u.Clear()
	if err = u.checkType(reflect.String); err != nil {
		return
	}
	size := len(u.Result)
	results = make([]string, 0, size)
	for i := 0; i < size; i++ {
		results = append(results, u.Result[i].(string))
	}
	return
}

func (u DeduplicateTest) Clear() {
	u.Result = nil
	u.Error = nil
}

func BenchmarkDeDuplicateInt1(b *testing.B) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, i%99000)
	}

	as := assert.New(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := deDuplicate1(data)
		as.Len(result, 99000)
	}
}

func BenchmarkDeDuplicateInt2(b *testing.B) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, i%99000)
	}

	as := assert.New(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := deDuplicate2(data)
		as.Len(result, 99000)
	}
}

func BenchmarkDeDuplicateInt3(b *testing.B) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, i%99000)
	}

	as := assert.New(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := deDuplicate3(data)
		as.Len(result, 99000)
	}
}

func BenchmarkDeDuplicateStrInterface1(b *testing.B) {
	data := make([]string, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, fmt.Sprintf("test-%08d", i%99000))
	}

	as := assert.New(b)
	b.ResetTimer()
	util := new(DeduplicateTest)
	for i := 0; i < b.N; i++ {
		result, _ := util.DeDuplicate(data).Str()
		as.Len(result, 99000)
	}
}

func BenchmarkDeDuplicateIntInterface1(b *testing.B) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, i%99000)
	}

	as := assert.New(b)
	b.ResetTimer()
	util := new(DeduplicateTest)
	for i := 0; i < b.N; i++ {
		result, _ := util.DeDuplicate(data).Int()
		as.Len(result, 99000)
	}
}

func BenchmarkDeDuplicateStrInterface2(b *testing.B) {
	data := make([]string, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, fmt.Sprintf("test-%08d", i%99000))
	}

	as := assert.New(b)
	b.ResetTimer()
	ded := dedupe.NewDeduplication()
	for i := 0; i < b.N; i++ {
		if err := ded.Do(&data); err != nil {
			log.Fatalln(err)
		}
		result, err := ded.String()
		if err != nil {
			log.Println(err)
		}
		as.Len(result, 99000)
	}
}

func BenchmarkDeDuplicateIntInterface2(b *testing.B) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, i%99000)
	}

	as := assert.New(b)
	b.ResetTimer()
	ded := dedupe.NewDeduplication()
	for i := 0; i < b.N; i++ {
		if err := ded.Do(&data); err != nil {
			log.Fatalln(err)
		}
		result, err := ded.Int()
		if err != nil {
			log.Println(err)
		}
		as.Len(result, 99000)
	}
}
