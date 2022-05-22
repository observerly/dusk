package dusk

import "testing"

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
