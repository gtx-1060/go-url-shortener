package services

import (
	"testing"
)

func TestRandomFromString_collision(t *testing.T) {
	hmap := make(map[string]struct{})
	url := "https://stackoverflow.com/questions/"
	i := 0
	for {
		gen := RandomFromString(url)
		if _, ok := hmap[gen]; ok {
			t.Skipf("collision after %v iterations\n", i)
		}
		hmap[gen] = struct{}{}
		i++
	}
}

func TestRandomString_collision(t *testing.T) {
	hmap := make(map[string]struct{})
	i := 0
	for {
		gen := RandomString(8)
		if _, ok := hmap[gen]; ok {
			t.Skipf("collision after %v iterations\n", i)
		}
		hmap[gen] = struct{}{}
		i++
	}
}

func BenchmarkRandomFromString(b *testing.B) {
	url := "https://stackoverflow.com/questions/"
	for i := 0; i < b.N; i++ {
		_ = RandomFromString(url)
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomString(8)
	}
}
