package dusk

import (
	"math"
	"testing"
)

func TestGetLunarMeanAnomaly(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarMeanAnomaly(J)

	var want float64 = 5.150833

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
