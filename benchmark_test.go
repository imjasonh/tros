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

func randLetter() string {
	return string(alphabet[rand.Intn(len(alphabet))])
}

func randSlice() []el {
	s := make([]el, sliceLen)
	for i := 0; i < sliceLen; i++ {
		s[i] = el{randLetter()}
	}
	return s
}

func BenchmarkTros(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		s := randSlice()
		Sort(s, "A")
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(1)
	for i := 0; i < b.N; i++ {
		s := sort.Interface(iface(randSlice()))
		sort.Sort(s)
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
