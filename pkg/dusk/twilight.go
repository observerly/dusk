package dusk

import (
	"time"

	tzm "github.com/zsefvlol/timezonemapper"
)

type SunriseStatus int

const (
	AboveHorizon = SunriseStatus(1)
	AtHorizon    = SunriseStatus(0)
	BelowHorizon = SunriseStatus(-1)
)

type Twilight struct {
	From     time.Time
	Until    time.Time
	Duration time.Duration
}

// For all twilight funcs, please reference for information on timezones and their respective locations:
// @see https://en.wikipedia.org/wiki/list_of_tz_database_time_zones
// @see https://www.iana.org/time-zones
// @see https://pkg.go.dev/time#LoadLocation

/*
	GetLocalTwilight()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@param degreesBelowHorizon - is the degrees below horizon for the designated "twilight period", with 0° being "night" e.g., as soon as the sun is below the horizon.
	@returns the start and end times of Civil Twilight, as designated by when the Sun is -6 degrees below the horizon.
*/
func GetLocalTwilight(datetime time.Time, longitude float64, latitude float64, elevation float64, degreesBelowHorizon float64) (*Twilight, *time.Location, error) {
	// get the corresponding timezone for the longitude and latitude provided:
	timezone := tzm.LatLngToTimezoneString(latitude, longitude)

	var s Sun = GetSunriseSunsetTimesInUTC(datetime, degreesBelowHorizon, longitude, latitude, elevation)

	var r Sun = GetSunriseSunsetTimesInUTC(datetime.Add(time.Hour*24), degreesBelowHorizon, longitude, latitude, elevation)

	// the corresponding local timezone for the observer, e..g, the location name corresponding to a file in the IANA Time Zone database, such as "Pacific/Honolulu":
	location, err := time.LoadLocation(timezone)

	if err != nil {
		return nil, nil, err
	}

	return &Twilight{
		From:     s.set.In(location),
		Until:    r.rise.In(location),
		Duration: r.set.Sub(r.rise),
	}, location, nil
}

/*
	GetLocalCivilTwilight()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@returns the start and end times of Civil Twilight, as designated by when the Sun is -6 degrees below the horizon.
*/
func GetLocalCivilTwilight(datetime time.Time, longitude float64, latitude float64, elevation float64) (*Twilight, *time.Location, error) {
	// civil twilight is designated as being 6 degrees below horizon:
	var degreesBelowHorizon float64 = -6

	return GetLocalTwilight(datetime, longitude, latitude, elevation, degreesBelowHorizon)
}

/*
	GetLocalNauticalTwilight()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@returns the start and end times of Nautical Twilight, as designated by when the Sun is -12 degrees below the horizon.
*/
func GetLocalNauticalTwilight(datetime time.Time, longitude float64, latitude float64, elevation float64) (*Twilight, *time.Location, error) {
	// nautical twilight is designated as being 6 degrees below horizon:
	var degreesBelowHorizon float64 = -12

	return GetLocalTwilight(datetime, longitude, latitude, elevation, degreesBelowHorizon)
}

/*
	GetLocalAstronomicalTwilight()

	@param datetime - the datetime of the observer (in UTC)
	@param longitude - is the longitude (west is negative, east is positive) in degrees of some observer on Earth
	@param latitude - is the latitude (south is negative, north is positive) in degrees of some observer on Earth
	@param elevation - is the elevation (above sea level) in meters of some observer on Earth
	@returns the start and end times of Astronomical Twilight, as designated by when the Sun is -18 degrees below the horizon.
*/
func GetLocalAstronomicalTwilight(datetime time.Time, longitude float64, latitude float64, elevation float64) (*Twilight, *time.Location, error) {
	// astronomical twilight is designated as being 18 degrees below horizon:
	var degreesBelowHorizon float64 = -18

	return GetLocalTwilight(datetime, longitude, latitude, elevation, degreesBelowHorizon)
}
