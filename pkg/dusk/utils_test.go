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

	got := GetAtmosphericRefraction(hz.Altitude)

	var want float64 = 0.005219

	if math.Abs(*got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", *got, want)
	}

	if *got < 0 || *got > 1.0 {
		t.Errorf("The atmospheric refraction must be between 0.0 and 0.5")
	}
}

func TestGetAtmosphericRefractionBelowHorizon(t *testing.T) {
	got := GetAtmosphericRefraction(-45)

	if got != nil {
		t.Errorf("The atmospheric refraction must be nil below the observer's horizon")
	}
}

func TestGetRelativeAirMass(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639})

	got := GetRelativeAirMass(hz.Altitude)

	var want float64 = 1.046558

	if math.Abs(*got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", *got, want)
	}

	if *got < 1 || *got > 40.0 {
		t.Errorf("The relative air mass must be a value bewteen 1 and approx. 40 at the observer's horizon")
	}
}

func TestGetRelativeAirMassAtZenith(t *testing.T) {
	got := GetRelativeAirMass(90)

	var want float64 = 1.0

	if math.Abs(*got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", *got, want)
	}

	if *got < 1 || *got > 40.0 {
		t.Errorf("The relative air mass must be a value bewteen 1 and approx. 40 at the observer's horizon")
	}
}

func TestGetRelativeAirMassAtHorizon(t *testing.T) {
	got := GetRelativeAirMass(0)

	var want float64 = 38

	if math.Abs(*got-want) > 2 {
		t.Errorf("The relative air mass must be approximately ~37 - 39 at the observer's horizon")
	}

	if *got < 1 || *got > 40.0 {
		t.Errorf("The relative air mass must be a value bewteen 1 and approx. 40 at the observer's horizon")
	}
}

func TestGetRelativeAirMassBelowHorizon(t *testing.T) {
	got := GetRelativeAirMass(-1)

	if got != nil {
		t.Errorf("The relative air mass must be nil below the observer's horizon")
	}
}

func TestGetApparentAltitude(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639})

	got := GetApparentAltitude(hz.Altitude)

	if math.Abs(*got-hz.Altitude) > 0.0053 {
		t.Errorf("got %f, wanted %f", *got, hz.Altitude)
	}
}

func TestGetApparentAltitudeBelowHorizon(t *testing.T) {
	got := GetApparentAltitude(-45)

	if got != nil {
		t.Errorf("The apparent altitude must have no adjustment below the observer's horizon")
	}
}
