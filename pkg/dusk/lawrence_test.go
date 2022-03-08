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
