package gcd

import (
	"testing"
)

type st struct {
	n []uint64
	r uint64
}

func TestGCDEuclid(t *testing.T) {
	tt := []st{
		{n: []uint64{0, 2}, r: 2},
		{n: []uint64{2, 0}, r: 2},
		{n: []uint64{2, 2}, r: 2},
		{n: []uint64{2, 4}, r: 2},
		{n: []uint64{8, 12}, r: 4},
		{n: []uint64{12, 28}, r: 4},
	}
	for _, tc := range tt {
		g := Euclid(tc.n[0], tc.n[1])
		if g != tc.r {
			t.Errorf("%v: expected %v, got %v", tc.n, tc.r, g)
			return
		}
	}
}
