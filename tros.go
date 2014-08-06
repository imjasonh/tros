package tros

import (
	"fmt"
	"reflect"
	"sort"
)

// SortInterface returns a sort.Interface suitable for sorting, reversing,
// searching, etc., a slice of structs by with the given field name fn
//
// i must be a slice of homogeneous structs. That is, all structs must have an
// exported field with the name fn, and those fields must all be of the same
// type.
//
// Operations using this sort.Interface are significantly (~10x) slower than the
// standard Go sorting idiom, because it relies heavily on reflection.
func SortInterface(i interface{}, fn string) (sort.Interface, error) {
	sval := reflect.ValueOf(i)
	if sval.Kind() != reflect.Slice {
		return nil, fmt.Errorf("non-slice interface, got %q", sval.Kind())
	}
	l := sval.Len()
	vals := make([]reflect.Value, l)
	fs := make([]reflect.Value, l)
	k := sval.Index(0).FieldByName(fn).Kind()
	if k == reflect.Invalid {
		return nil, fmt.Errorf("no field with name %q", fn)
	}
	if k > reflect.Float64 && k != reflect.String {
		return nil, fmt.Errorf("unsupported kind %q", k)
	}
	for i := 0; i < l; i++ {
		v := sval.Index(i)
		f := v.FieldByName(fn)
		if f.Kind() != k {
			return nil, fmt.Errorf("unmatched field kinds, %q vs %q", f.Kind(), k)
		}
		vals[i] = v
		fs[i] = f
	}
	return sort.Interface(sortable{vals, fs, k}), nil
}

// Sort sorts a slice of structs based on the values of the structs' fields with
// the given field name fn
//
// i must be a slice of homogeneous structs. That is, all structs must have an
// exported field with the name fn, and those fields must all be of the same
// type.
//
// Sort is significantly (~10x) slower than the standard Go sort.Sort function,
// because it relies heavily on reflection.
func Sort(i interface{}, fn string) error {
	s, err := SortInterface(i, fn)
	if err != nil {
		return err
	}
	sort.Sort(s)
	return nil
}

type sortable struct {
	vals []reflect.Value
	fs   []reflect.Value
	k    reflect.Kind
}

func (s sortable) Len() int { return len(s.vals) }

func (s sortable) Swap(i, j int) {
	a, b := s.vals[i], s.vals[j]
	tmp := reflect.New(a.Type()).Elem()
	tmp.Set(a)
	a.Set(b)
	b.Set(tmp)
}

func (s sortable) Less(i, j int) bool {
	af, bf := s.fs[i], s.fs[j]
	switch s.k {
	case reflect.Bool:
		return !af.Bool() && bf.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return af.Int() < bf.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return af.Uint() < bf.Uint()
	case reflect.Float32, reflect.Float64:
		return af.Float() < bf.Float()
	case reflect.String:
		return af.String() < bf.String()
	default:
		// Check in Sort should prevent this
		panic("unreachable")
	}
}
