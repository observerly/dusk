package dusk

import (
	"math"
	"time"
)

type Transit struct {
	rise *time.Time
	set  *time.Time
}

/*
  GetDoesObjectRiseOrSet()

  @returns a boolean which determines if the object's EquatorialCoordinate{} in question rises or sets for the given Observer's latitude
  @see p.117 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
*/
func GetDoesObjectRiseOrSet(eq EquatorialCoordinate, latitude float64) bool {
	// If |Ar| > 1, the object never rises above the horizon
	Ar := sinx(eq.Declination) / cosx(latitude)

	// If |H1| > 1, the object is always below the horizon
	H1 := tanx(latitude) * tanx(eq.Declination)

	return math.Abs(Ar) < 1 && math.Abs(H1) < 1
}

/*
	GetObjectRiseObjectSetTimesInUTC()

	@param datetime - the time to calculate the rise and set times for
	@param eq - the EquatorialCoordinate{} of the object to calculate the rise and set times for
	@param latitude - the latitude of the observer
	@param longitude - the longitude of the observer
	@returns a Transit struct which contains the rise and set times of the object in UTC
*/
func GetObjectRiseObjectSetTimesInUTC(datetime time.Time, eq EquatorialCoordinate, latitude float64, longitude float64) Transit {
	if !GetDoesObjectRiseOrSet(eq, latitude) {
		return Transit{
			rise: nil,
			set:  nil,
		}
	}

	// see p.117 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
	LSTr := 24 + eq.RightAscension/15 - GetArgumentOfLocalSiderealTimeForTransit(latitude, eq.Declination)/15

	GSTr := ConvertLocalSiderealTimeToGreenwhichSiderealTime(LSTr, longitude)

	UTr := ConvertGreenwhichSiderealTimeToUniversalTime(datetime, GSTr)

	// for highest accuracy, convert hours to milliseconds to add:
	rise := datetime.Add(time.Duration(UTr*3600000) * time.Millisecond)

	// see p.117 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
	LSTs := eq.RightAscension/15 + GetArgumentOfLocalSiderealTimeForTransit(latitude, eq.Declination)/15

	GSTs := ConvertLocalSiderealTimeToGreenwhichSiderealTime(LSTs, longitude)

	UTs := ConvertGreenwhichSiderealTimeToUniversalTime(datetime, GSTs)

	// for highest accuracy, convert hours to milliseconds to add:
	set := datetime.Add(time.Duration(UTs*3600000) * time.Millisecond)

	return Transit{
		rise: &rise,
		set:  &set,
	}
}
