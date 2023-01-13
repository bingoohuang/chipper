package tests

import (
	"math"

	"github.com/bingoohuang/chipper/core"
)

type piTest struct {
	baseStepTest
}

// NewPiTest creates new π math.Pi test
func NewPiTest(n uint) core.Test {
	return &piTest{baseStepTest{n: n}}
}

func (f *piTest) Name() string {
	return f.nameBase(LeibnizPi)
}

func (f *piTest) Start() {
	f.leibnizPi(float64(f.n))

	f.isDone = true
}

func (f *piTest) leibnizPi(iter float64) float64 {
	var sum, i float64 = 0, 0
	for ; i < iter; i++ {
		sum += math.Pow(-1, i) / (2*i + 1)
		f.p++
	}
	return sum * 4
}
