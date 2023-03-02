package dusk

import (
	"math"
	"time"

	tzm "github.com/zsefvlol/timezonemapper"
)

type Transit struct {
	Rise     *time.Time
	Maximum  *time.Time
	Set      *time.Time
	Duration time.Duration
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
GetObjectRiseObjectSetTimesInUTCForDay()

@param datetime - the time to calculate the rise and set times for
@param eq - the EquatorialCoordinate{} of the object to calculate the rise and set times for
@param latitude - the latitude of the observer
@param longitude - the longitude of the observer
@returns a Transit struct which contains the rise and set times of the object in UTC
*/
func GetObjectRiseObjectSetTimesInUTCForDay(datetime time.Time, eq EquatorialCoordinate, latitude float64, longitude float64) Transit {
	if !GetDoesObjectRiseOrSet(eq, latitude) {
		return Transit{
			Rise:     nil,
			Set:      nil,
			Duration: 0,
		}
	}

	var d = time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, time.UTC)

	// see p.117 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
	LSTr := 24 + eq.RightAscension/15 - GetArgumentOfLocalSiderealTimeForTransit(latitude, eq.Declination)/15

	GSTr := ConvertLocalSiderealTimeToGreenwhichSiderealTime(LSTr, longitude)

	UTr := ConvertGreenwhichSiderealTimeToUniversalTime(datetime, GSTr)

	// for highest accuracy, convert hours to milliseconds to add:
	rise := d.Add(time.Duration(UTr*3600000) * time.Millisecond)

	// see p.117 of Lawrence, J.L. 2015. Celestial Calculations - A Gentle Introduction To Computational Astronomy. Cambridge, Ma: The MIT Press
	LSTs := eq.RightAscension/15 + GetArgumentOfLocalSiderealTimeForTransit(latitude, eq.Declination)/15

	GSTs := ConvertLocalSiderealTimeToGreenwhichSiderealTime(LSTs, longitude)

	UTs := ConvertGreenwhichSiderealTimeToUniversalTime(d, GSTs)

	// for highest accuracy, convert hours to milliseconds to add:
	set := d.Add(time.Duration(UTs*3600000) * time.Millisecond)

	return Transit{
		Rise:     &rise,
		Set:      &set,
		Duration: rise.Sub(set),
	}
}

/*
GetObjectHorizontalCoordinatesForDay()

@param datetime - the datetime of the observer (in UTC)
@params eq - the EquatorialCoordinate{} of the object to calculate the rise and set times for
@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
@returns the horizontal coordinates of the target object for every minute of a given day.
*/
func GetObjectHorizontalCoordinatesForDay(datetime time.Time, eq EquatorialCoordinate, longitude float64, latitude float64) ([]TransitHorizontalCoordinate, error) {
	// create an empty list of horizontalCoordinate structs:
	horizontalCoordinates := make([]TransitHorizontalCoordinate, 1442)

	// get the corresponding timezone for the longitude and latitude provided:
	timezone := tzm.LatLngToTimezoneString(latitude, longitude)

	location, err := time.LoadLocation(timezone)

	if err != nil {
		return horizontalCoordinates, err
	}

	var d = time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, location).In(time.UTC)

	// Subtract one minute to ensure we are not over looking the rise time to be
	d = d.Add(time.Minute * -1)

	for i := range horizontalCoordinates {
		var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(d, longitude, latitude, eq)

		if i > 0 {
			horizontalCoordinates[i] = TransitHorizontalCoordinate{
				Datetime: d.In(location),
				Altitude: hz.Altitude,
				Azimuth:  hz.Azimuth,
				IsRise:   hz.Altitude > 0 && horizontalCoordinates[i-1].Altitude <= 0,
				IsSet:    hz.Altitude < 0 && horizontalCoordinates[i-1].Altitude >= 0,
			}
		} else {
			horizontalCoordinates[i] = TransitHorizontalCoordinate{
				Datetime: d.In(location),
				Altitude: hz.Altitude,
				Azimuth:  hz.Azimuth,
				IsRise:   false,
				IsSet:    false,
			}
		}

		d = d.Add(time.Minute)
	}

	return horizontalCoordinates[1:1441], nil
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
			Rise:     nil,
			Set:      nil,
			Duration: 0,
		}
	}

	transit := GetObjectRiseObjectSetTimesInUTCForDay(datetime, eq, latitude, longitude)

	// We need to ensure that if the transit rise is before the transit set,
	if transit.Rise != nil && transit.Set.Before(*transit.Rise) {
		tomorrow := GetObjectRiseObjectSetTimesInUTCForDay(datetime.Add(time.Hour*24), eq, latitude, longitude)
		transit.Set = tomorrow.Set
	}

	return Transit{
		Rise:     transit.Rise,
		Set:      transit.Set,
		Duration: transit.Rise.Sub(*transit.Set),
	}
}

/*
GetObjectRiseObjectSetTimes()

@param datetime - the time to calculate the rise and set times for
@param eq - the EquatorialCoordinate{} of the object to calculate the rise and set times for
@param latitude - the latitude of the observer
@param longitude - the longitude of the observer
@returns a Transit struct which contains the rise and set times of the object in local time
*/
func GetObjectRiseObjectSetTimes(datetime time.Time, eq EquatorialCoordinate, latitude float64, longitude float64) (*Transit, error) {
	if !GetDoesObjectRiseOrSet(eq, latitude) {
		return &Transit{
			Rise:     nil,
			Set:      nil,
			Duration: time.Duration(0),
		}, nil
	}

	// get the corresponding timezone for the longitude and latitude provided:
	timezone := tzm.LatLngToTimezoneString(latitude, longitude)

	// the corresponding local timezone for the observer, e..g, the location name corresponding to a file in the IANA Time Zone database, such as "Pacific/Honolulu":
	location, err := time.LoadLocation(timezone)

	if err != nil {
		return nil, err
	}

	transit := GetObjectRiseObjectSetTimesInUTC(datetime, eq, latitude, longitude)

	rise := transit.Rise.In(location)

	set := transit.Set.In(location)

	return &Transit{
		Rise:     &rise,
		Set:      &set,
		Duration: transit.Duration,
	}, nil
}

/*
GetObjectTransitMaximaTime()

@param datetime - the time to calculate the rise and set times for
@param eq - the EquatorialCoordinate{} of the object to calculate the rise and set times for
@param latitude - the latitude of the observer
@param longitude - the longitude of the observer
@returns a the Transit maxima time of the object in local time
*/
func GetObjectTransitMaximaTime(datetime time.Time, eq EquatorialCoordinate, latitude float64, longitude float64) (*time.Time, error) {
	transit, err := GetObjectRiseObjectSetTimes(datetime, eq, latitude, longitude)

	if err != nil {
		return nil, err
	}

	// find the number of minutes between the rise and set times:
	minutes := 1440

	if transit.Duration.Minutes() > 0 {
		minutes = int(math.Ceil(math.Abs(transit.Duration.Minutes())))
	}

	// create an empty list of horizontalCoordinate structs:
	horizontalCoordinates := make([]TransitHorizontalCoordinate, minutes)

	// Start at midnight on for the datetime provided:
	d := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())

	if transit.Rise != nil {
		d = *transit.Rise
	}

	for i := range horizontalCoordinates {
		// Get the current horizontal position of the object:
		var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(d, longitude, latitude, eq)

		horizontalCoordinates[i] = TransitHorizontalCoordinate{
			Datetime: d,
			Altitude: hz.Altitude,
			Azimuth:  hz.Azimuth,
		}

		d = d.Add(time.Minute)

		// Since our object's initial direction is rising, we can assume the following comparison:
		// if (i > 0) && (horizontalCoordinates[i].Altitude < horizontalCoordinates[i-1].Altitude) {
		// 	maxima = &horizontalCoordinates[i-1].Datetime
		// }
	}

	// Get the maximum altitude from the list of horizontal coordinates:

	maximum := &horizontalCoordinates[0]

	for i := range horizontalCoordinates {
		if horizontalCoordinates[i].Altitude > maximum.Altitude {
			maximum = &horizontalCoordinates[i]
		}
	}

	return &maximum.Datetime, nil
}

/*
GetObjectTransit()

@param datetime - the time to calculate the rise and set times for
@param eq - the EquatorialCoordinate{} of the object to calculate the rise and set times for
@param latitude - the latitude of the observer
@param longitude - the longitude of the observer
@returns a Transit struct which contains the rise, maximum and set times of the object in local time
*/
func GetObjectTransit(datetime time.Time, eq EquatorialCoordinate, latitude float64, longitude float64) (*Transit, error) {
	transit, err := GetObjectRiseObjectSetTimes(datetime, eq, latitude, longitude)

	if err != nil {
		return nil, err
	}

	if transit.Rise == nil || transit.Set == nil {
		return &Transit{
			Rise:     nil,
			Set:      nil,
			Maximum:  nil,
			Duration: 0,
		}, nil
	}

	// find the number of minutes between the rise and set times:
	minutes := int(math.Ceil(math.Abs(transit.Duration.Minutes())))

	// create an empty list of horizontalCoordinate structs:
	horizontalCoordinates := make([]TransitHorizontalCoordinate, minutes)

	d := *transit.Rise

	for i := range horizontalCoordinates {
		// Get the current horizontal position of the object:
		var hz HorizontalCoordinate = ConvertEquatorialCoordinateToHorizontal(d, longitude, latitude, eq)

		horizontalCoordinates[i] = TransitHorizontalCoordinate{
			Datetime: d,
			Altitude: hz.Altitude,
			Azimuth:  hz.Azimuth,
		}

		d = d.Add(time.Minute)

		// Since our object's initial direction is rising, we can assume the following comparison:
		if (i > 0) && (horizontalCoordinates[i].Altitude < horizontalCoordinates[i-1].Altitude) {
			transit.Maximum = &horizontalCoordinates[i-1].Datetime
		}
	}

	return &Transit{
		Rise:     transit.Rise,
		Set:      transit.Set,
		Maximum:  transit.Maximum,
		Duration: transit.Duration,
	}, nil
}
