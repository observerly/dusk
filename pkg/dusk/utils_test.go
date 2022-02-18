package dusk

import (
	"math"
	"testing"
	"time"
)

func TestGetEarthObliquity(t *testing.T) {
	var got float64 = GetEarthObliquity()

	var want float64 = 23.4397

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetMeanObliquityOfTheEcliptic(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.148 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
	// Va: Willmann - Bell.:
	var d time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetMeanObliquityOfTheEcliptic(J)

	var want float64 = 23.440947

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetNutationInLongitudeOfTheEcliptic(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.148 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
	// Va: Willmann - Bell.:
	var d time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var L float64 = GetSolarMeanLongitude(J)

	var l float64 = GetLunarMeanLongitude(J)

	var Ω float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var got float64 = GetNutationInLongitudeOfTheEcliptic(L, l, Ω)

	var want float64 = -0.000648203

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetNutationInObliquityOfTheEcliptic(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.148 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
	// Va: Willmann - Bell.:
	var d time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var L float64 = GetSolarMeanLongitude(J)

	var l float64 = GetLunarMeanLongitude(J)

	var Ω float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var got float64 = GetNutationInObliquityOfTheEcliptic(L, l, Ω)

	var want float64 = 0.002629996

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
