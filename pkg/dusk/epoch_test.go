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

func TestGetMeanSolarTime(t *testing.T) {
	var got float64 = GetMeanSolarTime(datetime, longitude)

	var want float64 = 7804.431856

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
