package tros

/*
* Sample run:
* PASS
* BenchmarkTros        200           6813560 ns/op
* BenchmarkSort       1000           3184105 ns/op
* BenchmarkSlice       500           3687327 ns/op
 */

import (
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/bradfitz/slice"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sliceLen = 10000
)

type el struct {
	A string
}

func randSlice() []el {
	s := make([]el, sliceLen)
	for i := 0; i < len(s); i++ {
		s[i] = el{string(alphabet[rand.Intn(len(alphabet))])}
	}
	return s
}

func BenchmarkTros(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		Sort(randSlice(), "A")
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.Interface(iface(randSlice())))
	}
}

type iface []el

func (s iface) Len() int           { return len(s) }
func (s iface) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s iface) Less(i, j int) bool { return s[i].A < s[j].A }

func BenchmarkSlice(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		s := randSlice()
		slice.Sort(s, func(i, j int) bool {
			return s[i].A < s[j].A
		})
	}
}

func randStructSlice() []container {
	s := make([]container, sliceLen)
	for i := 0; i < len(s); i++ {
		s[i] = container{lenLesser{strings.Repeat("a", rand.Intn(10))}}
	}
	return s
}
func BenchmarkTrosStruct(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		Sort(randStructSlice(), "A")
	}
}

func BenchmarkSortStruct(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.Interface(cSlice(randStructSlice())))
	}
}

type cSlice []container

func (s cSlice) Len() int           { return len(s) }
func (s cSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s cSlice) Less(i, j int) bool { return len(s[i].A.val) < len(s[j].A.val) }

func BenchmarkSliceStruct(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		s := randStructSlice()
		slice.Sort(s, func(i, j int) bool {
			return len(s[i].A.val) < len(s[j].A.val)
		})
	}
}
