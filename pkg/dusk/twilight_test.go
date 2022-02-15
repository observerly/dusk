package dusk

import (
	"testing"
	"time"
)

func TestGetLocalCivilTwilightFrom(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalCivilTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Time = twilight.from

	var want = time.Date(1992, 4, 12, 19, 03, 52, 618345344, timezone)

	if got.Before(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if want.After(got) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !got.Equal(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetLocalCivilTwilightUntil(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalCivilTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Time = twilight.until

	var want = time.Date(1992, 4, 13, 5, 39, 45, 686235136, timezone)

	if got.Before(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if want.After(got) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if !got.Equal(want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
