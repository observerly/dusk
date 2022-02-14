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

var latitude float64 = 19.798484

var elevation float64 = 0

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

func TestGetSolarDeclination(t *testing.T) {
	var J float64 = GetMeanSolarTime(d, longitude)

	var M float64 = GetSolarMeanAnomaly(J)

	var C float64 = GetSolarEquationOfCenter(M)

	var λ float64 = GetSolarEclipticLongitude(M, C)

	var got float64 = GetSolarDeclination(λ)

	var want float64 = 9.084711

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarHourAngle(t *testing.T) {
	var J float64 = GetMeanSolarTime(d, longitude)

	var M float64 = GetSolarMeanAnomaly(J)

	var C float64 = GetSolarEquationOfCenter(M)

	var λ float64 = GetSolarEclipticLongitude(M, C)

	var δ float64 = GetSolarDeclination(λ)

	var got float64 = GetSolarHourAngle(δ, latitude, elevation)

	var want float64 = 94.090408

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSunriseSunsetTimesRise(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	var sun Sun = GetSunriseSunsetTimesInUTC(d, longitude, latitude, elevation)

	var got time.Time = sun.rise.In(timezone)

	var want = time.Date(1992, 4, 12, 6, 05, 49, 72323712, timezone)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetSunriseSunsetTimesInUTCNoon(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	var sun Sun = GetSunriseSunsetTimesInUTC(d, longitude, latitude, elevation)

	var got time.Time = sun.noon.In(timezone)

	var want = time.Date(1992, 4, 12, 12, 22, 10, 770278016, timezone)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetSunriseSunsetTimesInUTCSet(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	var sun Sun = GetSunriseSunsetTimesInUTC(d, longitude, latitude, elevation)

	var got time.Time = sun.set.In(timezone)

	var want = time.Date(1992, 4, 12, 18, 38, 32, 468232192, timezone)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
