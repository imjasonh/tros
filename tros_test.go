package tros

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type s struct {
		A, B string
	}
	for _, c := range []struct {
		ss  []s
		fn  string
		exp []s
	}{{
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
		"B",
		[]s{{"c", "a"}, {"b", "b"}, {"a", "c"}},
	}, {
		[]s{{"c", "a"}, {"b", "b"}, {"a", "c"}},
		"A",
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
	}, {
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
		"A",
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
	}} {
		if err := Sort(c.ss, c.fn); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(c.ss, c.exp) {
			t.Errorf("unexpected result\n got %v\nwant %v", c.ss, c.exp)
		}
	}
}

func TestSort_Bool(t *testing.T) {
	type e struct {
		A bool
	}
	l := []e{e{true}, e{true}, e{false}, e{true}, e{false}}
	if err := Sort(l, "A"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got, want := l, []e{e{false}, e{false}, e{true}, e{true}, e{true}}; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected result, got %v want %v", got, want)
	}
}

// Demonstrates sorting slices of structs on fields that implements Lesser.
func TestSort_Lesser(t *testing.T) {
	for _, c := range []struct {
		s, exp interface{}
	}{{
		[]container{
			{lenLesserStruct{"xxx"}},
			{lenLesserStruct{"z"}},
			{lenLesserStruct{"wwwww"}},
			{lenLesserStruct{"yy"}},
			{lenLesserStruct{"z"}},
		},
		[]container{
			{lenLesserStruct{"z"}},
			{lenLesserStruct{"z"}},
			{lenLesserStruct{"yy"}},
			{lenLesserStruct{"xxx"}},
			{lenLesserStruct{"wwwww"}},
		},
	}, {
		// This case is different than the above case because if
		// lenLesserString didn't implement Less then the result would
		// be a slice sorted using string-sorting instead of by length.
		[]container2{
			{lenLesserString("xxx")},
			{lenLesserString("z")},
			{lenLesserString("wwwww")},
			{lenLesserString("yy")},
			{lenLesserString("z")},
		},
		[]container2{
			{lenLesserString("z")},
			{lenLesserString("z")},
			{lenLesserString("yy")},
			{lenLesserString("xxx")},
			{lenLesserString("wwwww")},
		},
	}} {
		if err := Sort(c.s, "A"); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(c.s, c.exp) {
			t.Errorf("unexpected result\n got %v\nwant %v", c.s, c.exp)
		}
	}
}

type container struct {
	A lenLesserStruct
}

type lenLesserStruct struct {
	val string
}

func (l lenLesserStruct) Less(o Lesser) bool {
	ol, ok := o.(lenLesserStruct)
	if !ok {
		panic("other is not lenLesserStruct")
	}
	return len(l.val) < len(ol.val)
}

type container2 struct {
	A lenLesserString
}

type lenLesserString string

func (l lenLesserString) Less(o Lesser) bool {
	ol, ok := o.(lenLesserString)
	if !ok {
		panic("other is not lenLesserString")
	}
	return len(l) < len(ol)
}
