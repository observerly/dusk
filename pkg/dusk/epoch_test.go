package dusk

import (
	"math"
	"testing"
	"time"
)

var datetime time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

func TestGetJulianDate(t *testing.T) {
	var got float64 = GetJulianDate(datetime)

	var want float64 = 2459348.5

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}