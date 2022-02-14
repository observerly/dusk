package dusk

import (
	"math"
	"testing"
	"time"
)

// For testing we need to specify a date because most calculations are
// differential w.r.t a time component. We set it to the date provided
// on p.342 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
// Va: Willmann - Bell.:
var d time.Time = time.Date(1992, 4, 12, 0, 0, 0, 0, time.UTC)

func TestGetSolarMeanAnomaly(t *testing.T) {
	var J float64 = GetMeanSolarTime(d, longitude)

	var got float64 = GetSolarMeanAnomaly(J)

	var want float64 = 98.561957

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarEquationOfCenter(t *testing.T) {
	var J float64 = GetMeanSolarTime(d, longitude)

	var M float64 = GetSolarMeanAnomaly(J)

	var got float64 = GetSolarEquationOfCenter(M)

	var want float64 = 1.887301

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarEclipticLongitude(t *testing.T) {
	var J float64 = GetMeanSolarTime(d, longitude)

	var M float64 = GetSolarMeanAnomaly(J)

	var C float64 = GetSolarEquationOfCenter(M)

	var got float64 = GetSolarEclipticLongitude(M, C)

	var want float64 = 383.386458

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarTransit(t *testing.T) {
	var J float64 = GetMeanSolarTime(d, longitude)

	var M float64 = GetSolarMeanAnomaly(J)

	var C float64 = GetSolarEquationOfCenter(M)

	var λ float64 = GetSolarEclipticLongitude(M, C)

	var got float64 = GetSolarTransitJulianDate(J, M, λ)

	var want float64 = 2448725.432069

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
