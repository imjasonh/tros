// Package tros provides methods to sort slices of structs using reflection
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
// Operations using this sort.Interface are significantly (~2.5x) slower than the
// standard Go sorting idiom, because it relies heavily on reflection.
func SortInterface(i interface{}, fn string) (sort.Interface, error) {
	sval := reflect.ValueOf(i)
	if sval.Kind() != reflect.Slice {
		return nil, fmt.Errorf("non-slice interface, got %q", sval.Kind())
	}
	l := sval.Len()
	if l == 0 {
		return nil, fmt.Errorf("slice is empty")
	}

	if fn == "" {
		return nil, fmt.Errorf("must specify field name")
	}
	reverse := false
	if fn[0] == '-' {
		reverse = true
		fn = fn[1:]
	}

	vals := make([]reflect.Value, l)
	var fs []reflect.Value
	var ls []Lesser

	f0 := sval.Index(0).FieldByName(fn)
	k := f0.Kind()
	if k == reflect.Invalid {
		return nil, fmt.Errorf("no field with name %q", fn)
	}
	if !f0.CanSet() {
		return nil, fmt.Errorf("field %q is not exported", fn)
	}
	if k > reflect.Float64 && k != reflect.String && k != reflect.Struct {
		return nil, fmt.Errorf("unsupported kind %q", k)
	}
	for i := 0; i < l; i++ {
		v := sval.Index(i)
		f := v.FieldByName(fn)
		if f.Kind() != k {
			return nil, fmt.Errorf("unmatched field kinds, %q vs %q", f.Kind(), k)
		}
		if less, ok := f.Interface().(Lesser); ok {
			ls = append(ls, less)
		} else if f.Kind() == reflect.Struct {
			return nil, fmt.Errorf("struct field %q does not implement Lesser", fn)
		} else {
			fs = append(fs, f)
		}
		vals[i] = v
	}
	tmp := reflect.New(vals[0].Type()).Elem()
	return sort.Interface(&sortable{vals, k, tmp, fs, ls, reverse}), nil
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
	k    reflect.Kind
	tmp  reflect.Value   // reused for swapping
	fs   []reflect.Value // used for comparing structs by non-Lesser fields
	ls   []Lesser        // used for comparing structs by Lesser fields
	rev  bool            // if true, reverse sort order
}

func (s *sortable) Len() int { return len(s.vals) }

func (s *sortable) Swap(i, j int) {
	a, b := s.vals[i], s.vals[j]
	s.tmp.Set(a)
	a.Set(b)
	b.Set(s.tmp)

	if s.ls != nil {
		s.ls[i], s.ls[j] = s.ls[j], s.ls[i]
	}
}

func (s *sortable) Less(i, j int) bool {
	if s.ls != nil {
		return s.ls[i].Less(s.ls[j])
	}

	af, bf := s.fs[i], s.fs[j]
	var r bool
	switch s.k {
	case reflect.Bool:
		r = !af.Bool() && bf.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r = af.Int() < bf.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r = af.Uint() < bf.Uint()
	case reflect.Float32, reflect.Float64:
		r = af.Float() < bf.Float()
	case reflect.String:
		r = af.String() < bf.String()
	default:
		panic("unreachable: invalid Kind") // Check in Sort should prevent this
	}
	if s.rev {
		r = !r
	}
	return r
}

// Lesser is an interface used to define custom comparison logic.

// Fields implementing this interface may be used to sort structs using the
// field's implementation of Less.
type Lesser interface {
	Less(other Lesser) bool
}
