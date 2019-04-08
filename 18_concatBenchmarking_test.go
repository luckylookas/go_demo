package main

import (
	"testing"
	"gotest.tools/assert"
)

/*

go tool pprof cpu.out
$ top
go tool pprof mem.out
$ top
go test -bench=. -benchmem -memprofile mem.out -cpuprofile cpu.out
go test -bench=. -cpu 1,2,4

 */

func Test_concatSPrintf(t *testing.T) {
	val := concatSprintf("abc", "def")
	assert.Assert(t, val == "abcdef")
}

func Test_concatSlice(t *testing.T) {
	val := concatSlice("abc", "def")
	assert.Assert(t, val == "abcdef")
}

func Test_concatBuilder(t *testing.T) {
	val := concatBuilder("abc", "def")
	assert.Assert(t, val == "abcdef")
}

func Test_concatBuilderGrow(t *testing.T) {
	val := concatBuilderGrow("abc", "def")
	assert.Assert(t, val == "abcdef")
}

func Test_concatAppend(t *testing.T) {
	val := concatAppend("abc", "def")
	assert.Assert(t, val == "abcdef")
}

func Benchmark_concatSprintf(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concatSprintf("abc", "def")
	}
}

func Benchmark_concatSlice(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concatSlice("abc", "def")
	}
}

func Benchmark_concatBuilder(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concatBuilder("abc", "def")
	}
}

func Benchmark_concatBuilderGrow(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concatBuilderGrow("abc", "def")
	}
}

func Benchmark_concatAppend(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concatAppend("abc", "def")
	}
}

//beware of misleading ns/op values here!
func Benchmark_concatAppendParalell(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			concatAppend("abc", "def")
		}
	})
}