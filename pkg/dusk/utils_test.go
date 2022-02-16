package dusk

import (
	"math"
	"testing"
)

func TestGetEarthObliquity(t *testing.T) {
	var got float64 = GetEarthObliquity()

	var want float64 = 23.4397

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
