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

func TestGetCurrentJulianDay(t *testing.T) {
	var got int = GetCurrentJulianDayRelativeToJ2000(datetime)

	var want int = 7804

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetMeanSolarTime(t *testing.T) {
	var got float64 = GetMeanSolarTime(datetime, longitude)

	var want float64 = 7804.431856

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
