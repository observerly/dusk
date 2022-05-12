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

var arcturus Coordinate = Coordinate{ϕ: 19.1825, θ: 213.9154}

var spica Coordinate = Coordinate{ϕ: -11.1614, θ: 201.2983}

var denebola Coordinate = Coordinate{ϕ: 14.5720581, θ: 177.2649}

func TestGetAngularSeparationArcturusSpica(t *testing.T) {
	var got float64 = GetAngularSeparation(arcturus, spica)

	var want float64 = 32.793027

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetAngularSeparationSpicaDenebola(t *testing.T) {
	var got float64 = GetAngularSeparation(spica, denebola)

	var want float64 = 35.064334

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetAngularSeparationDenebolaArcturus(t *testing.T) {
	var got float64 = GetAngularSeparation(denebola, arcturus)

	var want float64 = 35.309668

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestGetAngularSeparationZero(t *testing.T) {
	var coord1 Coordinate = Coordinate{ϕ: 0, θ: 0}

	var coord2 Coordinate = Coordinate{ϕ: 0, θ: 0}

	var got float64 = GetAngularSeparation(coord1, coord2)

	var want float64 = 0

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
