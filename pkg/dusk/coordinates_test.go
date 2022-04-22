package dusk

import (
	"math"
	"testing"
	"time"
)

func TestConvertEclipticCoordinateToEquatorialRA(t *testing.T) {
	// utilising the ecliptic position of the moon on the datetime provided:
	var eq EquatorialCoordinate = ConvertEclipticCoordinateToEquatorial(d, EclipticCoordinate{λ: 133.162655, β: -3.229126, Δ: 0})

	var got float64 = eq.RightAscension

	var want float64 = 134.688470

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEclipticCoordinateToEquatorialRAAlt(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	// utilising the ecliptic position of the moon on the datetime provided:
	var eq EquatorialCoordinate = ConvertEclipticCoordinateToEquatorial(datetime, EclipticCoordinate{λ: 50.279952, β: -2.981288, Δ: 0})

	var got float64 = eq.RightAscension

	var want float64 = 48.662544

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEclipticCoordinateToEquatorialDec(t *testing.T) {
	// utilising the ecliptic position of the moon on the datetime provided:
	var ec EquatorialCoordinate = ConvertEclipticCoordinateToEquatorial(d, EclipticCoordinate{λ: 133.162655, β: -3.229126, Δ: 0})

	var got float64 = ec.Declination

	var want float64 = 13.768368

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEclipticCoordinateToEquatorialDecAlt(t *testing.T) {
	// Date of observation:
	var datetime time.Time = time.Date(2015, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	// utilising the ecliptic position of the moon on the datetime provided:
	var ec EquatorialCoordinate = ConvertEclipticCoordinateToEquatorial(datetime, EclipticCoordinate{λ: 50.279952, β: -2.981288, Δ: 0})

	var got float64 = ec.Declination

	var want float64 = 14.941252

	if math.Abs(got-want) > 0.15 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEquatorialCoordinateTHorizontalAltitude(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639})

	var got float64 = hz.a

	var want float64 = 72.800882

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestConvertEquatorialCoordinateTHorizontalAzimuth(t *testing.T) {
	var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(datetime, longitude, latitude, EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639})

	var got float64 = hz.A

	var want float64 = 134.397750

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}
