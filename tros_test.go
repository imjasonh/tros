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
		err error
	}{{
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
		"B",
		[]s{{"c", "a"}, {"b", "b"}, {"a", "c"}},
		nil,
	}, {
		[]s{{"c", "a"}, {"b", "b"}, {"a", "c"}},
		"A",
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
		nil,
	}, {
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
		"A",
		[]s{{"a", "c"}, {"b", "b"}, {"c", "a"}},
		nil,
	}} {
		if err := Sort(c.ss, c.fn); err != c.err {
			t.Errorf("unexpected error: %v", err)
		}
		if c.err != nil && !reflect.DeepEqual(c.ss, c.exp) {
			t.Errorf("unexpected result\n got %v\nwant %v", c.ss, c.exp)
		}
	}

}
