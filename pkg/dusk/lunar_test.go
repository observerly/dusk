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

func TestGetLunarEquatorialPositionRightAscension(t *testing.T) {
	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = eq.ra

	var want float64 = 76.239624

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarEquatorialPositionDeclination(t *testing.T) {
	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = eq.dec

	var want float64 = 23.598793

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLunarHourAngle(t *testing.T) {
	var eq EquatorialCoordinate = GetLunarEquatorialPosition(datetime)

	var got float64 = GetLunarHourAngle(eq.dec, latitude, 0)

	var want float64 = 97.466043

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
