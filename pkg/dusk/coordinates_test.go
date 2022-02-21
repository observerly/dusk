package dusk

import (
	"math"
	"testing"
)

func TestConvertEclipticCoordinateToEquatorialRA(t *testing.T) {
	// utilising the ecliptic position of the moon on the datetime provided:
	var ec EquatorialCoordinate = ConvertEclipticCoordinateToEquatorial(d, EclipticCoordinate{λ: 133.162655, β: -3.229126, Δ: 0})

	var got float64 = ec.ra

	var want float64 = 134.688470

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEclipticCoordinateToEquatorialDec(t *testing.T) {
	// utilising the ecliptic position of the moon on the datetime provided:
	var ec EquatorialCoordinate = ConvertEclipticCoordinateToEquatorial(d, EclipticCoordinate{λ: 133.162655, β: -3.229126, Δ: 0})

	var got float64 = ec.dec

	var want float64 = 13.768368

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
