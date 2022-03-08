package dusk

import (
	"math"
	"testing"
	"time"
)

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
