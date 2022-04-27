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

	var got time.Time = twilight.From

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

	var got time.Time = twilight.Until

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

func TestGetLocalCivilTwilightDuration(t *testing.T) {
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

	var got time.Duration = twilight.Duration

	var want time.Duration = 48260273900544

	if got.Nanoseconds() != want.Nanoseconds() {
		t.Errorf("got %d, wanted %d", got.Nanoseconds(), want.Nanoseconds())
	}
}

func TestGetLocalNauticalTwilightFrom(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalNauticalTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Time = twilight.From

	var want = time.Date(1992, 4, 12, 19, 29, 24, 855337216, timezone)

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

func TestGetLocalNauticalTwilightUntil(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalNauticalTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Time = twilight.Until

	var want = time.Date(1992, 4, 13, 5, 14, 14, 269436032, timezone)

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

func TestGetLocalNauticalTwilightDuration(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalNauticalTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Duration = twilight.Duration

	var want time.Duration = 51323107498880

	if got.Nanoseconds() != want.Nanoseconds() {
		t.Errorf("got %d, wanted %d", got.Nanoseconds(), want.Nanoseconds())
	}
}

func TestGetLocalAstronomicalTwilightFrom(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalAstronomicalTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Time = twilight.From

	var want = time.Date(1992, 4, 12, 19, 55, 12, 555572352, timezone)

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

func TestGetLocalAstronomicalTwilightUntil(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalAstronomicalTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Time = twilight.Until

	var want = time.Date(1992, 4, 13, 4, 48, 27, 39767424, timezone)

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

func TestGetLocalAstronomicalTwilightDuration(t *testing.T) {
	timezone, err := time.LoadLocation("Pacific/Honolulu")

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	twilight, location, err := GetLocalAstronomicalTwilight(d, longitude, latitude, elevation)

	if err != nil {
		t.Errorf("got %q", err)
		return
	}

	if timezone.String() != location.String() {
		t.Errorf("got %q, wanted %q", location, timezone)
	}

	var got time.Duration = twilight.Duration

	var want time.Duration = 54417566835968

	if got.Nanoseconds() != want.Nanoseconds() {
		t.Errorf("got %d, wanted %d", got.Nanoseconds(), want.Nanoseconds())
	}
}
