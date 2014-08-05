package tros

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var r = rand.New(rand.NewSource(time.Now().Unix()))

type el struct {
	A, B string
}

func randLetter() string {
	return string(alphabet[r.Intn(len(alphabet))])
}

func genList(n int) []el {
	s := make([]el, n)
	for i := 0; i < n; i++ {
		s[i] = el{randLetter(), randLetter()}
	}
	return s
}

type iface []el

func (s iface) Len() int           { return len(s) }
func (s iface) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s iface) Less(i, j int) bool { return s[i].A < s[j].A }

func benchTros(l int, b *testing.B) {
	s := genList(l)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(s, "A")
	}
}

func BenchmarkTros10(b *testing.B)    { benchTros(10, b) }
func BenchmarkTros100(b *testing.B)   { benchTros(100, b) }
func BenchmarkTros1000(b *testing.B)  { benchTros(1000, b) }
func BenchmarkTros10000(b *testing.B) { benchTros(10000, b) }

func benchSort(l int, b *testing.B) {
	s := sort.Interface(iface(genList(l)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(s)
	}
}

func BenchmarkSort10(b *testing.B)    { benchSort(10, b) }
func BenchmarkSort100(b *testing.B)   { benchSort(100, b) }
func BenchmarkSort1000(b *testing.B)  { benchSort(1000, b) }
func BenchmarkSort10000(b *testing.B) { benchSort(10000, b) }
