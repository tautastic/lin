package rn

import "testing"

func TestLineFrom(t *testing.T) {
	for _, test := range []struct {
		p, q Vec
		want Line
	}{
		{Vec{2, []float64{0, 0}}, Vec{2, []float64{1, 0}}, Line{Vec{2, []float64{0, 0}}, Vec{2, []float64{1, 0}}}},
		{Vec{2, []float64{0, 0}}, Vec{2, []float64{1, 1}}, Line{Vec{2, []float64{0, 0}}, Vec{2, []float64{1, 1}}}},
		{Vec{2, []float64{8, 7}}, Vec{2, []float64{8, 5}}, Line{Vec{2, []float64{8, 7}}, Vec{2, []float64{0, -2}}}},
		{Vec{2, []float64{1, 0}}, Vec{2, []float64{0, 0}}, Line{Vec{2, []float64{1, 0}}, Vec{2, []float64{-1, 0}}}},
		{Vec{2, []float64{8, 7}}, Vec{2, []float64{0, 5}}, Line{Vec{2, []float64{8, 7}}, Vec{2, []float64{-8, -2}}}},
		{Vec{2, []float64{2, 2}}, Vec{2, []float64{1, 1}}, Line{Vec{2, []float64{2, 2}}, Vec{2, []float64{-1, -1}}}},

		{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 0, 0}}, Line{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 0, 0}}}},
		{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 1, 1}}, Line{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 1, 1}}}},
		{Vec{3, []float64{8, 0, 7}}, Vec{3, []float64{8, 6, 5}}, Line{Vec{3, []float64{8, 0, 7}}, Vec{3, []float64{0, 6, -2}}}},
		{Vec{3, []float64{1, 0, 0}}, Vec{3, []float64{0, 0, 0}}, Line{Vec{3, []float64{1, 0, 0}}, Vec{3, []float64{-1, 0, 0}}}},
		{Vec{3, []float64{8, 0, 7}}, Vec{3, []float64{0, 6, 5}}, Line{Vec{3, []float64{8, 0, 7}}, Vec{3, []float64{-8, 6, -2}}}},
		{Vec{3, []float64{2, 2, 2}}, Vec{3, []float64{1, 1, 1}}, Line{Vec{3, []float64{2, 2, 2}}, Vec{3, []float64{-1, -1, -1}}}},
	} {
		got := LineFromTo(test.p, test.q)
		if !VecEqual(got.A, test.want.A) || !VecEqual(got.B, test.want.B) {
			t.Errorf("error:\ngot:  %v\nwant: %v", got, test.want)
		}
	}
}
