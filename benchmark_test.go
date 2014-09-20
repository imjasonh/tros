package tros

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sliceLen = 1000
)

/// Regular field comparison

type el struct {
	A string
}

func randSlice() []el {
	s := make([]el, sliceLen)
	for i := 0; i < len(s); i++ {
		s[i] = el{strings.Repeat(
			string(alphabet[rand.Intn(len(alphabet))]),
			rand.Intn(10))}
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

/// Lesser field comparison

type el2 struct {
	A lenLesserString
}

func randLesserSlice() []el2 {
	s := make([]el2, sliceLen)
	for i := 0; i < len(s); i++ {
		s[i] = el2{lenLesserString(strings.Repeat(
			string(alphabet[rand.Intn(len(alphabet))]),
			rand.Intn(10)))}
	}
	return s
}

func BenchmarkTrosLesser(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		Sort(randLesserSlice(), "A")
	}
}

func BenchmarkSortLesser(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.Interface(iface2(randLesserSlice())))
	}
}

type iface2 []el2

func (s iface2) Len() int           { return len(s) }
func (s iface2) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s iface2) Less(i, j int) bool { return s[i].A.Less(s[j].A) }
