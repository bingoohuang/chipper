package tests

import "testing"

func benchmarkLeibnizPi(i int, b *testing.B) {
	pt := &piTest{}
	for n := 0; n < b.N; n++ {
		pt.leibnizPi(float64(i))
	}
}

func BenchmarkLeibnizPi1(b *testing.B)        { benchmarkLeibnizPi(1, b) }
func BenchmarkLeibnizPi10(b *testing.B)       { benchmarkLeibnizPi(10, b) }
func BenchmarkLeibnizPi100(b *testing.B)      { benchmarkLeibnizPi(100, b) }
func BenchmarkLeibnizPi1000(b *testing.B)     { benchmarkLeibnizPi(1000, b) }
func BenchmarkLeibnizPi10000(b *testing.B)    { benchmarkLeibnizPi(10000, b) }
func BenchmarkLeibnizPi100000(b *testing.B)   { benchmarkLeibnizPi(100000, b) }
func BenchmarkLeibnizPi1000000(b *testing.B)  { benchmarkLeibnizPi(1000000, b) }
func BenchmarkLeibnizPi10000000(b *testing.B) { benchmarkLeibnizPi(10000000, b) }
