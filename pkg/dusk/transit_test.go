package dusk

import (
	"testing"
	"time"
)

func TestGetDoesObjectRiseOrSetBetelgeuseNorthernHemisphere(t *testing.T) {
	var got bool = GetDoesObjectRiseOrSet(EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639}, 38.778132)

	var want bool = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetDoesObjectRiseOrSetBetelgeuseSouthernHemisphere(t *testing.T) {
	var got bool = GetDoesObjectRiseOrSet(EquatorialCoordinate{RightAscension: 88.7929583, Declination: 7.4070639}, -89.191006)

	var want bool = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetDoesObjectRiseOrSetArcturusNorthernHemisphere(t *testing.T) {
	var got bool = GetDoesObjectRiseOrSet(EquatorialCoordinate{RightAscension: 213.9153, Declination: 19.182409}, 38.778132)

	var want bool = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetDoesObjectRiseOrSetArcturusSouthernHemisphere(t *testing.T) {
	var got bool = GetDoesObjectRiseOrSet(EquatorialCoordinate{RightAscension: 213.9153, Declination: 19.182409}, -89.191006)

	var want bool = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetObjectRiseObjectSetTimesInUTCLawrenceChapter5Exercise1(t *testing.T) {
	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	var got Transit = GetObjectRiseObjectSetTimesInUTC(datetime, EquatorialCoordinate{RightAscension: 90, Declination: -60}, 45.250132, -100.300288)

	if got.rise != nil {
		t.Errorf("got %v, but expected the object to never rise for the given paramaters", got)
	}

	if got.set != nil {
		t.Errorf("got %v, but expected the object to never set for the given parameters", got)
	}
}

func TestGetObjectRiseObjectSetTimesInUTCLawrenceChapter5Exercise2(t *testing.T) {
	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	var got Transit = GetObjectRiseObjectSetTimesInUTC(datetime, EquatorialCoordinate{RightAscension: 243.675000, Declination: 25.9613889}, 38.250132, -78.300288)

	var rise = time.Date(2015, 6, 6, 20, 57, 48, 562000000, time.UTC)

	var set = time.Date(2015, 6, 6, 11, 59, 51, 410000000, time.UTC)

	if got.rise.String() != rise.String() {
		t.Errorf("got %v, wanted %v", *got.rise, rise)
	}

	if got.set.String() != set.String() {
		t.Errorf("got %v, wanted %v", *got.set, set)
	}
}
