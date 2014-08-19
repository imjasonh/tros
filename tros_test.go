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
