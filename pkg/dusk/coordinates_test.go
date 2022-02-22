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

func TestConvertEquatorialCoordinateTHorizontalAltitude(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{ra: 88.7929583, dec: 7.4070639})

	var got float64 = hz.a

	var want float64 = 72.800882

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEquatorialCoordinateTHorizontalAzimuth(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{ra: 88.7929583, dec: 7.4070639})

	var got float64 = hz.A

	var want float64 = 134.397750

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
