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

func TestGetArgumentOfLocalSiderealTime(t *testing.T) {
	var got float64 = GetArgumentOfLocalSiderealTimeForTransit(latitude, 7.4070639)

	var want float64 = 92.682420

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if got < 0 || got > 360 {
		t.Errorf("The argument of LST must be an angle between 0° and 360°")
	}
}

func TestGetAtmosphericRefraction(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639})

	var got float64 = GetAtmosphericRefraction(hz.Altitude)

	var want float64 = 0.005219

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if got < 0 || got > 1.0 {
		t.Errorf("The atmospheric refraction must be between 0.0 and 0.5")
	}
}
