package tros

import (
	"fmt"
	"reflect"
	"sort"
)

// Sort sorts a slice of structs based on the values of the structs' fields with the given field name fn
func Sort(i interface{}, fn string) error {
	sval := reflect.ValueOf(i)
	if sval.Kind() != reflect.Slice {
		return fmt.Errorf("non-slice interface, got %q", sval.Kind())
	}
	s := sortable{sval, fn, sval.Len(), nil}
	sort.Sort(sort.Interface(s))
	return s.err
}

type sortable struct {
	sval reflect.Value
	fn   string
	len  int
	err  error
}

func (s sortable) Len() int { return s.len }

func (s sortable) Swap(i, j int) {
	if s.err != nil {
		return
	}
	a, b := s.sval.Index(i), s.sval.Index(j)
	tmp := reflect.New(a.Type()).Elem()
	tmp.Set(a)
	a.Set(b)
	b.Set(tmp)
}

func (s sortable) Less(i, j int) bool {
	a, b := s.sval.Index(i), s.sval.Index(j)
	af, bf := a.FieldByName(s.fn), b.FieldByName(s.fn)
	afk, bfk := af.Kind(), bf.Kind()
	if afk != bfk {
		s.err = fmt.Errorf("unmatched kinds, %q vs %q", afk, bfk)
		return false
	}
	switch afk {
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
		s.err = fmt.Errorf("unsupported kind %q", afk)
		return false
	}
}
