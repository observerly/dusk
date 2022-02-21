package dusk

import (
	"math"
	"testing"
)

func TestSinX(t *testing.T) {
	var got float64 = sinx(45)

	var want float64 = 0.70710678118

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestCosX(t *testing.T) {
	var got float64 = cosx(45)

	var want float64 = 0.70710678118

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestSinCosX(t *testing.T) {
	var sinGot, cosGot = sincosx(45)

	var sinWant = 0.70710678118

	var cosWant = 0.70710678118

	if math.Abs(sinGot-sinWant) > 0.00001 {
		t.Errorf("got %f, wanted %f", sinGot, sinWant)
	}

	if math.Abs(cosGot-cosWant) > 0.00001 {
		t.Errorf("got %f, wanted %f", cosWant, cosWant)
	}
}

func TestTanX(t *testing.T) {
	var got float64 = tanx(45)

	var want float64 = 1

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestASinX(t *testing.T) {
	var got float64 = asinx(45)

	var want float64 = 0.903339111

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestACosX(t *testing.T) {
	var got float64 = acosx(45)

	var want float64 = 0.667457216

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestATanX(t *testing.T) {
	var got float64 = atanx(45)

	var want float64 = 88.726970

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestATan2YX(t *testing.T) {
	var got float64 = atan2yx(45, 45)

	var want float64 = 45

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
