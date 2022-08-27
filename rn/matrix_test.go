package rn

import (
	"math"
	"testing"
)

func TestMatGet(t *testing.T) {
	for _, test := range []struct {
		m    Mat
		i, j int
		want float64
	}{
		{Mat{M: 1, N: 2, Data: []float64{1, 2}}, 0, 0, 1},
		{Mat{M: 1, N: 2, Data: []float64{1, 2}}, 0, 1, 2},
		{Mat{M: 2, N: 1, Data: []float64{1, 2}}, 0, 0, 1},
		{Mat{M: 2, N: 1, Data: []float64{1, 2}}, 1, 0, 2},

		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 0, 0, 1},
		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 1, 0, 2},
		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 0, 1, 3},
		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 1, 1, 4},

		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 0, 1},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 0, 2},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 1, 3},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 1, 4},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 2, 5},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 2, 6},

		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 0, 1},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 0, 2},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 2, 0, 3},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 1, 4},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 1, 5},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 2, 1, 6},

		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 0, 1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, 0, 2},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2, 0, 3},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 1, 4},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, 1, 5},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2, 1, 6},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 2, 7},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, 2, 8},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2, 2, 9},
	} {
		got := test.m.Get(test.i, test.j)
		if got != test.want {
			t.Errorf(
				"error:\ngot=%v\nwant=%v",
				got, test.want,
			)
		}
	}
}

func TestMatSet(t *testing.T) {
	for _, test := range []struct {
		m    Mat
		i, j int
		val  float64
	}{
		{Mat{M: 1, N: 2, Data: []float64{1, 2}}, 0, 0, -1},
		{Mat{M: 1, N: 2, Data: []float64{1, 2}}, 0, 1, -1},
		{Mat{M: 2, N: 1, Data: []float64{1, 2}}, 0, 0, -1},
		{Mat{M: 2, N: 1, Data: []float64{1, 2}}, 1, 0, -1},

		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 0, 0, -1},
		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 1, 0, -1},
		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 0, 1, -1},
		{Mat{M: 2, N: 2, Data: []float64{1, 2, 3, 4}}, 1, 1, -1},

		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 0, -1},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 0, -1},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 1, -1},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 1, -1},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 2, -1},
		{Mat{M: 2, N: 3, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 2, -1},

		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 0, -1},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 0, -1},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 2, 0, -1},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 0, 1, -1},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 1, 1, -1},
		{Mat{M: 3, N: 2, Data: []float64{1, 2, 3, 4, 5, 6}}, 2, 1, -1},

		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 0, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, 0, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2, 0, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 1, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, 1, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2, 1, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0, 2, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, 2, -1},
		{Mat{M: 3, N: 3, Data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2, 2, -1},
	} {
		test.m.Set(test.i, test.j, test.val)
		got := test.m.Get(test.i, test.j)
		if got != test.val {
			t.Errorf(
				"error:\ngot=%v\nwant=%v",
				got, test.val,
			)
		}
	}
}

func TestMatGaussSolve(t *testing.T) {
	for _, test := range []struct {
		m    Mat
		want Vec
	}{
		{
			Mat{M: 3, N: 4, Data: []float64{16, 81, 256, 4, 9, 16, 1, 1, 1, -3, 3.25, 33}},
			Vec{N: 3, X: []float64{0.25, -2, 1}},
		},
		{
			Mat{M: 3, N: 4, Data: []float64{2, -3, 2, 10, -6, 4, 4, -5, 6, -14, -4, 8}},
			Vec{N: 3, X: []float64{4, -3, 2}},
		},
		{
			Mat{M: 3, N: 4, Data: []float64{1, 2, 3, -2, -1, -3, 3, 4, 2, 1, 1, 2}},
			Vec{N: 3, X: []float64{0.3333333333334, -0.3333333333333, 0}},
		},
		{
			Mat{M: 3, N: 4, Data: []float64{1, -2, -1, 2, -1, -5, -3, 6, 3, 1, 4, -7}},
			Vec{N: 3, X: []float64{math.NaN(), 2, math.NaN()}},
		},
		{
			Mat{M: 4, N: 5, Data: []float64{3, 7, 8, 12, 4, -8, 9, 0, 11, -6, 20, 50, 1, 9, -1, 1, 20, 11, 35, 64}},
			Vec{N: 4, X: []float64{1, 1, 1, 2}},
		},
		{
			Mat{M: 4, N: 5, Data: []float64{3, 2, 20, 12, 4, 7, -1, 7, 2, -8, 0, -5, 8, 6, 8, -6, 6, 3, 35, 64}},
			Vec{N: 4, X: []float64{2.8508439545298, 3.3033643357447, 1.4737627741417, -2.3391893443564}},
		},
		{
			Mat{M: 4, N: 5, Data: []float64{
				0.537, 0.306, 0.902, 0.622,
				0.759, 0.254, 0.144, 0.015,
				0.778, 0.024, 0.148, 0.446,
				0.758, 0.562, 0.668, 0.821,
				0.729, 0.853, 0.439, 0.383}},
			Vec{N: 4, X: []float64{-0.5556995200839, 0.9762349711047, -1.017857257634, 1.4226137805486}},
		},
	} {
		_, gotX := test.m.GaussSolve()
		if !gotX.Equal(test.want) {
			t.Errorf(
				"error:\ngot=%v\nwant=%v",
				gotX.X, test.want.X,
			)
		}
	}
}
