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

	if got.Rise != nil {
		t.Errorf("got %v, but expected the object to never rise for the given paramaters", got)
	}

	if got.Set != nil {
		t.Errorf("got %v, but expected the object to never set for the given parameters", got)
	}
}

func TestGetObjectRiseObjectSetTimesInUTCLawrenceChapter5Exercise2(t *testing.T) {
	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	var got Transit = GetObjectRiseObjectSetTimesInUTC(datetime, EquatorialCoordinate{RightAscension: 243.675000, Declination: 25.9613889}, 38.250132, -78.300288)

	var rise = time.Date(2015, 6, 6, 20, 57, 48, 562000000, time.UTC)

	var set = time.Date(2015, 6, 7, 11, 55, 55, 501000000, time.UTC)

	if got.Rise.String() != rise.String() {
		t.Errorf("got %v, wanted %v", *got.Rise, rise)
	}

	if got.Set.String() != set.String() {
		t.Errorf("got %v, wanted %v", *got.Set, set)
	}
}

func TestGetObjectRiseObjectSetTimesChapter5Exercise1(t *testing.T) {
	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	got, err := GetObjectRiseObjectSetTimes(datetime, EquatorialCoordinate{RightAscension: 90, Declination: -60}, 45.250132, -100.300288)

	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}

	if got.Rise != nil {
		t.Errorf("got %v, but expected the object to never rise for the given paramaters", got)
	}

	if got.Set != nil {
		t.Errorf("got %v, but expected the object to never set for the given parameters", got)
	}
}

func TestGetObjectRiseObjectSetTimesChapter5Exercise2(t *testing.T) {
	timezone, _ := time.LoadLocation("America/New_York")

	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	got, err := GetObjectRiseObjectSetTimes(datetime, EquatorialCoordinate{RightAscension: 243.675000, Declination: 25.9613889}, 38.250132, -78.300288)

	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}

	var rise = time.Date(2015, 6, 6, 16, 57, 48, 562000000, timezone)

	var set = time.Date(2015, 6, 7, 7, 55, 55, 501000000, timezone)

	if rise.After(set) {
		t.Errorf("the object must rise before it sets")
	}

	if got.Rise.String() != rise.String() {
		t.Errorf("got %v, wanted %v", *got.Rise, rise)
	}

	if got.Set.String() != set.String() {
		t.Errorf("got %v, wanted %v", *got.Set, set)
	}
}

func TestGetObjectTransitMaximaTime(t *testing.T) {
	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	transit, err := GetObjectRiseObjectSetTimes(datetime, EquatorialCoordinate{RightAscension: 243.675000, Declination: 25.9613889}, 38.250132, -78.300288)

	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}

	got, err := GetObjectTransitMaximaTime(datetime, EquatorialCoordinate{RightAscension: 243.675000, Declination: 25.9613889}, 38.250132, -78.300288)

	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}

	if got.Before(*transit.Rise) || got.After(*transit.Set) {
		t.Errorf("maxima time must be between rise and set")
	}
}

func TestGetObjectTransitMaximaTimeNoRiseNoSet(t *testing.T) {
	var datetime time.Time = time.Date(2015, 6, 6, 0, 0, 0, 0, time.UTC)

	got, err := GetObjectTransitMaximaTime(datetime, EquatorialCoordinate{RightAscension: 90, Declination: -60}, 45.250132, -100.300288)

	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}

	if got != nil {
		t.Errorf("got %v, but expected the object to never reach a maxima above the horizon for the given paramaters", got)
	}

}
