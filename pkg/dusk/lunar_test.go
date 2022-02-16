package dusk

import (
	"math"
	"testing"
)

func TestGetLunarMeanLongitude(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarMeanLongitude(J)

	var want float64 = 134.290182

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
func TestGetLunarMeanAnomaly(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarMeanAnomaly(J)

	var want float64 = 5.150833

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarArgumentOfLatitude(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarArgumentOfLatitude(J)

	var want float64 = 219.889721

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarHorizontalLongitude(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var M float64 = GetLunarMeanAnomaly(J)

	var L float64 = GetLunarMeanLongitude(J)

	var got float64 = GetLunarHorizontalLongitude(M, L)

	var want float64 = 134.854795

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
