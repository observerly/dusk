package dusk

import (
	"math"
	"testing"
	"time"
)

func TestGetLunarMeanLongitude(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarMeanLongitude(J)

	var want float64 = 134.290182

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarMeanEclipticLongitude(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got = GetLunarMeanEclipticLongitude(datetime)

	var want = 59.716785

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarTrueEclipticLongitude(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got = GetLunarTrueEclipticLongitude(datetime)

	var want = 64.972240

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarMeanEclipticLongitudeOfTheAscendingNode(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got = GetLunarMeanEclipticLongitudeOfTheAscendingNode(datetime)

	var want = 194.877008

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarCorrectedEclipticLongitudeOfTheAscendingNode(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got = GetLunarCorrectedEclipticLongitudeOfTheAscendingNode(datetime)

	var want = 194.881180

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarMeanElongation(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarMeanElongation(J)

	var want float64 = 113.842304

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

func TestGetLunarTrueAnomaly(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var got float64 = GetLunarTrueAnomaly(datetime)

	var want float64 = 6.302688

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

func TestGetLunarHorizontalLatitude(t *testing.T) {
	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var F float64 = GetLunarArgumentOfLatitude(J)

	var got float64 = GetLunarHorizontalLatitude(F)

	var want float64 = 356.711352

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarLongitudeOfTheAscendingNode(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.148 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
	// Va: Willmann - Bell.:
	var d time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var got float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var want float64 = 11.253083

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarLongitudeOfNutation(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.148 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
	// Va: Willmann - Bell.:
	var d time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var L float64 = GetSolarMeanLongitude(J)

	var l float64 = GetLunarMeanLongitude(J)

	var Ω float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var got float64 = GetLunarLongitudeOfNutation(L, l, Ω)

	var want float64 = -0.000648203

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarObliquityOfNutation(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.148 of Meeus, Jean. 1991. Astronomical algorithms.Richmond,
	// Va: Willmann - Bell.:
	var d time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var J float64 = GetCurrentJulianCenturyRelativeToJ2000(d)

	var L float64 = GetSolarMeanLongitude(J)

	var l float64 = GetLunarMeanLongitude(J)

	var Ω float64 = GetLunarLongitudeOfTheAscendingNode(J)

	var got float64 = GetLunarObliquityOfNutation(L, l, Ω)

	var want float64 = 0.002629996

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarAnnualEquationCorrection(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var M float64 = GetSolarMeanAnomalyLawrence(datetime)

	var got float64 = GetLunarAnnualEquationCorrection(M)

	var want float64 = -0.004845

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEvectionCorrection(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var M float64 = GetLunarMeanAnomalyLawrence(datetime)

	var λ float64 = GetLunarMeanEclipticLongitude(datetime)

	var Msol = GetSolarMeanAnomalyLawrence(datetime)

	var Csol = GetSolarEquationOfCenterLawrence(Msol)

	var λsol float64 = GetSolarEclipticLongitudeLawrence(Msol, Csol)

	var got float64 = GetLunarEvectionCorrection(M, λ, λsol)

	var want float64 = -0.237282

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarMeanAnomalyCorrection(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 2, 3, 0, 0, 0, time.UTC)

	var M float64 = GetLunarMeanAnomalyLawrence(datetime)

	var λ float64 = GetLunarMeanEclipticLongitude(datetime)

	var Msol = GetSolarMeanAnomalyLawrence(datetime)

	var Csol = GetSolarEquationOfCenterLawrence(Msol)

	var λsol float64 = GetSolarEclipticLongitudeLawrence(Msol, Csol)

	var Ae float64 = GetLunarAnnualEquationCorrection(M)

	var Eν float64 = GetLunarEvectionCorrection(M, λ, λsol)

	var got float64 = GetLunarMeanAnomalyCorrection(M, Msol, Ae, Eν)

	var want float64 = 85.497682

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEquatorialPositionRightAscension(t *testing.T) {
	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = eq.α

	var want float64 = 76.239624

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEquatorialPositionDeclination(t *testing.T) {
	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = eq.δ

	var want float64 = 23.598793

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
func TestGetLunarEclipticPositionLongitude(t *testing.T) {
	var ec EclipticCoordinate = GetLunarEclipticPosition(d)

	var got float64 = ec.λ

	var want float64 = 133.162655

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEclipticPositionLongitudeAlt(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	var ec EclipticCoordinate = GetLunarEclipticPosition(datetime)

	var got float64 = ec.λ

	var want float64 = 50.604878

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEclipticPositionLatitude(t *testing.T) {
	var ec EclipticCoordinate = GetLunarEclipticPosition(d)

	var got float64 = ec.β

	var want float64 = -3.229126

	if math.Abs(got-want) > 0.1 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEclipticPositionLatitudeAlt(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	var ec EclipticCoordinate = GetLunarEclipticPosition(datetime)

	var got float64 = ec.β

	var want float64 = -2.981288

	if math.Abs(got-want) > 0.1 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEclipticPositionDistance(t *testing.T) {
	var ec EclipticCoordinate = GetLunarEclipticPosition(d)

	var got float64 = ec.Δ

	var want float64 = 368403.226858

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarHorizontalParallax(t *testing.T) {
	var ec EclipticCoordinate = GetLunarEclipticPosition(d)

	var got float64 = GetLunarHorizontalParallax(ec.Δ)

	var want float64 = 0.992007

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarHourAngle(t *testing.T) {
	var ec EclipticCoordinate = GetLunarEclipticPosition(datetime)

	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var π float64 = GetLunarHorizontalParallax(ec.Δ)

	var got float64 = GetLunarHourAngle(eq.δ, latitude, 0, π)

	var want float64 = 97.500858

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetEclipticLongitudeInXHours(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	var M float64 = GetLunarMeanAnomalyLawrence(datetime)

	var λ float64 = GetLunarMeanEclipticLongitude(datetime)

	var Msol = GetSolarMeanAnomalyLawrence(datetime)

	var Csol = GetSolarEquationOfCenterLawrence(Msol)

	var λsol float64 = GetSolarEclipticLongitudeLawrence(Msol, Csol)

	var Ae float64 = GetLunarAnnualEquationCorrection(M)

	var Eν float64 = GetLunarEvectionCorrection(M, λ, λsol)

	var Ca float64 = GetLunarMeanAnomalyCorrection(M, Msol, Ae, Eν)

	var ec EclipticCoordinate = GetLunarEclipticPosition(datetime)

	var got float64 = GetLunarEclipticLongitudeInXHours(ec.λ, Ca, 12)

	var want float64 = 57.438144

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetEclipticLatitudeInXHours(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	var ec EclipticCoordinate = GetLunarEclipticPosition(datetime)

	var Ωprime1 float64 = GetLunarCorrectedEclipticLongitudeOfTheAscendingNode(datetime)

	var λt1 float64 = GetLunarTrueEclipticLongitude(datetime)

	var got float64 = GetLunarEclipticLatitudeInXHours(ec.β, Ωprime1, λt1, 12)

	var want float64 = -3.470089

	if math.Abs(got-want) > 0.1 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarTransitJulianDate(t *testing.T) {
	var eq EquatorialCoordinate = GetLunarEquatorialPosition(d)

	var ϑ float64 = GetApparentGreenwhichSiderealTimeInDegrees(d)

	var got float64 = GetLunarTransitJulianDate(datetime, eq.α, longitude, ϑ)

	var want float64 = 2459348.890048

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarHorizontalCoordinatesForDay(t *testing.T) {
	horizontalCoordinates, err := GetLunarHorizontalCoordinatesForDay(datetime, longitude, latitude)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if len(horizontalCoordinates) != 1440 {
		t.Errorf("there is not enough horizontal coordinates for the day, expected 1440")
	}
}
