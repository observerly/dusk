package dusk

import (
	"math"
	"testing"
	"time"
)

func TestGetObliquityOfTheEclipticLawrence(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

	var T = GetCurrentJulianCenturyRelativeToJ2000(datetime)

	var got = GetObliquityOfTheEclipticLawrence(T)

	var want = 23.437992

	if math.Abs(got-want) > 0.0001 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetLunarMeanAnomalyLawrence(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got = GetLunarMeanAnomalyLawrence(datetime)

	var want = 85.910642

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEclipticPositionLawrenceX(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var Ωprime float64 = GetLunarCorrectedEclipticLongitudeOfTheAscendingNode(datetime)

	var λt float64 = GetLunarTrueEclipticLongitude(datetime)

	var got float64 = cosx(λt - Ωprime)

	var want float64 = -0.638869

	if math.Abs(got-want) > 0.1 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetLunarEclipticPositionLawrenceY(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var Ωprime float64 = GetLunarCorrectedEclipticLongitudeOfTheAscendingNode(datetime)

	var λt float64 = GetLunarTrueEclipticLongitude(datetime)

	// the inclination of the Moon's orbit with respect to the ecliptic
	var ι float64 = 5.1453964

	var got float64 = sinx(λt-Ωprime) * cosx(ι)

	var want float64 = -0.766215

	if math.Abs(got-want) > 0.1 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetLunarEclipticPositionLawrenceLongitudeQuadrant(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var Ωprime float64 = GetLunarCorrectedEclipticLongitudeOfTheAscendingNode(datetime)

	var λt float64 = GetLunarTrueEclipticLongitude(datetime)

	// the inclination of the Moon's orbit with respect to the ecliptic
	var ι float64 = 5.1453964

	var x float64 = cosx(λt - Ωprime)

	var y float64 = sinx(λt-Ωprime) * cosx(ι)

	// utilise atan2yx to determine a quadrant adjustment for arctan
	var got float64 = math.Mod(atan2yx(y, x), 360)

	// correct for negative angles
	if got < 0 {
		got += 360
	}

	var want = 230.178711

	if math.Abs(got-want) > 0.25 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetLunarEclipticPositionLawrenceLongitude(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var ec EclipticCoordinate = GetLunarEclipticPositionLawrence(datetime)

	var got float64 = ec.λ

	var want float64 = 65.059853

	if math.Abs(got-want) > 0.25 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetLunarEclipticPositionLawrenceLatitude(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var ec EclipticCoordinate = GetLunarEclipticPositionLawrence(datetime)

	var got float64 = ec.β

	var want float64 = -3.956258

	if math.Abs(got-want) > 0.1 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEquatorialPositionLawrenceRightAscension(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = eq.RightAscension

	var want float64 = 63.86571

	if math.Abs(got-want) > 2 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetLunarEquatorialPositionLawrenceDeclination(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = eq.Declination

	var want float64 = 17.248880

	if math.Abs(got-want) > 1 {
		t.Errorf("quad %f, wanted %f and difference %f", got, want, math.Abs(got-want))
	}
}

func TestGetSolarMeanAnomalyLawrence(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 2, 5, 17, 0, 0, 0, time.UTC)

	var got = GetSolarMeanAnomalyLawrence(datetime)

	var want = 32.592589

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarMeanAnomalyLawrenceAlt(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got = GetSolarMeanAnomalyLawrence(datetime)

	var want = 358.505618

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarEquationOfCenterLawrence(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 2, 5, 17, 0, 0, 0, time.UTC)

	var M = GetSolarMeanAnomalyLawrence(datetime)

	var got = GetSolarEquationOfCenterLawrence(M)

	var want = 1.031320

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarEclipticLongitudeLawrence(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 2, 5, 17, 0, 0, 0, time.UTC)

	var M = GetSolarMeanAnomalyLawrence(datetime)

	var C = GetSolarEquationOfCenterLawrence(M)

	var got = GetSolarEclipticLongitudeLawrence(M, C)

	var want = 316.562255

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetSolarEclipticLongitudeLawrenceAlt(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var M = GetSolarMeanAnomalyLawrence(datetime)

	var C = GetSolarEquationOfCenterLawrence(M)

	var got = GetSolarEclipticLongitudeLawrence(M, C)

	var want = 281.394034

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
