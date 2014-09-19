package tros

/*
* Sample run:
* PASS
* BenchmarkTros        200           7784816 ns/op
* BenchmarkSort        500           3147062 ns/op
 */

import (
	"math/rand"
	"sort"
	"testing"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sliceLen = 10000
)

type el struct {
	A string
}

func randLetter() string {
	return string(alphabet[rand.Intn(len(alphabet))])
}

func slice(n int) []el {
	s := make([]el, n)
	for i := 0; i < n; i++ {
		s[i] = el{randLetter()}
	}
	return s
}

func BenchmarkTros(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		s := slice(sliceLen)
		Sort(s, "A")
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		s := sort.Interface(iface(slice(sliceLen)))
		sort.Sort(s)
	}
}

type iface []el

func (s iface) Len() int           { return len(s) }
func (s iface) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s iface) Less(i, j int) bool { return s[i].A < s[j].A }
