/*
Package dedupe - Deduplication among various types of slices
*/
package dedupe

import (
	"reflect"

	"golang.org/x/xerrors"
)

// Deduplication duplicate exclusion
//  supports type:
//    []bool, []float32, []float64, []int, []int64,
//    []uint, []uint64, []string, []struct, []*struct, []AnyType
//  no supports type: []func
type Deduplication struct {
	SliceBool    []bool
	SliceFloat32 []float32
	SliceFloat64 []float64
	SliceInt     []int
	SliceInt64   []int64
	SliceUint    []uint
	SliceUint64  []uint64
	SliceString  []string
	SliceAny     interface{}

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
// NotChange field: true
func NewDeduplication() *Deduplication {
	return &Deduplication{NotChange: true}
}

func (d *Deduplication) validation(args interface{}) error {
	switch reflect.TypeOf(args).Kind() {
	case reflect.Ptr:
		switch reflect.TypeOf(args).Elem().Kind() {
		case reflect.Slice, reflect.Array, reflect.Struct:
			if d.typeElem(reflect.ValueOf(args).Type()).Kind() == reflect.Func {
				return xerrors.New("function is not supported")
			}
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

func (d *Deduplication) typeElem(src reflect.Type) reflect.Type {
	el := src.Elem()
	switch el.Kind() {
	case reflect.Array, reflect.Slice:
		return d.typeElem(el)
	case reflect.Ptr:
		return d.typeElem(el)
	default:
		return el
	}
}

func (d *Deduplication) valueElem(src reflect.Value) reflect.Value {
	el := src.Elem()
	switch el.Kind() {
	case reflect.Ptr:
		return d.valueElem(el)
	default:
		return el
	}
}

func (d *Deduplication) duplication() reflect.Value {
	slice := reflect.MakeSlice(d.Value.Type(), 0, 0)
	encountered := make(map[interface{}]struct{}, d.Value.Len())
	for i := 0; i < d.Value.Len(); i++ {
		element := d.Value.Index(i)
		key := element.Interface()
		if element.Kind() == reflect.Ptr {
			key = d.valueElem(element).Interface()
		}
		if _, ok := encountered[key]; !ok {
			encountered[key] = struct{}{}
			slice = reflect.Append(slice, element)
		}
	}

	switch d.Value.Type().Elem().Kind() {
	case reflect.Bool:
		item, ok := slice.Interface().([]bool)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceBool = item
		}
	case reflect.Float32:
		item, ok := slice.Interface().([]float32)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceFloat32 = item
		}
	case reflect.Float64:
		item, ok := slice.Interface().([]float64)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceFloat64 = item
		}
	case reflect.Int:
		item, ok := slice.Interface().([]int)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceInt = item
		}
	case reflect.Int64:
		item, ok := slice.Interface().([]int64)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceInt64 = item
		}
	case reflect.Uint:
		item, ok := slice.Interface().([]uint)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceUint = item
		}
	case reflect.Uint64:
		item, ok := slice.Interface().([]uint64)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceUint64 = item
		}
	case reflect.String:
		item, ok := slice.Interface().([]string)
		if ok {
			d.Error = xerrors.Errorf("cast failed. typecast yourself using `Any()`")
			d.SliceString = item
		}
	}

	d.SliceAny = slice.Interface()
	return slice
}

// Do divide processing by type
func (d *Deduplication) Do(args interface{}) error {
	d.clear()
	if err := d.validation(args); err != nil {
		return xerrors.Errorf("error in validation method: %w", err)
	}

	var value reflect.Value
	switch args.(type) {
	case *[]bool, *[]int, *[]int64, *[]uint, *[]uint64,
		*[]float32, *[]float64, *[]string:
		value = d.duplication()
	case []bool, []int, []int64, []uint, []uint64,
		[]float32, []float64, []string:
		return xerrors.New("please pass by address")
	default:
		value = d.duplication()
		if d.Error != nil {
			return xerrors.Errorf("error in duplication method: %w", d.Error)
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
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceFloat32, nil
}

// Float64 returns the deduplicated slice
func (d *Deduplication) Float64() ([]float64, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Float64); err != nil {
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceFloat64, nil
}

// Int returns the deduplicated slice
func (d *Deduplication) Int() ([]int, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Int); err != nil {
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceInt, nil
}

// Int64 returns the deduplicated slice
func (d *Deduplication) Int64() ([]int64, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Int64); err != nil {
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceInt64, nil
}

// Uint returns the deduplicated slice
func (d *Deduplication) Uint() ([]uint, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Uint); err != nil {
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceUint, nil
}

// Uint64 returns the deduplicated slice
func (d *Deduplication) Uint64() ([]uint64, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.Uint64); err != nil {
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceUint64, nil
}

// String returns the deduplicated slice
func (d *Deduplication) String() ([]string, error) {
	defer d.clear()
	if err := d.errorCheck(reflect.String); err != nil {
		return nil, xerrors.Errorf("error in errorCheck method: %w", err)
	}
	return d.SliceString, nil
}

// Any returns the deduplicated slice
func (d *Deduplication) Any() (interface{}, error) {
	defer d.clear()
	if d.Error != nil {
		return nil, xerrors.Errorf("error in Deduplication.Error field: %w", d.Error)
	}
	return d.SliceAny, nil
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
