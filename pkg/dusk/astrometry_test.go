package dusk

import (
	"math"
	"testing"
)

func TestGetHourAngle(t *testing.T) {
	var LST float64 = GetLocalSiderealTime(datetime, longitude)

	var got float64 = GetHourAngle(88.7929583, LST)

	var want float64 = 347.698366

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetHourAngleBonus(t *testing.T) {
	var LST float64 = GetLocalSiderealTime(d, longitude)

	var got float64 = GetHourAngle(88.7929583, LST)

	var want float64 = 316.180845

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
