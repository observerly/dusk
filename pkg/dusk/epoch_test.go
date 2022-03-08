package dusk

import (
	"math"
	"testing"
	"time"
)

var datetime time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

var longitude float64 = -155.468094

func TestGetJulianDate(t *testing.T) {
	var got float64 = GetJulianDate(datetime)

	var want float64 = 2459348.5

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetUniversalTime(t *testing.T) {
	var got time.Time = GetUniversalTime(2459348.5)

	var want time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetGreenwhichSiderealTime(t *testing.T) {
	var got float64 = GetGreenwhichSiderealTime(datetime)

	var want float64 = 15.463990399019053

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetLocalSiderealTime(t *testing.T) {
	var got float64 = GetLocalSiderealTime(datetime, longitude)

	var want float64 = 5.099450799019053

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetCurrentJulianDayJLLawrence(t *testing.T) {
	var datetime time.Time = time.Date(2015, 2, 5, 12, 0, 0, 0, time.UTC)

	var got int = GetCurrentJulianDayRelativeToJ2000(datetime)

	var want int = 5514

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetFractionalJulianDayStandardEpoch(t *testing.T) {
	var datetime time.Time = time.Date(2015, 2, 5, 17, 0, 0, 0, time.UTC)

	var got float64 = GetFractionalJulianDaysSinceStandardEpoch(datetime)

	var want float64 = 5514.208333

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetCurrentJulianDay(t *testing.T) {
	var got int = GetCurrentJulianDayRelativeToJ2000(datetime)

	var want int = 7804

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCurrentJulianCentury(t *testing.T) {
	var got float64 = GetCurrentJulianCenturyRelativeToJ2000(datetime)

	var want float64 = 0.21364818617385353

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetCurrentJulianPeriodJulianDate(t *testing.T) {
	var period JulianPeriod = GetCurrentJulianPeriod(datetime)

	var got float64 = period.JD

	var want float64 = 2459348.5

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetCurrentJulianPeriodJulianCenturies(t *testing.T) {
	var period JulianPeriod = GetCurrentJulianPeriod(datetime)

	var got float64 = period.T

	var want float64 = 0.21364818617385353

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetMeanGreenwhichSiderealTimeInDegrees(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.89 of Meeus, Jean. 1991. Astronomical algorithms. Richmond,
	// Va: Willmann - Bell.:
	var datetime time.Time = time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC)

	var got float64 = GetMeanGreenwhichSiderealTimeInDegrees(datetime)

	var want float64 = 197.693195

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetApparentGreenwhichSiderealTimeInDegrees(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.88 of Meeus, Jean. 1991. Astronomical algorithms. Richmond,
	// Va: Willmann - Bell.:
	var datetime time.Time = time.Date(1987, 4, 10, 0, 0, 0, 0, time.UTC)

	var got float64 = GetApparentGreenwhichSiderealTimeInDegrees(datetime)

	var want float64 = 197.692600

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetApparentGreenwhichSiderealTimeInDegreesBonus(t *testing.T) {
	// For testing we need to specify a date because most calculations are
	// differential w.r.t a time component. We set it to the date provided
	// on p.103 of Meeus, Jean. 1991. Astronomical algorithms. Richmond,
	// Va: Willmann - Bell.:
	var datetime time.Time = time.Date(1988, 3, 20, 0, 0, 0, 0, time.UTC)

	var got float64 = GetApparentGreenwhichSiderealTimeInDegrees(datetime)

	var want float64 = 177.741993

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetMeanSolarTime(t *testing.T) {
	var got float64 = GetMeanSolarTime(datetime, longitude)

	var want float64 = 7804.431856

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
