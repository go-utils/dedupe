package dedupe

import (
	"fmt"
	"reflect"
	"sort"

	"golang.org/x/xerrors"
)

// Deduplication duplicate exclusion
//   └── Supports type: []bool, []float32, []float64, []int, []int64, []uint, []uint64, []string, []AnyStruct
type Deduplication struct {
	SliceBool    []bool
	SliceFloat32 []float32
	SliceFloat64 []float64
	SliceInt     []int
	SliceInt64   []int64
	SliceUint    []uint
	SliceUint64  []uint64
	SliceString  []string
	SliceStruct  interface{}

	NotChange bool
	Value     *reflect.Value
	Error     error
}

var defaultDeduplication = new(Deduplication)

// Do simply dedupe
func Do(args interface{}) error {
	return defaultDeduplication.Do(args)
}

// NewDeduplication constructor
func NewDeduplication() *Deduplication {
	return &Deduplication{NotChange: true}
}

func (d *Deduplication) validation(args interface{}) error {
	switch reflect.TypeOf(args).Kind() {
	case reflect.Ptr:
		switch reflect.TypeOf(args).Elem().Kind() {
		case reflect.Slice, reflect.Array, reflect.Struct:
			targetValue := reflect.Indirect(reflect.ValueOf(args))
			d.Value = &targetValue
		default:
			d.Error = xerrors.New("invalid type")
		}
	default:
		d.Error = xerrors.New("please pass by address")
	}
	return d.Error
}

func (d *Deduplication) duplicationBool() reflect.Value {
	encountered := make(map[bool]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i).Bool()
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceBool = append(d.SliceBool, element)
		}
	}
	sort.Slice(d.SliceBool, func(i, j int) bool { return fmt.Sprint(d.SliceBool[i]) < fmt.Sprint(d.SliceBool[j]) })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]bool{}), len(d.SliceBool), len(d.SliceBool))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceBool))
	return newSlice
}

func (d *Deduplication) duplicationFloat32() reflect.Value {
	encountered := make(map[float32]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := float32(d.Value.Index(i).Float())
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceFloat32 = append(d.SliceFloat32, element)
		}
	}
	sort.Slice(d.SliceFloat32, func(i, j int) bool { return d.SliceFloat32[i] < d.SliceFloat32[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]float32{}), len(d.SliceFloat32), len(d.SliceFloat32))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceFloat32))
	return newSlice
}

func (d *Deduplication) duplicationFloat64() reflect.Value {
	encountered := make(map[float64]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i).Float()
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceFloat64 = append(d.SliceFloat64, element)
		}
	}
	sort.Slice(d.SliceFloat64, func(i, j int) bool { return d.SliceFloat64[i] < d.SliceFloat64[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]float64{}), len(d.SliceFloat64), len(d.SliceFloat64))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceFloat64))
	return newSlice
}

func (d *Deduplication) duplicationInt() reflect.Value {
	encountered := make(map[int]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := int(d.Value.Index(i).Int())
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceInt = append(d.SliceInt, element)
		}
	}
	sort.Slice(d.SliceInt, func(i, j int) bool { return d.SliceInt[i] < d.SliceInt[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]int{}), len(d.SliceInt), len(d.SliceInt))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceInt))
	return newSlice
}

func (d *Deduplication) duplicationInt64() reflect.Value {
	encountered := make(map[int64]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i).Int()
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceInt64 = append(d.SliceInt64, element)
		}
	}
	sort.Slice(d.SliceInt64, func(i, j int) bool { return d.SliceInt64[i] < d.SliceInt64[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]int64{}), len(d.SliceInt64), len(d.SliceInt64))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceInt64))
	return newSlice
}

func (d *Deduplication) duplicationUint() reflect.Value {
	encountered := make(map[uint]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := uint(d.Value.Index(i).Uint())
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceUint = append(d.SliceUint, element)
		}
	}
	sort.Slice(d.SliceUint, func(i, j int) bool { return d.SliceUint[i] < d.SliceUint[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]uint{}), len(d.SliceUint), len(d.SliceUint))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceUint))
	return newSlice
}

func (d *Deduplication) duplicationUint64() reflect.Value {
	encountered := make(map[uint64]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i).Uint()
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceUint64 = append(d.SliceUint64, element)
		}
	}
	sort.Slice(d.SliceUint64, func(i, j int) bool { return d.SliceUint64[i] < d.SliceUint64[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]uint64{}), len(d.SliceUint64), len(d.SliceUint64))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceUint64))
	return newSlice
}

func (d *Deduplication) duplicationString() reflect.Value {
	encountered := make(map[string]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i).String()
		if _, ok := encountered[element]; !ok {
			encountered[element] = struct{}{}
			d.SliceString = append(d.SliceString, element)
		}
	}
	sort.Slice(d.SliceString, func(i, j int) bool { return d.SliceString[i] < d.SliceString[j] })
	newSlice := reflect.MakeSlice(reflect.TypeOf([]string{}), 0, len(d.SliceString))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(d.SliceString))
	return newSlice
}

func (d *Deduplication) duplicationStruct(args interface{}) reflect.Value {
	rv := reflect.Indirect(reflect.ValueOf(args))
	slice := reflect.MakeSlice(rv.Type(), 0, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i)
		for j := 0; j < slice.Len(); j++ {
			if !reflect.DeepEqual(element.Interface(), slice.Index(j).Interface()) {
				slice = reflect.Append(slice, element)
			}
		}
		if i == 0 {
			slice = reflect.Append(slice, element)
		}
	}
	d.SliceStruct = slice.Interface()
	return slice
}

// Do divide processing by type
func (d *Deduplication) Do(args interface{}) error {
	d.clear()
	if err := d.validation(args); err != nil {
		return xerrors.Errorf("error in validation: %w", err)
	}

	var value reflect.Value
	switch types := args.(type) {
	case *[]bool:
		value = d.duplicationBool()
	case *[]float32:
		value = d.duplicationFloat32()
	case *[]float64:
		value = d.duplicationFloat64()
	case *[]int:
		value = d.duplicationInt()
	case *[]int64:
		value = d.duplicationInt64()
	case *[]uint:
		value = d.duplicationUint()
	case *[]uint64:
		value = d.duplicationUint64()
	case *[]string:
		value = d.duplicationString()
	case []bool, []int, []int64, []uint, []uint64,
		[]float32, []float64, []string:
		return xerrors.New("please pass by address")
	default:
		switch d.Value.Type().Elem().Kind() {
		case reflect.Struct, reflect.Ptr:
			value = d.duplicationStruct(args)
		default:
			return xerrors.Errorf("invalid type: %#v", types)
		}
	}

	if !d.NotChange {
		renewal := reflect.Indirect(reflect.ValueOf(args))
		renewal.Set(value)
	}
	return nil
}

func (d *Deduplication) errorCheck(types ...reflect.Kind) (err error) {
	element := d.Value.Type().Elem()
	ok := d.typeCheck(element, types...)
	switch {
	case !ok:
		err = xerrors.Errorf("invalid Type -> %s", element.String())
	case d.Error != nil:
		err = d.Error
	case len(d.SliceFloat32) == 0 &&
		len(d.SliceFloat64) == 0 &&
		len(d.SliceInt) == 0 &&
		len(d.SliceInt64) == 0 &&
		len(d.SliceUint) == 0 &&
		len(d.SliceUint64) == 0 &&
		len(d.SliceString) == 0 &&
		len(types) == 1:
		err = xerrors.Errorf("0 size slice: %s", d.Value.Type().String())
	}
	return
}

func (d *Deduplication) typeCheck(element reflect.Type, types ...reflect.Kind) bool {
	for _, typ := range types {
		if typ == element.Kind() {
			return true
		}
	}
	return false
}

// Float32 returns the deduplicated slice
func (d *Deduplication) Float32() ([]float32, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Float32); err != nil {
		return nil, err
	}
	return d.SliceFloat32, nil
}

// Float64 returns the deduplicated slice
func (d *Deduplication) Float64() ([]float64, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Float64); err != nil {
		return nil, err
	}
	return d.SliceFloat64, nil
}

// Int returns the deduplicated slice
func (d *Deduplication) Int() ([]int, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Int); err != nil {
		return nil, err
	}
	return d.SliceInt, nil
}

// Int64 returns the deduplicated slice
func (d *Deduplication) Int64() ([]int64, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Int64); err != nil {
		return nil, err
	}
	return d.SliceInt64, nil
}

// Uint returns the deduplicated slice
func (d *Deduplication) Uint() ([]uint, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Uint); err != nil {
		return nil, err
	}
	return d.SliceUint, nil
}

// Uint64 returns the deduplicated slice
func (d *Deduplication) Uint64() ([]uint64, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Uint64); err != nil {
		return nil, err
	}
	return d.SliceUint64, nil
}

// String returns the deduplicated slice
func (d *Deduplication) String() ([]string, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.String); err != nil {
		return nil, err
	}
	return d.SliceString, nil
}

// Struct returns the deduplicated slice
func (d *Deduplication) Struct() (interface{}, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Struct, reflect.Ptr); err != nil {
		return nil, err
	}
	return d.SliceStruct, nil
}

func (d *Deduplication) clear() {
	d.SliceFloat32 = make([]float32, 0)
	d.SliceFloat64 = make([]float64, 0)
	d.SliceInt = make([]int, 0)
	d.SliceInt64 = make([]int64, 0)
	d.SliceUint = make([]uint, 0)
	d.SliceUint64 = make([]uint64, 0)
	d.SliceString = make([]string, 0)
	d.Value = new(reflect.Value)
	d.Error = nil
}
